package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3
const countdownEnd = 0

func Countdown(buffer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > countdownEnd; i-- {
		fmt.Fprintln(buffer, i) // ln adds a newline
		sleeper.Sleep()
	}
	fmt.Fprint(buffer, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
