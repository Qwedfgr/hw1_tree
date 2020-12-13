package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func dirTree(out *os.File, path string, printFiles bool) (err error) {
	file, err := os.Open(path)
	files, _ := file.Readdir(-1)
	prefix := "├───"
	level := 0
	tabPrefix := strings.Repeat("\t", level) + prefix
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	for _, v := range files {
		size := ""
		if !v.IsDir() {
			if v.Size() == 0 {
				size = "(empty)"
			} else {
				size = fmt.Sprintf("(%db)", v.Size())
			}
		}
		fmt.Printf("%s %s %s \n", tabPrefix, v.Name(), size)
	}
	return nil
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
