package hashchecker

import (
	"testing"
)

func TestBruteForceOnString(t *testing.T) {

	testStr := "test"
	brute := BruteForce("MD5", testStr, 100000, 4)

	if testStr == brute {
		t.Error("Success: The brute force result matches the test string.", testStr, brute)
	} else {
		t.Error("Failure: The brute force result does not match the test string.", testStr, brute)
	}

}
