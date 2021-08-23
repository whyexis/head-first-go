package prose // part of the same package as the code to be tested

import "testing" // *required* to use a type from the testing package, import the standard library's testing package

// function name should being with "Test"
// name after test is up to user
// function will be passed a single parameter: a pointer to a testing.T value
func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	if JoinWithCommas(list) != "apple and orange" {
		t.Error("didn't match expected value") // call a method on the testing.T to fail the test
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	if JoinWithCommas(list) != "apple, orange, and pear" {
		t.Error("didn't match expected value") // most methods accep a string with a message explaining the reason the test failed
	}
}
