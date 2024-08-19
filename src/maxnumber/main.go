package main

import "fmt"

func maximun(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]

	for _, val := range numbers {
		if max < val {
			max = val
		}
	}

	return max
}

func main() {
	fmt.Println(maximun(1, 5, 2, 3, 11, 7, 3))
}
