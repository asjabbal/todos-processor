package internal

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

type readTest struct {
	typeArg, sizeArg string
	expected         map[string]interface{}
}

var readTests = []readTest{
	{"-type=even", "-size=100", map[string]interface{}{
		"type": "even",
		"size": 100,
	}},
	{"-type=odd", "-size=200", map[string]interface{}{
		"type": "odd",
		"size": 200,
	}},
	{"-type=all", "-size=300", map[string]interface{}{
		"type": "all",
		"size": 300,
	}},
	{"-type=even", "-size=1500", nil},
	{"-type=odd", "-size=1100", nil},
	{"-type=all", "-size=1001", nil},
	{"-type=prime", "-size=100", nil},
	{"-type=prime", "-size=2000", nil},
}

func TestRead(t *testing.T) {
	for _, test := range readTests {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		os.Args = []string{"cmd", test.typeArg, test.sizeArg}
		output, _ := ArgsReader{}.Read()

		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
