package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"uuid"

	"github.com/atotto/clipboard"
)

var noCopy = flag.Bool("no-copy", false, "Don't copy output to clipboard.")
var noPrint = flag.Bool("no-print", false, "Don't print output to stdout.")
var help = flag.Bool("h", false, "Show this help message and exit.")
var version = flag.Int("v", 4, "Determine the UUID version to use. 4 or 7 accepted.")
var amount = flag.Int("n", 1, "The amount of UUIDs to generate.")

func printHelp() {
	fmt.Println("Usage: guid [OPTIONS]\nGenerate a GUID/UUID and copy it to the clipboard.\n\n" +
		"Options:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = printHelp
	flag.Parse()
	if *help {
		printHelp()
		os.Exit(0)
	}
	var generatorFunction func() uuid.UUID = uuid.NewV4
	switch *version {
	case 4:
		generatorFunction = uuid.NewV4
	case 7:
		generatorFunction = uuid.NewV7
	default:
		_, _ = fmt.Fprintf(os.Stderr, "Unsupported version: %d. Should be 4 or 7.\n", *version)
		os.Exit(1)
	}
	if *amount <= 0 {
		_, _ = fmt.Fprintf(os.Stderr, "Amount must be greater than 0.\n")
		os.Exit(1)
	}
	if *amount == 1 {
		newUuid := generatorFunction().String()
		if !*noPrint {
			fmt.Println(newUuid)
		}
		if !*noCopy {
			err := clipboard.WriteAll(newUuid)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "Could not copy to clipboard:", err)
			}
		}
		os.Exit(0)
	}

	output := make([]string, 0, *amount)
	for i := 0; i < *amount; i++ {
		newUuid := generatorFunction().String()
		output = append(output, newUuid)
	}
	outputStr := strings.Join(output, "\n")
	if !*noPrint {
		_, _ = fmt.Fprintln(os.Stdout, outputStr)
	}
	if !*noCopy {
		err := clipboard.WriteAll(outputStr)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Could not copy to clipboard:", err)
		}
	}
}
