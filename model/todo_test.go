package model

import (
	"testing"
)

type printValuesTest struct {
	todo     Todo
	expected bool
}

var printValuesTests = []printValuesTest{
	{Todo{Id: 10, UserId: 101, Title: "The title", Completed: true}, true},
	{Todo{Id: 11, UserId: 111, Title: "The title", Completed: false}, true},
	{Todo{Id: 0, UserId: 121, Title: "The title", Completed: true}, false},
	{Todo{Id: 12, UserId: 131, Title: "", Completed: true}, false},
}

func TestGenerate(t *testing.T) {
	for _, test := range printValuesTests {
		output := test.todo.PrintValues()

		if output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
