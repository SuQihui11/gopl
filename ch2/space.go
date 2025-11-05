package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string = "sqh"

func init() {
	cwd, err := os.Getwd() // compile error: unused: cwd
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	fmt.Println("cwd:", cwd)
}
func main() {
	fmt.Printf("%s", cwd)
}
