package main

import "fmt"

var print = fmt.Println

func main() {
	fmt.Println("Hello, World!")
	var people = `checking string literal`

	print(people)

}
