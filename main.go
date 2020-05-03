package main

import (
	"fmt"
	"strings"

	"github.com/cuckooq/learngo/something"
)

// function
func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpperNaked(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func supperAdd(numbers ...int) {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
}

func supperAdd2(numbers ...int) {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
}

func supperAdd3(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// if
func canIDrinkIf1(age int) bool {
	if age >= 20 {
		return true
	}

	return false
}

func canIDrinkIf2(age int) bool {
	if koreanAge := age - 2; koreanAge >= 18 {
		return true
	}

	return false
}

// switch
func canIDrinkSwitch1(age int) bool {
	switch {
	case age > 18:
		return true
	case age == 18:
		return true
	case age < 18:
		return false
	}

	return false
}

func canIDrinkSwitch2(age int) bool {
	switch age {
	case 18:
		return true
	case 20:
		return true
	case 50:
		return false
	}

	return false
}

func canIDrinkSwitch3(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 18:
		return true
	case 20:
		return true
	case 50:
		return false
	}

	return false
}

// structs
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// variables, constants
	const name1 string = "cuckoo"
	var name2 string = "cuckoo2"
	name3 := "cuckoo3"
	name2 = "cuckoo3"
	name3 = "cuckoo4"

	// function
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

	length, uppercase := lenAndUpperNaked("name")
	fmt.Println(length)
	fmt.Println(uppercase)

	// for
	supperAdd(1, 2, 3, 4)
	supperAdd2(1, 2, 3, 4)
	total := supperAdd3(1, 2, 3, 4)
	fmt.Println(total)

	// if
	fmt.Println(canIDrinkIf1((20)))
	fmt.Println(canIDrinkIf2((19)))

	// switch
	fmt.Println(canIDrinkSwitch1(20))
	fmt.Println(canIDrinkSwitch2(20))
	fmt.Println(canIDrinkSwitch3(20))

	// pointer
	a := 2
	b := a
	a = 10
	fmt.Println(a, b)
	fmt.Println(&a, &b)
	c := &a
	fmt.Println(c)
	fmt.Println(*c)
	*c = 2222
	fmt.Println(a)

	// arrays and slices
	names := [5]string{"cuckoo", "q", "cho"}
	fmt.Println(names)
	names[3] = "kemy"
	names[4] = "senz"
	fmt.Println(names)

	names2 := []string{"cuckoo", "q", "cho"}
	fmt.Println(names2)
	names2 = append(names2, "kemy")
	names2 = append(names2, "senz")
	fmt.Println(names2)

	// map
	namesMap := map[string]string{
		"KEY1": "cuckoo",
		"KEY2": "cho",
	}
	for key, value := range namesMap {
		fmt.Println(key, value)
	}

	// struct
	favFood := []string{"mandu", "kimchi"}
	cuckoo := person{
		name:    "cuckoo",
		age:     31,
		favFood: favFood,
	}
	fmt.Println(cuckoo)
}
