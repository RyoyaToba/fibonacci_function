package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_get_fibonacci(t *testing.T) {
	tests := map[string]struct {
		input  int
		output func(*testing.T, int)
	}{
		"case1": {
			input: 0,
			output: func(t *testing.T, answer int) {
				assert.Equal(t, 0, answer)
			},
		},
		"case2": {
			input: 1,
			output: func(t *testing.T, answer int) {
				assert.Equal(t, 1, answer)
			},
		},
		"case3": {
			input: 2,
			output: func(t *testing.T, answer int) {
				assert.Equal(t, 1, answer)
			},
		},
		"case4": {
			input: 7,
			output: func(t *testing.T, answer int) {
				assert.Equal(t, 13, answer)
			},
		},
		"case5": {
			input: 10,
			output: func(t *testing.T, answer int) {
				assert.Equal(t, 55, answer)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			input := tt.input
			answer := get_fibonacci(input)
			tt.output(t, answer)
		})
	}
}
