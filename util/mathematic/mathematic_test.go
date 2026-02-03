package mathematic_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/pobyzaarif/belajar-go-cli/util/mathematic"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	testCases := map[string]int{
		"1,2,3,4,5": 15,
		"1,2,3":     6,
		"1":         1,
	}

	for question, answer := range testCases {
		testSplit := strings.Split(question, ",")
		listNumber := make([]int, 0)

		for _, v := range testSplit {
			number, _ := strconv.Atoi(v)
			listNumber = append(listNumber, number)
		}

		result := mathematic.Sum(listNumber...)

		if result != answer {
			t.Fatalf("wrong with question: %v, expectation: %v result: %v", question, answer, result)
		}
	}
}

func TestSumSecond(t *testing.T) {
	testCases := map[string]int{
		"1,2,3,4,5": 15,
		"1,2,3":     6,
		"1":         1,
	}

	for question, answer := range testCases {
		testSplit := strings.Split(question, ",")
		listNumber := make([]int, 0)

		for _, v := range testSplit {
			number, _ := strconv.Atoi(v)
			listNumber = append(listNumber, number)
		}

		result := mathematic.Sum(listNumber...)

		assert.Equal(t, answer, result, fmt.Sprintf("Result has to be %v", answer))
	}
}
