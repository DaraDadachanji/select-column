package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	column := parseColumn()
	flag.Parse()
	for {
		line, err := reader.ReadString('\n')
		words := strings.Split(line, " ")
		var index int
		if column >= 0 {
			index = column
		} else { //reverse from final column
			index = len(words) - column
		}
		if index < len(words) && index >= 0 {
			fmt.Println(words[index])
		} else {
			fmt.Println()
		}
		if err == io.EOF {
			return
		}
	}
}

func parseColumn() int {
	flag.Parse()
	if flag.NArg() == 0 {
		return 0
	}
	arg := flag.Arg(0)
	integer, err := strconv.Atoi(arg)
	if err != nil {
		os.Stderr.WriteString("invalid column selector, must be an integer")
		os.Exit(1)
	}
	return integer
}
