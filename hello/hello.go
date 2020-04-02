package main

import "fmt"

func Hello() string {
	return "Hello, world"
}

func HelloYou(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello())
	fmt.Print(HelloYou("world"))
}
