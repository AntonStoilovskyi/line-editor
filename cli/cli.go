package cli

import (
	"flag"
	"io/ioutil"

	// "fmt"
	"log"
	"os"
)

type Params struct {
	File   string
	Script string
}

func Cli() Params {
	var parameters Params
	filePtr := flag.String("file", "", "File to parse. (Required)")
	scriptSnippetPtr := flag.String("script", "input", "JavaScript snippet. {input} is always processing line. Example: \"String(input).split('t')[0];\"")
	scriptSnippetFilePtr := flag.String("script-file", "", "Path to JavaScript file")
	flag.Parse()

	if *filePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Printf("\nfile: %s\nscript: %s\n", *filePtr, *scriptSnippetPtr)

	parameters.File = *filePtr
	if *scriptSnippetFilePtr != "" {
		parameters.Script = readScriptFile(*scriptSnippetFilePtr)
	} else {
		parameters.Script = *scriptSnippetPtr
	}

	return parameters
}

func readScriptFile(path string) string {

	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
