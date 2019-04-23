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

type MockSleeper struct {
	Calls int
}

func (s *MockSleeper) Sleep() {
	s.Calls++
}

type RealSleeper struct {}

func (d *RealSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const countdownStart = 3
const final = "Go!"

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, final)
}

func main() {
	sleeper := &RealSleeper{}
	Countdown(os.Stdout, sleeper)
}