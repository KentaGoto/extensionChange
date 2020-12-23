package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func main() {
	var dir string
	var cExt string

	a := os.Args
	aN := len(a)

	if aN == 1 {
		fmt.Print("Dir: ")
		fmt.Scan(&dir)
		fmt.Print("Ext: ")
		fmt.Scan(&cExt)
	} else if aN == 3 {
		dir = os.Args[1]  // Source directory
		cExt = os.Args[2] // Extensions name to be changed
	} else {
		fmt.Println("The number of arguments specified is incorrect.")
		os.Exit(1)
	}

	paths := dirwalk(dir)

	for _, path := range paths {
		dirname := filepath.Dir(path)
		filebasename := filepath.Base(path)
		ext := filepath.Ext(path)

		// Replace file extensions name
		replacedFilebasename := strings.Replace(filebasename, ext, "."+cExt, 1)
		if err := os.Rename(path, dirname+"\\"+replacedFilebasename); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Done!")
}
