package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if len(str) <= 1 {
		return str, nil
	}

	var sb strings.Builder
	runes := []rune(str)
	var rune1 rune
	var rune2 rune

	for i := 1; i < len(runes); {
		rune1 = runes[i-1]
		rune2 = runes[i]

		// Тут был if-else, но линтер заставил поменять его на switch
		// Имхо, выглядеть стало хуже
		switch {
		case unicode.IsDigit(rune2):
			digit, _ := strconv.Atoi(string(rune2))
			if unicode.IsDigit(rune1) {
				return "", ErrInvalidString
			}
			sb.WriteString(strings.Repeat(string(rune1), digit))
			i += 2
		case !unicode.IsDigit(rune1):
			sb.WriteRune(rune1)
			i++
		default:
			return "", ErrInvalidString
		}
	}

	lastRune := runes[len(runes)-1]

	if !unicode.IsDigit(lastRune) {
		sb.WriteRune(lastRune)
	} else if unicode.IsDigit(runes[len(runes)-2]) {
		return "", ErrInvalidString
	}

	return sb.String(), nil
}
