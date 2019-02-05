package goconfig

import (
	"reflect"
	"runtime/debug"
	"testing"
)

// AssertNil checks whether the obtained field is equal to nil,
// failing in other case.
func AssertNil(t *testing.T, obtained interface{}) {

	if nil == obtained {
		return
	}

	if reflect.ValueOf(obtained).IsNil() {
		return
	}

	line := GetStackLine(2)
	t.Errorf("\nExpected: nil \nObtained:%#v\nat %s\n", obtained, line)
}

// AssertNotNil checks whether the obtained field is distinct to nil,
// failing in other case.
func AssertNotNil(t *testing.T, obtained interface{}) {

	line := GetStackLine(2)
	if nil == obtained {
		t.Errorf("\nExpected: not nil \nObtained:%#v\nat %s\n", obtained, line)
		return
	}

	if reflect.ValueOf(obtained).IsNil() {
		t.Errorf("\nExpected: not nil \nObtained:%#v\nat %s\n", obtained, line)
		return
	}

}

// AssertEqual checks whether the obtained and expected fields are equal
// failing in other case.
func AssertEqual(t *testing.T, obtained, expected interface{}) bool {
	if reflect.DeepEqual(expected, obtained) {
		return true
	}

	line := GetStackLine(2)
	t.Errorf("\nExpected: %#v\nObtained: %#v\nat %s\n", expected, obtained, line)

	return false
}

// GetStackLine accesses the stack trace to get some lines
// so they can be showed by the tests in case of error.
func GetStackLine(linesToSkip int) string {

	stack := debug.Stack()
	lines := make([]string, 0)
	index := 0
	for i := 0; i < len(stack); i++ {
		if stack[i] == []byte("\n")[0] {
			lines = append(lines, string(stack[index:i-1]))
			index = i + 1
		}
	}
	return lines[linesToSkip*2+3] + "\n" + lines[linesToSkip*2+4]
}
