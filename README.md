# line-editor

This is a simple edit tool for a text file. Change each line in the stream, using javaScript snippet for rebuild these lines.

Please install Golang if you do not have Golang alredy installed.

* Add GOPATH/bin to your PATH:
    * for bash:
    `echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc && source ~/.bashrc`
    * for zsh:
    `echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc && source ~/.zshrc`
* Install line-editor:
    `go get github.com/AntonStolov/line-editor`

## How to use:

* Select file using '--file' flag
* Write javascript code to change each line from the file using '--script' flag

## Example:
Change commas in the CSV file to the pipe lines '|' if number of rows equels 4

`line-editor --file test.csv --script "String(input).split(',').length == 4 ? String(input).split(',').join('|') : ',,,'"`