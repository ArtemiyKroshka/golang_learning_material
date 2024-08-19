package main

import (
	"fmt"
	"goCourse/src/getaverage/getstring"
	"log"
	"path/filepath"
	"runtime"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)

	// Формируем полный путь к соседнему файлу "data.txt"
	dataFilePath := filepath.Join(path, "data.txt")
	data, err := getstring.GetStrings(dataFilePath)
	hash := make(map[string]int)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range data {
		hash[v]++
	}

	fmt.Println(hash)
}
