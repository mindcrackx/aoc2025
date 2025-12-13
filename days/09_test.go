package days_test

import (
	"strings"
	"testing"

	"github.com/mindcrackx/aoc2025/days"
)

func TestNine_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name: "example 1",
			input: `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`,
			expect: "50",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := days.Nine_1(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("got unexpected error: %s", err.Error())
			}

			if result != tt.expect {
				t.Fatalf("expected %q but got %q instead", tt.expect, result)
			}
		})
	}
}

func TestNine_2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		expect string
	}{}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := days.Nine_2(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("got unexpected error: %s", err.Error())
			}

			if result != tt.expect {
				t.Fatalf("expected %q but got %q instead", tt.expect, result)
			}
		})
	}
}
