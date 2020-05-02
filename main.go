package main

import (
	"fmt"
	"strings"

	"github.com/cuckooq/learngo/something"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
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
	fmt.Println(multiply(3, 2))

	totalLengh, uppername := lenAndUpper("name")
	fmt.Println(totalLengh)
	fmt.Println(uppername)

	totalLenghRe, _ := lenAndUpper("test")
	fmt.Println(totalLenghRe)

	repeatMe("name", "name2")
}
