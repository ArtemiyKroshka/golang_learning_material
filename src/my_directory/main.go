package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func walkDir(path string) {
	files, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fullPath := filepath.Join(path, file.Name())
		if file.IsDir() {
			fmt.Println("Directory:", fullPath)
			walkDir(fullPath)
		} else {
			fmt.Println("File:", fullPath)
		}
	}
}

func main() {
	defer reportPanic()
	panic("some other issue")
	walkDir("my_directory")
}
