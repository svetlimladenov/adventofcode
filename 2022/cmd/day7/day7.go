package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var root *Directory

func main() {
	root = NewDirectory("/")

	currentDirectory := *root

	lines := parseInput(input)
	for _, l := range lines {
		lineParts := parseLine(l)
		if lineParts[0] == "$" {
			processCmd(lineParts, &currentDirectory)
		} else if lineParts[0] == "dir" {
			fmt.Println("Directory :", lineParts[1])
			currentDirectory.SubDirectories[lineParts[1]] = NewDirectory(lineParts[1])
			currentDirectory.SubDirectories[lineParts[1]].ParentDirectory = &currentDirectory
			// Directory
		} else {
			// File
			file := processFile(lineParts)
			currentDirectory.Files = append(currentDirectory.Files, *file)
		}
	}
}

func processCmd(cmd []string, dir *Directory) {
	if cmd[1] == "cd" {
		newPath := cmd[2]

		fmt.Printf("Changing directory from %s to %s\n", dir.Path, newPath)

		if newPath == "/" {
			// go to root
			*dir = *root
		} else if newPath == ".." {
			// go to parent
			*dir = *dir.ParentDirectory
		} else {
			// go to child directory
			newDir := *dir.SubDirectories[newPath]
			newDir.ParentDirectory = dir // ParentDir points to currentDir
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
