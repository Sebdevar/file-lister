package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NameHeader = "Name"
const SizeHeader = "Size (bytes)"
const ModTimeHeader = "Last Modified On"

var NameColumnLength = len(NameHeader)
var SizeColumnLength = len(SizeHeader)
var ModTimeColumnLength = len(ModTimeHeader)

func setDisplayLengths(fileSet []os.FileInfo) {
	var largestFileSize int
	for _, file := range fileSet {
		fileNameLength := len(file.Name())
		if file.IsDir() {
			fileNameLength++
		}
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

func addTableLine(name, size, modTime string) string {
	nameEntry := name + strings.Repeat(" ", NameColumnLength-len(name))
	sizeEntry := size + strings.Repeat(" ", SizeColumnLength-len(size))
	modTimeEntry := modTime + strings.Repeat(" ", ModTimeColumnLength-len(modTime))

	return fmt.Sprintf("= %s | %s | %s =\n", nameEntry, sizeEntry, modTimeEntry)
}

func addTableDivision() string {
	return fmt.Sprintln(strings.Repeat("=", NameColumnLength+SizeColumnLength+ModTimeColumnLength+10))
}

func addTotalStatistics(amountOfFiles, totalSizeInBytes int) (output string) {
	output = fmt.Sprintln("Amount of files: ", amountOfFiles)
	output += fmt.Sprintln("Total size (bytes): ", totalSizeInBytes)
	return
}

func getTableDisplayAsString(fileSet []os.FileInfo, addDirectorySizesToTotal bool) (output string) {
	setDisplayLengths(fileSet)

	output += addTableDivision()
	output += addTableLine(NameHeader, SizeHeader, ModTimeHeader)
	output += addTableDivision()

	var totalSize int
	for _, file := range fileSet {
		fileName := file.Name()
		fileSize := strconv.Itoa(int(file.Size()))
		fileModTime := file.ModTime().String()

		if file.IsDir() {
			fileName += "/"
			if addDirectorySizesToTotal {
				totalSize += int(file.Size())
			}
		} else {
			totalSize += int(file.Size())
		}

		output += addTableLine(fileName, fileSize, fileModTime)
	}

	output += addTableDivision()
	output += addTotalStatistics(len(fileSet), totalSize)
	return
}
