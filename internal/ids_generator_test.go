package internal

import (
	"reflect"
	"testing"
)

type generateTest struct {
	todosType     string
	numberOfTodos int
	expected      []int
}

var generateTests = []generateTest{
	{"even", 10, []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}},
	{"odd", 5, []int{1, 3, 5, 7, 9}},
	{"all", 15, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}},
	{"even", 1100, nil},
	{"odd", 1200, nil},
	{"all", 1001, nil},
	{"prime", 10, nil},
	{"prime", 1001, nil},
}

func TestGenerate(t *testing.T) {
	for _, test := range generateTests {
		output, _ := IdsGenerator{}.Generate(test.todosType, test.numberOfTodos)

		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
