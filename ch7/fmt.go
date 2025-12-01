package main

import (
	"fmt"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (n int, err error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	_, err := c.Write([]byte("hello"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

}
