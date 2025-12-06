package days_test

import (
	"strings"
	"testing"

	"github.com/mindcrackx/aoc2025/days"
)

func TestSix_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name: "example 1",
			input: `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `,
			expect: "4277556",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := days.Six_1(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("got unexpected error: %s", err.Error())
			}

			if result != tt.expect {
				t.Fatalf("expected %q but got %q instead", tt.expect, result)
			}
		})
	}
}

func TestSix_2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name: "example 1",
			input: `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `,
			expect: "3263827",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := days.Six_2(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("got unexpected error: %s", err.Error())
			}

			if result != tt.expect {
				t.Fatalf("expected %q but got %q instead", tt.expect, result)
			}
		})
	}
}
