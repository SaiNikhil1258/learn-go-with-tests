package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bounjour, "
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Nikhil", "Spanish"))
}
