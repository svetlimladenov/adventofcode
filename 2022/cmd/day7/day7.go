package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var root Directory
var currentDir *Directory

func main() {
	root = *NewDirectory("/")

	currentDir = &root

	lines := parseInput(input)
	for _, l := range lines {
		lineParts := parseLine(l)
		if lineParts[0] == "$" {
			processCmd(lineParts)
		} else if lineParts[0] == "dir" {
			fmt.Println("Directory :", lineParts[1])
			currentDir.SubDirectories[lineParts[1]] = NewDirectory(lineParts[1])
			currentDir.SubDirectories[lineParts[1]].ParentDirectory = currentDir
			// Directory
		} else {
			// File
			file := processFile(lineParts)
			currentDir.Files = append(currentDir.Files, *file)
		}
	}

	fmt.Println(root)
}

func processCmd(cmd []string) {
	if cmd[1] == "cd" {
		newPath := cmd[2]

		fmt.Printf("Changing directory from %s to %s\n", currentDir.Path, newPath)

		if newPath == "/" {
			// go to root
			currentDir = &root
			return
		}

		if newPath == ".." {
			// go to parent
			currentDir = currentDir.ParentDirectory
		} else {
			// go to child directory
			newDir := currentDir.SubDirectories[newPath]
			newDir.ParentDirectory = currentDir // ParentDir points to currentDir
			currentDir = newDir
		}
	}
}

func processFile(lineParts []string) *File {
	fileSize, err := strconv.ParseInt(lineParts[0], 10, 64)
	if err != nil {
		panic("Cannot parse fize size")
	}

	return &File{Size: fileSize, Name: lineParts[1]}
}

func parseLine(line string) []string {
	return strings.Split(line, " ")
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

type File struct {
	Name string
	Size int64
}

type Directory struct {
	Path            string
	SubDirectories  map[string]*Directory
	ParentDirectory *Directory
	Files           []File
}

func NewDirectory(path string) *Directory {
	return &Directory{
		Path:           path,
		SubDirectories: map[string]*Directory{},
		Files:          make([]File, 0),
	}
}

// {
// 	"/": {
// 		"a": {
// 			"a": 1234
// 		},
// 		"b" {
// 			"b": 1234
// 		}
// 	}
// }
