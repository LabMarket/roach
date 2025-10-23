package main

import (
	"fmt"
	"io"
	"os"

	"LabMarket/roach/formatter"
)

func main() {
	args := os.Args[1:]

	var f []byte
	var err error
	if len(args) == 0 {
		f, err = io.ReadAll(os.Stdin)
	} else {
		f, err = os.ReadFile(args[0])
		if err != nil {
			fmt.Println("Formatter: cannot read file", err.Error())
			os.Exit(1)
		}
	}

	formatter := formatter.New(string(f))
	formatter.Format()
}
