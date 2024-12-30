package server

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func Test_logickChecker(t *testing.T) {
	testCases := []struct {
		name     string
		gameMap  [3][3]string
		expected string
	}{
		{
			name: "Empty map",
			gameMap: [3][3]string{
				{"", "", ""},
				{"", "", ""},
				{"", "", ""},
			},
			expected: "",
		}, {
			name: "Top line",
			gameMap: [3][3]string{
				{"X", "X", "X"},
				{"", "O", ""},
				{"O", "", "O"},
			},
			expected: "X",
		}, {
			name: "Diagonal L-Up to R-Down",
			gameMap: [3][3]string{
				{"O", "", ""},
				{"X", "O", "X"},
				{"", "X", "O"},
			},
			expected: "O",
		}, {
			name: "Midle vertical",
			gameMap: [3][3]string{
				{"O", "X", ""},
				{"", "X", "O"},
				{"", "X", "O"},
			},
			expected: "X",
		}, {
			name: "Diagonal R-Up to L-Down",
			gameMap: [3][3]string{
				{"", "X", "O"},
				{"", "O", ""},
				{"O", "X", "X"},
			},
			expected: "O",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Checker(tc.gameMap)
			assert.Equal(t, tc.expected, result)
		})
	}
}
