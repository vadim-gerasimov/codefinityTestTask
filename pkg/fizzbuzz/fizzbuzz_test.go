package fizzbuzz

import (
	"bytes"
	"io"
	"strconv"
	"testing"
)

var rules = Rules{
	func(i int) (string, bool) { return "FizzBuzz", i%3 == 0 && i%5 == 0 },
	func(i int) (string, bool) { return "Fizz", i%3 == 0 },
	func(i int) (string, bool) { return "Buzz", i%5 == 0 },
}

func defaultInc(i int) int {
	return i + 1
}

func formatLine(output string, _, isLast bool) string {
	if isLast {
		return output + "."
	}
	return output + ", "
}

func formatLines(output string, _, _ bool) string {
	return output + "\n"
}

func fibonacciInc(i int) int {
	a, b := 0, 1
	for b <= i {
		a, b = b, a+b
	}
	return b
}

func formatFibonacciLine(output string, isFirst, isLast bool) string {
	formatted := output + ", "
	if isFirst {
		return "1, " + formatted
	}
	if isLast {
		return output + "."
	}
	return formatted
}

func formatFibonacciLines(output string, isFirst, _ bool) string {
	formatted := output + "\n"
	if isFirst {
		return "1\n" + formatted
	}
	return formatted
}

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		firstN, lastN int
		out           io.Writer
		rules         Rules
		incrementer   Incrementer
		formatter     Formatter
		wantOut       string
		wantOutput    []string
	}{
		{
			firstN: 1,
			lastN:  15,
			out:    &bytes.Buffer{},
			rules: Rules{
				func(i int) (string, bool) { return "FizzBuzz", i%3 == 0 || i%5 == 0 },
			},
			incrementer: defaultInc,
			formatter:   formatLine,
			wantOut:     "1, 2, FizzBuzz, 4, FizzBuzz, FizzBuzz, 7, 8, FizzBuzz, FizzBuzz, 11, FizzBuzz, 13, 14, FizzBuzz.",
			wantOutput:  []string{"1", "2", "FizzBuzz", "4", "FizzBuzz", "FizzBuzz", "7", "8", "FizzBuzz", "FizzBuzz", "11", "FizzBuzz", "13", "14", "FizzBuzz"},
		},
		{
			firstN:      3,
			lastN:       3,
			out:         &bytes.Buffer{},
			rules:       rules,
			incrementer: defaultInc,
			formatter:   formatLine,
			wantOut:     "Fizz.",
			wantOutput:  []string{"Fizz"},
		},
		{
			firstN:      1,
			lastN:       15,
			out:         &bytes.Buffer{},
			rules:       rules,
			incrementer: defaultInc,
			formatter:   formatLine,
			wantOut:     "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz.",
			wantOutput:  []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
		{
			firstN:      1,
			lastN:       15,
			out:         &bytes.Buffer{},
			rules:       rules,
			incrementer: defaultInc,
			formatter:   formatLines,
			wantOut:     "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n14\nFizzBuzz\n",
			wantOutput:  []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
		{
			firstN:      1,
			lastN:       15,
			out:         &bytes.Buffer{},
			rules:       rules,
			incrementer: fibonacciInc,
			formatter:   formatFibonacciLine,
			wantOut:     "1, 1, 2, Fizz, Buzz, 8, 13.",
			wantOutput:  []string{"1", "2", "Fizz", "Buzz", "8", "13"},
		},
		{
			firstN:      1,
			lastN:       15,
			out:         &bytes.Buffer{},
			rules:       rules,
			incrementer: fibonacciInc,
			formatter:   formatFibonacciLines,
			wantOut:     "1\n1\n2\nFizz\nBuzz\n8\n13\n",
			wantOutput:  []string{"1", "2", "Fizz", "Buzz", "8", "13"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.wantOut, func(t *testing.T) {
			output, err := FizzBuzz(tt.firstN, tt.lastN, tt.out, tt.rules, tt.incrementer, tt.formatter)
			if err != nil {
				t.Errorf("TestFizzBuzz: FizzBuzz: %v", err)
			}

			out := tt.out.(*bytes.Buffer).String()
			if out != tt.wantOut {
				t.Errorf("TestFizzBuzz: want %s, got %s", tt.wantOut, out)
			}

			for i := range output {
				if output[i] != tt.wantOutput[i] {
					t.Errorf("TestFizzBuzz: want %s, got %s", tt.wantOutput[i], output[i])
				}
			}
		})
	}
}

func TestGetRuleFor(t *testing.T) {
	tests := []struct {
		n     int
		rules Rules
		want  int
	}{
		{
			n:     42,
			rules: rules,
			want:  1,
		},
		{
			n:     43,
			rules: rules,
			want:  -1,
		},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			rule := GetRuleFor(tt.n, tt.rules)
			if rule != tt.want {
				t.Errorf("TestGetRuleFor: want %d, got %d", tt.want, rule)
			}
		})
	}
}
