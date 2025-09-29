package main

import (
	"fmt"
	"os"

	"github.com/LabMarket/roach/cli"
)

func main() {
	if err := cli.Main(); err != nil {

		// When running the standard roach binary, we want to exit on error.
		// A custom build that imports the cli package can choose to handle
		// the error differently.
		if !cli.IsCustom {
			fmt.Fprintf(os.Stderr, "Error:\n%s\n", err)
			os.Exit(1)
		}
	}
}
