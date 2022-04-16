package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	column := flag.Int("c", 1, "which column to select")
	index := *column - 1
	flag.Parse()
	for {
		line, err := reader.ReadString('\n')
		words := strings.Split(line, " ")
		if index < len(words) {
			fmt.Println(words[index])
		} else {
			return
		}
		if err == io.EOF {
			return
		}
	}
}
