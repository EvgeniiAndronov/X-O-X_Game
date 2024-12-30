package server

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestMidlware_Test(t *testing.T) {
	testCases := []struct {
		name     string
		value    string
		expected string
	}{
		{
			name:     "Strange",
			value:    "___",
			expected: "948a13f1eb402f8ff16472e217608fb6",
		}, {
			name:     "Single",
			value:    "qwerty",
			expected: "d8578edf8458ce06fbc5bb76a58c5ca4",
		}, {
			name:     "only nums",
			value:    "321654987",
			expected: "d1b2cc725d846f0460ff290c60925070",
		}, {
			name:     "nums and literals",
			value:    "qwerty12356",
			expected: "865ac29d54fb192421d38b4e2a7440b4",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := makeHash(testCase.value)
			assert.Equal(t, testCase.expected, res)
		})
	}
}
