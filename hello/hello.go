package main

import (
	"fmt"
)

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const ending = "!"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name + ending
	}
	return englishHelloPrefix + name + ending
}

func main() {
	fmt.Println(Hello("Chris", "")) //副作用、文字列はドメイン。これを分離する。→　テスタビリティ
}
