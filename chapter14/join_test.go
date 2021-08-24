package prose // part of the same package as the code to be tested

import (
	"fmt"
	"testing" // *required* to use a type from the testing package, import the standard library's testing package
)

type testData struct {
	list []string
	want string
}

// function name should being with "Test"
// name after test is up to user
// function will be passed a single parameter: a pointer to a testing.T value
func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		testData{list: []string{"apple"}, want: "apple"},
		testData{list: []string{"apple", "orange"}, want: "apple and orange"},
		testData{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Error(errorString(test.list, got, test.want)) // call a method on the testing.T to fail the test
		}
	}
}

// errorString generates the error message if the test fails
func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}