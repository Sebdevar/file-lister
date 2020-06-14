package main

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func getFileSetFromFolder(folderLocation string) (fileSet []os.FileInfo) {
	fileSet, err := ioutil.ReadDir(folderLocation)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func sortFileSetBySize(fileSet []os.FileInfo) {
	sort.SliceStable(fileSet, func(i, j int) bool {
		if fileSet[i].IsDir() {
			if fileSet[j].IsDir() {
				return fileSet[i].Name() < fileSet[j].Name()
			}
			return true
		}
		return !fileSet[j].IsDir() && fileSet[i].Size() > fileSet[j].Size()
	})
}
