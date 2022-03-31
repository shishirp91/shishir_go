package main

import (
	"fmt"
)

func Age(name string, age int) {
	fmt.Printf("%s is %d years old.\n", name, age)
}

func Returnsports() string {
	return "volleyball"
}

func main() {
	Returnsports()
}
