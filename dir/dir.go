package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getFilePath(arg string) string {
	path, err := filepath.Abs(arg)
	check(err)
	removeFilePath(&path)
	return path
}

func removeFilePath(path *string) {
	fileInfo, err := os.Stat(*path)
	check(err)
	if fileInfo.Mode()&os.ModeType != os.ModeDir {
		*path = filepath.Dir(*path)
	}
}

func checkExists(path string) {
	_, err := os.Stat(path)
	check(err)
}

func main() {
	args := os.Args[1:]
	if length := len(args); length != 2 {
		panic("dir requires exactly 2 arguments")
	}
	origin := getFilePath(args[0])
	destination := getFilePath(args[1])
	if origin == destination {
		panic("directories must be different")
	}

	checkExists(origin)

	fmt.Println(origin)
	fmt.Println(destination)
}
