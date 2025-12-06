package utils_test

import (
	"slices"
	"testing"

	"github.com/mindcrackx/aoc2025/utils"
)

func TestRotateMatrix2DCounterClockwise(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  [][]int
		expect [][]int
	}{
		{
			name: "same dimensions",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expect: [][]int{
				{3, 6, 9},
				{2, 5, 8},
				{1, 4, 7},
			},
		},
		{
			name: "different dimensions 1",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{0, 11, 12},
			},
			expect: [][]int{
				{3, 6, 9, 12},
				{2, 5, 8, 11},
				{1, 4, 7, 0},
			},
		},
		{
			name: "different dimensions 2",
			input: [][]int{
				{1, 2, 3, 0},
				{4, 5, 6, 11},
				{7, 8, 9, 12},
			},
			expect: [][]int{
				{0, 11, 12},
				{3, 6, 9},
				{2, 5, 8},
				{1, 4, 7},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := utils.RotateMatrix2DCounterClockwise(tt.input)

			if result == nil {
				t.Fatal("expected result to be not nil")
			}

			if len(result) != len(tt.expect) {
				t.Fatalf("expected len expect (%d) to be len result (%d)", len(tt.expect), len(result))
			}

			for i := range tt.expect {
				if slices.Compare(result[i], tt.expect[i]) != 0 {
					t.Fatalf("arr %d expected %v but got %v instead", i, tt.expect, result)
				}
			}

		})
	}
}

func TestRotateMatrix2DCounterClockwise_StringType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  [][]string
		expect [][]string
	}{
		{
			name: "test_1",
			input: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			expect: [][]string{
				{"3", "6", "9"},
				{"2", "5", "8"},
				{"1", "4", "7"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := utils.RotateMatrix2DCounterClockwise(tt.input)

			if result == nil {
				t.Fatal("expected result to be not nil")
			}

			if len(result) != len(tt.expect) {
				t.Fatalf("expected len expect (%d) to be len result (%d)", len(tt.expect), len(result))
			}

			for i := range tt.expect {
				if slices.Compare(result[i], tt.expect[i]) != 0 {
					t.Fatalf("arr %d expected %v but got %v instead", i, tt.expect, result)
				}
			}

		})
	}
}
