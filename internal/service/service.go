package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func IdentifyData(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", errors.New("Empty string")
	}

	for _, r := range s {
		if unicode.IsLetter(r) {
			result := morse.ToMorse(s)
			if strings.TrimSpace(result) == "" {
				return "", errors.New("Could not convert to morse")
			}
			return result, nil
		}
	}

	result := morse.ToText(s)
	if strings.TrimSpace(result) == "" {
		return "", errors.New("Could not convert to text")
	}
	return result, nil
}
