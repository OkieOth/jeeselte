package main

import (
	"flag"
	"fmt"
	"os"
)

var usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func checkCommandlineParams(templateFile *string, inputFile *string, outputFile *string) {
	if len(*templateFile) == 0 {
		usage()
		fmt.Printf("\nno template file is given!\n\n")
		os.Exit(1)
	}
	if len(*inputFile) == 0 {
		usage()
		fmt.Printf("no input file is given!\n\n")
		os.Exit(1)
	}

	fmt.Printf("template file to use: %s\n", *templateFile)
	fmt.Printf("input file to use: %s\n", *inputFile)
	if len(*outputFile) == 0 {
		fmt.Printf("output will be written to stdout\n")
	} else {
		fmt.Printf("output file to use: %s\n", *outputFile)
	}
}

func main() {
	templateFile := flag.String("template", "", "path to template file to define the processing (required)")
	inputFile := flag.String("input", "", "path to input file to process (required)")
	ouputFile := flag.String("output", "", "path to the output file, if not given stdout is used insteed")
	flag.Parse()
	checkCommandlineParams(templateFile, inputFile, ouputFile)

	fmt.Println("\nhello :)")
	/*
		stylesheet := new(types.Stylesheet)
		fmt.Printf("Template file to use: %s\n", *templateFile)
		fmt.Println(stylesheet)
	*/
}
