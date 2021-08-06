package main

import "fmt"

var englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(Hello("Puneet"))
}
