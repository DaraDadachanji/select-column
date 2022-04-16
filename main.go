package main

import (
	"bufio"
	"errors"
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
		line, readErr := reader.ReadString('\n')
		words := strings.Split(line, " ")
		index, err := getColumnIndex(column, len(words))
		if err != nil {
			fmt.Println() //blank line if index out of range
		} else {
			fmt.Println(words[index])
		}
		if readErr == io.EOF {
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

func getColumnIndex(column int, wordCount int) (int, error) {
	var index int
	if column >= 0 {
		index = column
	} else { //reverse from final column
		index = wordCount - column
	}
	if index < wordCount && index >= 0 {
		return index, nil
	} else {
		return 0, errors.New("index out of range")
	}
}
