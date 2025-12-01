package main

import (
	"errors"
	"fmt"
	"syscall"
)

func main() {
	fmt.Println(errors.New("eof") == errors.New("eof"))

	var err error
	fmt.Printf("%T\n", err)
	err = errors.New("eof")
	fmt.Printf("%T\n", err)
	err = syscall.Errno(2)
	fmt.Printf("%T\n", err)
	fmt.Println(err.Error())
}
