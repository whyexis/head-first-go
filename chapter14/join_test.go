package prose // part of the same package as the code to be tested

import "testing" // *required* to use a type from the testing package, import the standard library's testing package

// function name should being with "Test"
// name after test is up to user
// function will be passed a single parameter: a pointer to a testing.T value
func TestTwoElements(t *testing.T) {
	t.Error("No test written yet") // call a method on the testing.T to fail the test
}

func TestThreeElements(t *testing.T) {
	t.Error("no test here either") // most methods accep a string with a message explaining the reason the test failed
}
