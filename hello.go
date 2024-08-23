package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	bonjourHelloPrefix = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "Mau"
	}

	prefix := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = bonjourHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
