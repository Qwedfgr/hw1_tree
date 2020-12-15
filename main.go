package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"sort"
)

func dirTree(out io.Writer, path string, printFiles bool) (err error) {
	recursiveDir(out, path, printFiles, "")
	return nil
}

func getPrintInfo(out io.Writer, v os.FileInfo, arrayLen int, isLastPosition bool, prefix string) {
	usualPrefix := "├───"
	lastPositionPrefix := "└───"
	size := ""
	tabPrefix := ""

	if isLastPosition {
		tabPrefix = prefix + lastPositionPrefix
	} else {
		tabPrefix = prefix + usualPrefix
	}

	if !v.IsDir() {
		if v.Size() == 0 {
			size = " (empty)"
		} else {
			size = fmt.Sprintf(" (%db)", v.Size())
		}
	}
	fmt.Fprintf(out, "%s%s%s\n", tabPrefix, v.Name(), size)
}

func recursiveDir(out io.Writer, pathDir string, printFiles bool, prefix string) {
	file, _ := os.Open(pathDir)
	files, _ := file.Readdir(-1)
	newPrefix := ""
	var newFiles []os.FileInfo
	file.Close()
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	if !printFiles {
		for i := 0; i < len(files); i++ {
			if files[i].IsDir() {
				newFiles = append(newFiles, files[i])
			}
		}
		files = newFiles
	}

	for i := 0; i < len(files); i++ {
		isLastPosition := i+1 == len(files)
		if isLastPosition {
			newPrefix = "\t"
		} else {
			newPrefix = "│\t"
		}
		if files[i].IsDir() {
			getPrintInfo(out, files[i], len(files), isLastPosition, prefix)
			recursiveDir(out, path.Join(pathDir, files[i].Name()), printFiles, prefix+newPrefix)
		} else if printFiles {
			getPrintInfo(out, files[i], len(files), isLastPosition, prefix)
		}
	}
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
