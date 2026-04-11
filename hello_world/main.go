package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name, lang string) string {
	switch lang {
	case "Spanish":
		if name == "" {
			name = "Mundo"
		}
		return spanishHelloPrefix + name
	default:
		if name == "" {
			name = "World"
		}
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(Hello("Jack", "Spanish"))
}
