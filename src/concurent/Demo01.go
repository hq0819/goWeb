package main

import (
	"bufio"
	"os"
)

func main() {
	dir := "src/pic/pic01.jpg"
	open, _ := os.Open(dir)

	bufio.NewReader(open)
	//bufio.NewWriter()
}
