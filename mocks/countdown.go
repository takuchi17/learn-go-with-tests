package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(buffer io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprint(buffer, i)
	}
	fmt.Fprintln(buffer, "Go!") // ln adds a newline
}

func main() {
	Countdown(os.Stdout)
}
