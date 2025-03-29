package main

import (
	"fmt"
)

const SPANISH = "Spanish"
const FRENCH = "French"

const ENGLISH_HELLO_PREFIX = "Hello, "
const SPANISH_HELLO_PREFIX = "Hola, "
const FRENCH_HELLO_PREFIX = "Bonjour, "

const ENDING_PREFIX = "!"

// 公開したくない関数は先頭を小文字にする。
func greetingPrefix(language string) (prefix string) {
	switch language {
	case FRENCH:
		prefix = FRENCH_HELLO_PREFIX
	case SPANISH:
		prefix = SPANISH_HELLO_PREFIX
	default:
		prefix = ENGLISH_HELLO_PREFIX
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name + ENDING_PREFIX
}

func main() {
	fmt.Println(Hello("Chris", "")) //副作用、文字列はドメイン。これを分離する。→　テスタビリティ
}
