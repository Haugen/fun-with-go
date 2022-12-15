package greetings

import (
	"errors"
	"unicode"
)

func ToUpper(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty string")
	}

	r := []rune(s)
	for i := range r {
		r[i] = unicode.ToUpper(r[i])
	}

	return string(r), nil
}
