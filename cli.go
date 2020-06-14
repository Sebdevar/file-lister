package main

import (
	"flag"
)

type Cli struct {
	addDirectorySizesToTotal bool
	directoryLocation        string
}

func (cli *Cli) init() {
	flag.BoolVar(&cli.addDirectorySizesToTotal, "dir-size", false, "Add the system size of directories to the calculated total (Default: false)")
	flag.Parse()
	cli.directoryLocation = flag.Arg(0)
}
