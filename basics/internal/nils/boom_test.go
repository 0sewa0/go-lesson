package nils

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func isEven(number int) bool {
	return number % 2 == 0
}

func TestIsEven(t *testing.T) {
	type TestCase struct {
		desc string
		input int
		expected bool
	}
	testCases := []TestCase{
		{
			desc: "Is 1 even?",
			input: 1,
			expected: false,
		},
		{
			desc: "Is 2 even?",
			input: 2,
			expected: true,
		},
		{
			desc: "3 is even",
			input: 3,
			expected: false,
		},
		{
			desc: "4 is even",
			input: 4,
			expected: true,
		},
		{
			desc: "1232143254256 even",
			input: 1232143254256,
			expected: true,
		},
		{
			desc: "0 goes boom",
			input: 0,
			expected: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			require.NotEqual(t, 0, test.input)
			assert.Equal(t, test.expected, isEven(test.input))
		})
	}
}

func TestMapsGoBoom(t *testing.T) {
	t.Run("will go boom", func(t *testing.T) {
		MapsGoBoom()
	})

	t.Run("recover", func(t *testing.T) {
		defer func() {
			if recover() != nil {
				t.Error("map went boom")
			}
		}()
		MapsGoBoom()
	})

}

func TestSlicesAreSafe(t *testing.T) {
	SlicesAreSafe()
}
