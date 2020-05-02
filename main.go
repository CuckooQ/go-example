package main

import (
	"fmt"

	"github.com/cuckooq/learngo/something"
)

func multiply(a int, b int) int {
	return a * b
}

func main() {
	const name1 string = "cuckoo"
	var name2 string = "cuckoo2"
	name3 := "cuckoo3"
	name2 = "cuckoo3"
	name3 = "cuckoo4"

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)
	something.SayBye()
	fmt.Println(multiply(2, 2))
}
