package utils

import (
	"errors"
	"os"
	"strings"
	"testing"
)

// testError returns an error
func testError() error {
	var err error
	_, err = os.Open("some not existing file")
	return WrapError(err)
}

//
func TestWrapError(t *testing.T) {
	err := testError()
	if err == nil {
		t.Fatal("error is nil it will never happen but I'm making Goland happy")
	}
	errFunc := strings.Split(err.Error(), ":")[0]
	if errFunc != "testError" {
		t.Fatalf("context func should be testError, got: `%s`", errFunc)
	}
}

func testErrorNamedReturnValue() (err error) {
	defer WrapErrorP(&err)
	err = errors.New("some error")
	return err
}

func TestWrapErrorP(t *testing.T) {
	err := testErrorNamedReturnValue()
	if err == nil {
		t.Fatal("error is nil it will never happen but I'm making Goland happy")
	}
	errFunc := strings.Split(err.Error(), ":")[0]
	if errFunc != "testErrorNamedReturnValue" {
		t.Fatalf("context func should be testError, got: `%s`", errFunc)
	}
}
