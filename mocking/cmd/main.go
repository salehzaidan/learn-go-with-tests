package main

import (
	"learn-go-with-tests/mocking"
	"os"
	"time"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, SleepFn: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
