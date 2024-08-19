package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var sum int

	for _, val := range os.Args[1:] {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}

	fmt.Printf("Sum of numbers: %d", sum)

}
