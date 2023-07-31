package fizzbuzz

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Rule takes an integer to check. If the integer satisfies the rule which Rule
// represents, it returns the string which the rule defines and true, otherwise
// it returns an empty string and false.
type Rule func(int) (string, bool)

// Rules is a slice in which an index is equal to a priority of Rule. The lower
// an index is, the higher the priority of Rule.
type Rules []Rule

// Incrementer takes the current integer and returns the new one.
type Incrementer func(int) int

// Formatter takes answer and returns it formatted. Also it takes isFirst and
// isLast arguments to make conditional formattings of the first and the last
// answers.
type Formatter func(answer string, isFirst, isLast bool) string

// GetRuleFor returns the index of Rule which is satisfied by n, otherwise it
// returns -1.
func GetRuleFor(n int, rules Rules) int {
	for i, rule := range rules {
		_, ok := rule(n)
		if ok {
			return i
		}
	}
	return -1
}

var errBadInput = errors.New("first N more then last N station")

// FizzBuzz reads the last number from in, then starting from 1 and until n
// with a step defined by the return value of increment it outputs strings
// checking if the current number satisfies any rule from rules. If it does, it
// uses the string returned from the rule, otherwise it returns the current
// number as a string. Then it formats the answer with format.
func FizzBuzz(
	firstN,
	lastN int,
	out io.Writer,
	rules Rules,
	increment Incrementer,
	format Formatter,
) ([]string, error) {
	var output []string

	if firstN > lastN {
		return nil, errBadInput
	}

	for i := firstN; i <= lastN; i = increment(i) {
		answer := strconv.Itoa(i)
		rule := GetRuleFor(i, rules)
		if rule > -1 {
			answer, _ = rules[rule](i)
		}
		output = append(output, answer)
		formattedAnswer := format(answer, i == 1, i == lastN || increment(i) > lastN)
		_, err := out.Write([]byte(formattedAnswer))
		if err != nil {
			return nil, fmt.Errorf("fizzBuzz: out.Write: %v", err)
		}
	}
	return output, nil
}
