package fizzbuzz

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
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
	tests := map[string]struct {
		firstN, lastN int
		out           io.Writer
		rules         Rules
		incrementer   Incrementer
		formatter     Formatter
		expOut        string
		expOutput     []string
		expErr        error
	}{
		"only_FizzBuzz_rule": {
			firstN: 1,
			lastN:  15,
			out:    &bytes.Buffer{},
			rules: Rules{
				func(i int) (string, bool) { return "FizzBuzz", i%3 == 0 || i%5 == 0 },
			},
			incrementer: defaultInc,
			formatter:   formatLine,
			expOut:      "1, 2, FizzBuzz, 4, FizzBuzz, FizzBuzz, 7, 8, FizzBuzz, FizzBuzz, 11, FizzBuzz, 13, 14, FizzBuzz.",
			expOutput:   []string{"1", "2", "FizzBuzz", "4", "FizzBuzz", "FizzBuzz", "7", "8", "FizzBuzz", "FizzBuzz", "11", "FizzBuzz", "13", "14", "FizzBuzz"},
			expErr:      nil,
		},
		"server_received": {
			firstN:      3,
			lastN:       3,
			out:         &bytes.Buffer{},
			rules:       rules,
			incrementer: defaultInc,
			formatter:   formatLine,
			expOut:      "Fizz.",
			expOutput:   []string{"Fizz"},
			expErr:      nil,
		},
		"basic_case": {
			firstN:      1,
			lastN:       15,
			out:         &bytes.Buffer{},
			rules:       rules,
			incrementer: defaultInc,
			formatter:   formatLine,
			expOut:      "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz.",
			expOutput:   []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
			expErr:      nil,
		},
		"line_by_line": {
			firstN:      1,
			lastN:       15,
			out:         &bytes.Buffer{},
			rules:       rules,
			incrementer: defaultInc,
			formatter:   formatLines,
			expOut:      "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n14\nFizzBuzz\n",
			expOutput:   []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
			expErr:      nil,
		},
		"fibonacci_case": {
			firstN:      1,
			lastN:       15,
			out:         &bytes.Buffer{},
			rules:       rules,
			incrementer: fibonacciInc,
			formatter:   formatFibonacciLine,
			expOut:      "1, 1, 2, Fizz, Buzz, 8, 13.",
			expOutput:   []string{"1", "2", "Fizz", "Buzz", "8", "13"},
			expErr:      nil,
		},
		"fibonacci_case_line_by_line": {
			firstN:      1,
			lastN:       15,
			out:         &bytes.Buffer{},
			rules:       rules,
			incrementer: fibonacciInc,
			formatter:   formatFibonacciLines,
			expOut:      "1\n1\n2\nFizz\nBuzz\n8\n13\n",
			expOutput:   []string{"1", "2", "Fizz", "Buzz", "8", "13"},
			expErr:      nil,
		},
		"bad_input": {
			firstN:      5,
			lastN:       3,
			out:         &bytes.Buffer{},
			rules:       rules,
			incrementer: defaultInc,
			formatter:   formatLine,
			expOut:      "",
			expOutput:   nil,
			expErr:      errBadInput,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := FizzBuzz(tt.firstN, tt.lastN, tt.out, tt.rules, tt.incrementer, tt.formatter)
			assert.Equal(t, tt.expOutput, actual)
			if tt.expErr != nil {
				assert.EqualError(t, err, tt.expErr.Error())
			} else {
				assert.NoError(t, err)
			}

			out := tt.out.(*bytes.Buffer).String()
			assert.Equal(t, tt.expOut, out)

			for i := range actual {
				assert.Equal(t, tt.expOutput[i], actual[i])
			}
		})
	}
}

func TestGetRuleFor(t *testing.T) {
	tests := map[string]struct {
		n     int
		rules Rules
		want  int
	}{
		"rule_fizz": {
			n:     42,
			rules: rules,
			want:  1,
		},
		"wrong_rule": {
			n:     43,
			rules: rules,
			want:  -1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			rule := GetRuleFor(tt.n, tt.rules)
			assert.Equal(t, rule, tt.want)
		})
	}

}
