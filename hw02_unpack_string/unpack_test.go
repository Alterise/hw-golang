package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "", expected: ""},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", "aaa10"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestUnpackUnicode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "привет", expected: "привет"},
		{input: "при3ве4т5", expected: "прииивееееттттт"},
		{input: "㌚㌓", expected: "㌚㌓"},
		{input: "㌚㌚㌓", expected: "㌚㌚㌓"},
		{input: "㌚㌚3㌓", expected: "㌚㌚㌚㌚㌓"},
		{input: "🤪💀🐳", expected: "🤪💀🐳"},
		{input: "🤪1💀5🐳0", expected: "🤪💀💀💀💀💀"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}
