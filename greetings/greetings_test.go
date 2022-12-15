package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "WhiteFluffy"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("WhiteFluffy")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("WhiteFluffy") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestToUpper(t *testing.T) {
	name := "WHITEFLUFFY"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := ToUpper("WhiteFluffy")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`ToUpper("WhiteFluffy") = %q, %v, want "WHITEFLUFFY", error`, msg, err)
	}
}
