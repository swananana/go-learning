package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishPrefix + name 
	}

	if language == "French" {
		return frenchPrefix + name
	}
	
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}