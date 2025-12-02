package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var priceTable = template.Must(template.New("priceAndItem").Parse(`
<!DOCTYPE html>
<html>
<head>
<style>
table { border-collapse: collapse; width: 50%; }
th, td { text-align: left; padding: 8px; border-bottom: 1px solid #ddd; }
th { background-color: #f2f2f2; }
a { text-decoration: none; color: black; font-weight: bold; }
a:hover { color: blue; }
</style>
</head>
<body>

<h2>DataBase</h2>
<table>
  <tr>
    <th><a>Item</a></th>
    <th><a>Price</a></th>
  </tr>
	{{range $k, $v := .}}
  <tr>
    <td>{{$k}}</td> <td>{{$v}}</td> </tr>
  </tr>
  {{end}}
</table>

</body>
</html>
`))

func main() {
	db := database{"shoes": 50, "socks": 5}
	//log.Fatal(http.ListenAndServe(":8080", db))
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/update", db.crateOrUpdate)
	mux.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.FormValue("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		//w.WriteHeader(http.StatusNotFound)
		//fmt.Fprintf(w, "no such page: %q\n", req.URL.Path)
		msg := fmt.Sprintf("no such file or directory: %q\n", req.URL.Path)
		http.Error(w, msg, http.StatusNotFound)
	}
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if priceTable.Execute(w, db) != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error\n")
	}
}
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
func (db database) crateOrUpdate(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	price := req.FormValue("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "missing arguments\n")
		return
	}
	v, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid arguments\n")
		return
	}
	db[item] = dollars(v)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s updated to %s\n", item, dollars(v))
	return
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db, item)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s deleted from %s\n", item, db[item])
	return
}
