package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	column := parseColumnArgument()
	for {
		line, readErr := reader.ReadString('\n')
		words := splitColumns(line)
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

// parseColumnArgument parses the column selector from command line arguments
func parseColumnArgument() int {
	args := os.Args[1:]
	if len(args) > 1 {
		os.Stderr.WriteString("too many arguments")
		os.Exit(1)
	} else if len(args) == 0 {
		return 1
	}
	arg := args[0]

	integer, err := strconv.Atoi(arg)
	if err != nil {
		os.Stderr.WriteString("invalid column selector, must be an integer")
		os.Exit(1)
	}
	return integer
}

// getColumnIndex calculates the column index to select
// if positive, converts 1-based index into 0-based index
// if negative, returns the nth last column
// (e.g. column -2 for 5 columns would be 3, which is the 4th column)
func getColumnIndex(column int, wordCount int) (int, error) {
	var index int
	if column > 0 {
		index = column - 1
	} else if column == 0 {
		index = 0
	} else { //reverse from final column
		index = wordCount - column
	}
	if index < wordCount && index >= 0 {
		return index, nil
	} else {
		return 0, errors.New("index out of range")
	}
}

// splitColumns splits a string using any amount of whitespace as a delimniter
func splitColumns(line string) []string {
	line = strings.TrimSpace(line)
	split := strings.Split(line, " ")
	var words []string
	for _, word := range split {
		if word != "" {
			words = append(words, word)
		}
	}
	return words
}
