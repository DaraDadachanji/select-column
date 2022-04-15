package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	line := os.Args[0]
	index := os.Args[1]
	columnNumber, err := strconv.Atoi(index)
	if err != nil {
		log.Fatal("column index must be an integer")
	}
	columns := strings.Split(line, " ")
	if columnNumber < len(columns) {
		column := columns[columnNumber]
		fmt.Print(column)
	}

}
