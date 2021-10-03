# line-editor

This is a simple edit tool for a text file. Change each line in the stream, using a javaScript snippet to rebuild these lines.

Please install Golang if you do not have Golang already installed.

* Add GOPATH/bin to your PATH:
    * for bash:
    `echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc && source ~/.bashrc`
    * for zsh:
    `echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc && source ~/.zshrc`
* Install line-editor:
    * Golang 1.16<=: `go get github.com/AntonStolov/line-editor`
    * Golang >=1.17: `go install github.com/AntonStolov/line-editor@latest`
## How to use:

* Select file using '--file' flag
* Write javascript code to change each line from the file using '--script' flag

## Example:
Replace commas in the CSV file to the pipelines '|' if the number of rows equals 4

`line-editor --file test.csv --script "String(input).split(',').length == 4 ? String(input).split(',').join('|') : ',,,'"`