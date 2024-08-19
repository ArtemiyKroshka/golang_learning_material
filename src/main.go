package main

import (
	"fmt"
	"goCourse/src/prose"
)

func main() {
	phrases := []string{"me", "marina", "andrii"}
	result := prose.JoinWithCommas(phrases)
	fmt.Println(result)
}
