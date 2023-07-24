package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/vadim-gerasimov/codefinityTestTask/pkg/fizzbuzz"
)

// rules defines rules for printing Fizz, Buzz and FizzBuzz.
var rules = fizzbuzz.Rules{
	// Checks if the number must be replaced with FizzBuzz.
	func(i int) (string, bool) { return "FizzBuzz", i%3 == 0 && i%5 == 0 },
	// Checks if the number must be replaced with Fizz.
	func(i int) (string, bool) { return "Fizz", i%3 == 0 },
	// Checks if the number must be replaced with Buzz.
	func(i int) (string, bool) { return "Buzz", i%5 == 0 },
}

// inc takes an integer and returns the next one.
func inc(i int) int {
	return i + 1
}

// format takes output and returns output with a newline in the end.
func format(output string, _, _ bool) string {
	return output + "\n"
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(fmt.Errorf("input: io.ReadAll: %v", err))
	}
	n, err := strconv.Atoi(string(input))
	if err != nil {
		log.Fatal(fmt.Errorf("n: strconv.Atoi: %v", err))
	}
	_, err = fizzbuzz.FizzBuzz(1, n, os.Stdout, rules, inc, format)
	if err != nil {
		log.Fatal(err)
	}
}
