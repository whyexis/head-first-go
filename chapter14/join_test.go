package prose // part of the same package as the code to be tested

import "testing" // *required* to use a type from the testing package, import the standard library's testing package

// function name should being with "Test"
// name after test is up to user
// function will be passed a single parameter: a pointer to a testing.T value
func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want) // call a method on the testing.T to fail the test
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want) // call a method on the testing.T to fail the test
	}
}
