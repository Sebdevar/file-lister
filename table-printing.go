package main

import (
	"fmt"
	"io"
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

func printTableLine(outputStream io.Writer, name, size, modTime string) (err error) {
	nameEntry := name + strings.Repeat(" ", NameColumnLength-len(name))
	sizeEntry := size + strings.Repeat(" ", SizeColumnLength-len(size))
	modTimeEntry := modTime + strings.Repeat(" ", ModTimeColumnLength-len(modTime))

	_, err = fmt.Fprintf(outputStream, "= %s | %s | %s =\n", nameEntry, sizeEntry, modTimeEntry)
	return
}

func printTableDivision(outputStream io.Writer) (err error) {
	_, err = fmt.Fprintln(outputStream, strings.Repeat("=", NameColumnLength+SizeColumnLength+ModTimeColumnLength+10))
	return
}

func printTotalStatistics(outputStream io.Writer, amountOfFiles, totalSizeInBytes int) (err error) {
	_, err = fmt.Fprintln(outputStream, "Amount of files: ", amountOfFiles)
	if err != nil {
		return
	}

	_, err = fmt.Fprintln(outputStream, "Total size (bytes): ", totalSizeInBytes)
	return
}

func printTable(outputStream io.Writer, fileSet []os.FileInfo, addDirectorySizesToTotal bool) (err error) {
	setDisplayLengths(fileSet)

	err = printTableDivision(outputStream)
	if err != nil {
		return
	}
	err = printTableLine(outputStream, NameHeader, SizeHeader, ModTimeHeader)
	if err != nil {
		return
	}
	err = printTableDivision(outputStream)
	if err != nil {
		return
	}

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

		err = printTableLine(outputStream, fileName, fileSize, fileModTime)
		if err != nil {
			return
		}
	}

	err = printTableDivision(outputStream)
	if err != nil {
		return
	}
	err = printTotalStatistics(outputStream, len(fileSet), totalSize)
	return
}
