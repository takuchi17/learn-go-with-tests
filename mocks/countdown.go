package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const countdownStart = 3
const countdownEnd = 0

func Countdown(buffer io.Writer) {
	for i := countdownStart; i > countdownEnd; i-- {
		fmt.Fprint(buffer, i)
	}
	fmt.Fprintln(buffer, finalWord) // ln adds a newline
}

func main() {
	Countdown(os.Stdout)
}
