package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

const NameHeader = "Name"
const SizeHeader = "Size (bytes)"
const ModTimeHeader = "Last Modified On"

var NameColumnLength = len(NameHeader)
var SizeColumnLength = len(SizeHeader)
var ModTimeColumnLength = len(ModTimeHeader)

func (cli *Cli) setDisplayLengths(fileSet []os.FileInfo) {
	var largestFileSize int
	for _, file := range fileSet {
		fileNameLength := len(file.Name())
		fileSize := int(file.Size())
		fileModTimeLength := len(file.ModTime().String())

		if fileNameLength > NameColumnLength {
			NameColumnLength = fileNameLength
		}
		if fileSize > largestFileSize {
			largestFileSize = fileSize
		}
		if fileModTimeLength > ModTimeColumnLength {
			ModTimeColumnLength = fileModTimeLength
		}
	}

	if lengthOfLargestFileSize := len(strconv.Itoa(largestFileSize)); lengthOfLargestFileSize > SizeColumnLength {
		SizeColumnLength = lengthOfLargestFileSize
	}
}

func (cli *Cli) printTableLine(name, size, modTime string) {
	nameEntry := name + strings.Repeat(" ", NameColumnLength-len(name))
	sizeEntry := size + strings.Repeat(" ", SizeColumnLength-len(size))
	modTimeEntry := modTime + strings.Repeat(" ", ModTimeColumnLength-len(modTime))

	fmt.Printf("= %s | %s | %s =\n", nameEntry, sizeEntry, modTimeEntry)
}

func (cli *Cli) printTableDivision() {
	fmt.Println(strings.Repeat("=", NameColumnLength+SizeColumnLength+ModTimeColumnLength+10))
}

func (cli *Cli) printTotalStatistics(amountOfFiles, totalSizeInBytes int) {
	fmt.Println("Amount of files: ", amountOfFiles)
	fmt.Println("Total size (bytes): ", totalSizeInBytes)
}

func (cli *Cli) printTable(fileSet []os.FileInfo) {
	cli.setDisplayLengths(fileSet)

	cli.printTableDivision()
	cli.printTableLine(NameHeader, SizeHeader, ModTimeHeader)
	cli.printTableDivision()

	var totalSize int
	for _, file := range fileSet {
		fileName := file.Name()
		fileSize := strconv.Itoa(int(file.Size()))
		fileModTime := file.ModTime().String()

		if file.IsDir() {
			fileName += "/"
			if cli.addDirectorySizesToTotal {
				totalSize += int(file.Size())
			}
		} else {
			totalSize += int(file.Size())
		}

		cli.printTableLine(fileName, fileSize, fileModTime)
	}

	cli.printTableDivision()
	cli.printTotalStatistics(len(fileSet), totalSize)
}
