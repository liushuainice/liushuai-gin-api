package main

import (
	"bytes"
	"fmt"
)

func appendStr() string {
	var b bytes.Buffer
	b.WriteString("hello")
	b.WriteString(" world")
	return b.String()
}

func main() {
	str := appendStr()
	fmt.Println(str)
}
