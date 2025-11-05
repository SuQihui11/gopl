package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
)

func main() {

	var method string

	flag.StringVar(&method, "method", "1", "编码方式（1 - SHA256, 2 - SHA384，3 - SHA512）")
	flag.Parse()

	ValidateMethod(method)

	fmt.Print("Enter a string: ")
	var input string
	fmt.Scanln(&input)
	var inputBytes = []byte(input)
	if method == "1" {
		fmt.Printf("%x\n", sha256.Sum256(inputBytes))
	} else if method == "2" {
		fmt.Printf("%x\n", sha512.Sum384(inputBytes))
	} else {
		fmt.Printf("%x\n", sha512.Sum512(inputBytes))
	}
}

func ValidateMethod(m string) {
	if m != "1" && m != "2" && m != "3" {
		log.Fatalf("method %s is not supported.", m)
	}
}
