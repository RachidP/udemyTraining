package main

import (
	"bufio"
	"fmt"
)

// ByteCounter exemple
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	advance, token, err := bufio.ScanWords(p, atEOF)

	return len(p), nil
}

func main() {
	var c WordCounter
	c.Write([]byte("hello world Ciao bello"))
	fmt.Println(c)
}
