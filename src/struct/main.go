package main

import (
	"fmt"
	"goCourse/src/struct/programmer"
)

func main() {
	zxc := programmer.NewProgrammer("Artem", 22, 2200, "Vue", programmer.Company{Name: "Holding House"})
	fmt.Println(zxc)

}
