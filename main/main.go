package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		inFile  *os.File
		outFile *os.File
		err     error
	)
	if len(os.Args) != 3 && len(os.Args) != 1 {
		printUsage()
		os.Exit(1)
	}
	if len(os.Args) == 1 {
		inFile = os.Stdin
		outFile = os.Stdout
	} else {
		inFile, err = os.Open(os.Args[1])
		if os.IsNotExist(err) || os.IsPermission(err) {
			fmt.Println(err)
			os.Exit(2)
		}
		defer inFile.Close()
		outFile, err = os.Create(os.Args[1])
		if os.IsNotExist(err) || os.IsPermission(err) {
			fmt.Println(err)
			os.Exit(3)
		}
	}
	defer outFile.Close()

}

func printUsage() {
	fmt.Printf("Usage: \n\t\t%s inputFile.envtmpl output", os.Args[0])
	fmt.Printf(" \t\tcat inputFile.envtmpl < %s > output", os.Args[0])
	fmt.Println()
}
