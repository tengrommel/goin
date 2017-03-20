package main

import (
	"fmt"
	"net"
	"io"
)

func main() {
	c, _ := net.Dial("tcp", ":2100")
	var r io.Reader
	r = c
	if _, ok := r.(io.Writer); ok{
		fmt.Println("We didn't forget there is a write inside value c")
	}
}
