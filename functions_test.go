package utils

import "testing"

func testCalled() string {
	return CallerFunctionName()
}
func testCaller() string {
	return testCalled()
}

func TestCallerFunctionName(t *testing.T) {
	caller := testCaller()
	if caller != "testCaller" {
		t.Fatalf("Name should be: caller, got: %s", caller)
	}
}
