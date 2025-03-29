package main

import (
	"fmt"
)

func Hello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(Hello("chris")) //副作用、文字列はドメイン。これを分離する。→　テスタビリティ
}
