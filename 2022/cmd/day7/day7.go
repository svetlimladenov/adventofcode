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
var smallestToDelete int64 = 70000000

func main() {
	root = *NewDirectory("/")

	currentDir = &root

	lines := parseInput(input)
	for _, l := range lines {
		lineParts := parseLine(l)
		if lineParts[0] == "$" {
			processCmd(lineParts)
		} else if lineParts[0] == "dir" {
			currentDir.SubDirectories[lineParts[1]] = NewDirectory(lineParts[1])
			currentDir.SubDirectories[lineParts[1]].ParentDirectory = currentDir
			// Directory
		} else {
			// File
			file := processFile(lineParts)
			currentDir.Files = append(currentDir.Files, *file)
		}
	}

	usedSpace := root.Sum()
	fmt.Printf("usedSpace: %v\n", usedSpace)

	freeSpace := 70000000 - usedSpace

	neededForUpdate := 30000000 - freeSpace

	fmt.Printf("neededForUpdate: %v\n", neededForUpdate)

	root.ClosestTo(neededForUpdate)
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

var result int64

func (d *Directory) Sum() int64 {
	var sum int64
	sum = 0
	for _, f := range d.Files {
		sum += f.Size
	}

	for _, sd := range d.SubDirectories {
		sdSum := sd.Sum()
		sum += sdSum
	}

	if sum < 100000 {
		result += sum
		fmt.Printf("Dir sum: %d\n", sum)
	}

	return sum
}

func (d *Directory) ClosestTo(spaceNeeded int64) int64 {
	var sum int64
	sum = 0
	for _, f := range d.Files {
		sum += f.Size
	}

	for _, sd := range d.SubDirectories {
		sdSum := sd.ClosestTo(spaceNeeded)
		sum += sdSum
	}

	fmt.Printf("Dir sum %s: %d\n", d.Path, sum)

	if sum >= spaceNeeded {
		if smallestToDelete > sum {
			fmt.Printf("ok to delete %q - %d\n", d.Path, sum)
			smallestToDelete = sum
		}
	}

	return sum
}
