package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {

	path := flag.String("path", ".", "absolute/path/to/target")
	flag.Parse()

	// Copy CMake files
	cmakeFiles := listFilesInDir("/Project-Files/CMakeFiles_proj")
	moveAllFilesInDir(cmakeFiles, *path)

	// Copy src dir
	srcFiles := listFilesInDir("/Project-Files/src")
	os.Mkdir(*path+"/src", 0755)
	moveAllFilesInDir(srcFiles, *path+"/src")
}

//
// Return a list of all the files in the directory
//
func listFilesInDir(dir string) []string {
	var files []string

	wd, osErr := os.Getwd()

	if osErr != nil {
		fmt.Print(osErr)
	}

	root := wd + "/" + dir + "/"
	err := filepath.Walk(root, func(_path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, _path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return files
}

func moveAllFilesInDir(files []string, path string) {
	for _, file := range files {
		move(path, file)
	}
}

//
// Copy a file to the destination
//
func move(path string, file string) {
	inputFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	outputFile, err := os.Create(path + "/" + filepath.Base(file))
	if err != nil {
		inputFile.Close()
		panic(err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		panic(err)
	}
}
