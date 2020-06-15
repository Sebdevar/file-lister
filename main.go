package main

import (
	"fmt"
)

func main() {
	cli := Cli{}
	cli.init()
	if cli.directoryLocation != "" {
		files := getFileSetFromFolder(cli.directoryLocation)
		sortFileSetBySize(files)
		fmt.Println(getTableDisplayAsString(files, cli.addDirectorySizesToTotal))
	}
}
