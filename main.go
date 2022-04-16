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
	index := flag.Int("c", 0, "which column to select")
	flag.Parse()
	for {
		line, err := reader.ReadString('\n')
		words := strings.Split(line, " ")
		if *index < len(words) {
			fmt.Print(words[*index])
		} else {
			return
		}
		if err == io.EOF {
			os.Stderr.WriteString("")
			return
		}
	}
}
