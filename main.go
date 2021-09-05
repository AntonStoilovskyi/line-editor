package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/robertkrimen/otto"
	"github.com/AntonStolov/line-editor/cli"

)

func main() {

	paramiters := cli.Cli()
	counter := 0;
	
	file, err := os.Open(paramiters.File)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		counter++;
		manipulator(paramiters.Script, scanner.Text(), fmt.Sprintf("%d", counter))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	
}

func manipulator(jsCommand string, input string, lineCounter string) {
	script := `
		var line =` + lineCounter + `;
		var value = (function() {
			var input = "` + input + `";
			` +
			`return ` + jsCommand +
		`
		})();
	`
	// fmt.Printf("full processing script: %s\n", script)
	// fmt.Println(script)
	vm := otto.New()
	vm.Run(script)

	if value, err := vm.Get("value"); err == nil {
		if value_int, err := value.ToString(); err == nil {
			fmt.Println(value_int)
		}
	}
}