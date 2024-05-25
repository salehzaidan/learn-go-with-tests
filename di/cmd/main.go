package main

import (
	"learn-go-with-tests/di"
	"os"
)

func main() {
	di.Greet(os.Stdout, "salehzaidan")
}
