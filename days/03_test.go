package days_test

import (
	"strings"
	"testing"

	"github.com/mindcrackx/aoc2025/days"
)

func TestThree_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name: "example 1",
			input: `987654321111111
811111111111119
234234234234278
818181911112111`,
			expect: "357",
		},
		{
			name:   "example 2",
			input:  "987654321111111",
			expect: "98",
		},
		{
			name:   "example 3",
			input:  "811111111111119",
			expect: "89",
		},
		{
			name:   "example 4",
			input:  "234234234234278",
			expect: "78",
		},
		{
			name:   "example 5",
			input:  "818181911112111",
			expect: "92",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := days.Three_1(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("got unexpected error: %s", err.Error())
			}

			if result != tt.expect {
				t.Fatalf("expected %q but got %q instead", tt.expect, result)
			}
		})
	}
}

func TestThree_2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name: "example 1",
			input: `987654321111111
811111111111119
234234234234278
818181911112111`,
			expect: "3121910778619",
		},
		{
			name:   "example 2",
			input:  "987654321111111",
			expect: "987654321111",
		},
		{
			name:   "example 3",
			input:  "811111111111119",
			expect: "811111111119",
		},
		{
			name:   "example 4",
			input:  "234234234234278",
			expect: "434234234278",
		},
		{
			name:   "example 5",
			input:  "818181911112111",
			expect: "888911112111",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := days.Three_2(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("got unexpected error: %s", err.Error())
			}

			if result != tt.expect {
				t.Fatalf("expected %q but got %q instead", tt.expect, result)
			}
		})
	}
}
