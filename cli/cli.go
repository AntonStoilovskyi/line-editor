package cli

import (
	"flag"
	// "fmt"
	"log"
	"os"
)

type Params struct {
    File string
    Script  string
}

func Cli() Params {
	var paramiters Params
    textPtr := flag.String("file", "", "File to parse. (Required)")
    metricPtr := flag.String("script", "input", "JavaScript snippet. {input} is always processing line. Example: \"String(input).split(\"t\")[0];\"")
    flag.Parse()

    if *textPtr == "" {
        flag.PrintDefaults()
        os.Exit(1)
    }

    log.Printf("\nfile: %s\nscript: %s\n", *textPtr, *metricPtr)

	paramiters.File = *textPtr
	paramiters.Script = *metricPtr

	return paramiters
}