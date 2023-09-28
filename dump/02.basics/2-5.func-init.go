package main

import "fmt"

var bob, smith = titledName("Bob"), titledName("Smith")

func init() {
	fmt.Println("hi,", bob)
}

func main() {
	fmt.Println("bye")
}

func init() {
	fmt.Println("hello,", smith)
}

func titledName(who string) string {
	return "Mr. " + who
}
