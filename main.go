package main

func main() {
	cli := Cli{}
	cli.init()
	if cli.directoryLocation != "" {
		files := getFileSetFromFolder(cli.directoryLocation)
		sortFileSetBySize(files)
		cli.printTable(files)
	}
}
