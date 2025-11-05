package test

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type XKCD struct {
	Month  string `json:"month"`
	Year   string `json:"year"`
	Title  string `json:"title"`
	ImgURL string `json:"img"`
}

const BaseUrl = `https://xkcd.com/`

func GetXKCDAndSave() {
	// 创建文件
	file, err := os.Create("./xkcd.csv")
	if err != nil {
		fmt.Printf("Create file error, err =  %s\n", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	err = writer.Write([]string{"Month", "Year", "Title", "ImgURL"})
	if err != nil {
		fmt.Printf("Write file error, err =  %s\n", err)
		return
	}
	for i := 1; i <= 571; i++ {
		xkcd := &XKCD{}
		url := fmt.Sprintf("%s%d/info.0.json", BaseUrl, i)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("http.Get error, err =  %s\n", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("resp.StatusCode error, err =  %s\n", resp.Status)
		}
		if err = json.NewDecoder(resp.Body).Decode(&xkcd); err != nil {
			fmt.Printf("json.Decode error, err =  %s\n", err)
		}
		err = writer.Write([]string{xkcd.Month, xkcd.Year, xkcd.Title, xkcd.ImgURL})
		if err != nil {
			fmt.Printf("Write file error, err =  %s\n", err)
		}
		writer.Flush()
	}
}

func ResImgUrl(key string) string {
	file, err := os.Open("./xkcd.csv")
	if err != nil {
		fmt.Printf("Open file error, err =  %s\n", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Read file error, err =  %s\n", err)
		}
		if key == record[2] {
			return record[3]
		}
	}
	return "未找到"
}
