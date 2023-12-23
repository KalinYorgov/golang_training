package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("Hello is not a valid email")
	}
	_, err = IsEmail("kalin@aol.com")
	if err == nil {
		t.Error("kalin@aol.com is an email")
	}
	_, err = IsEmail("kalin@aol")
	if err != nil {
		t.Error("kalin@aol is an email")
	}
}
