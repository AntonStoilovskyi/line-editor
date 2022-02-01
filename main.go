package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/AntonStolov/line-editor/cli"
	"github.com/robertkrimen/otto"
)

func main() {

	parameters := cli.Cli()
	counter := 0
	lineBefore := ""

	file, err := os.Open(parameters.File)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counter++
		manipulator(parameters.Script, scanner.Text(), fmt.Sprintf("%d", counter), lineBefore)
		lineBefore = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func manipulator(jsCommand string, input string, lineCounter string, lineBefore string) {
	script := `
		var lineBefore = "` + lineBefore + `";
		var line =` + lineCounter + `;
		var value = (function() {
			var input = "` + input + `";
			` +
		`return ` + jsCommand +
		`
		})();
	`

	vm := otto.New()
	vm.Run(script)

	if value, err := vm.Get("value"); err == nil {
		if valueInt, err := value.ToString(); err == nil {
			fmt.Println(valueInt)
		}
	}
}
