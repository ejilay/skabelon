package main

import (
	"fmt"
	"github.com/ejilay/skabelon"
	"io/ioutil"
	"os"
	"path"
)

func main() {

	var (
		inFile, outFile *os.File
		err             error
	)

	if len(os.Args) != 3 {
		printUsage()
		os.Exit(1)
	}
	inFile, err = os.Open(os.Args[1])
	if os.IsPermission(err) || os.IsNotExist(err) {
		exitWithCode(err, 2)
	}

	outFile, err = os.Create(os.Args[2])
	if os.IsPermission(err) || os.IsNotExist(err) {
		exitWithCode(err, 3)
	}

	tmplBuf, err := ioutil.ReadAll(inFile)
	if err != nil {
		exitWithCode(err, 4)
	}
	skabelon := skabelon.New()
	if err = skabelon.Parse(path.Base(os.Args[1]), string(tmplBuf)); err != nil {
		exitWithCode(err, 5)
	}
	if err = skabelon.Exec(outFile); err != nil {
		exitWithCode(err, 6)
	}
}
func exitWithCode(err error, code int) {
	fmt.Println(err)
	os.Exit(code)
}

func printUsage() {
	fmt.Printf("Usage: \n\t\t%s inputFile.envtmpl output", os.Args[0])
	fmt.Printf(" \t\tcat inputFile.envtmpl < %s > output", os.Args[0])
	fmt.Println()
}
