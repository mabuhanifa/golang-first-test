package main

import "fmt"

var print = fmt.Println

func main() {
	fmt.Println("Hello, World!")
	var people string = `checking string literal`

	print(people)

}
