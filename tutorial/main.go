package main

import "fmt"

var print = fmt.Println

func main() {
	fmt.Println("Hello, World!")
	var people string = `checking string literal`
	name := "Alice"
	age := 30
	message := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	print(message, people)
}
