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

func main() {
	args := os.Args[1:]
	if length := len(args); length != 2 {
		panic("dir requires exactly 2 arguments")
	}
	origin, err := filepath.Abs(args[0])
	check(err)
	destination, _ := filepath.Abs(args[1])
	check(err)
	fmt.Println(origin)
	fmt.Println(destination)
}
