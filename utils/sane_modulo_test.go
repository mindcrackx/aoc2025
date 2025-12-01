package utils_test

import (
	"testing"

	"github.com/mindcrackx/aoc2025/utils"
)

func TestModuloSane(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		num      int
		modulo   int
		expected int
	}{
		"positive": {
			5,
			4,
			1,
		},
		"zero": {
			0,
			4,
			0,
		},
		"negative": {
			-4,
			6,
			2,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := utils.ModuloSane(test.num, test.modulo)
			if res != test.expected {
				t.Errorf("expected %d but got %d for num=%d and mod=%d", test.expected, res, test.num, test.modulo)
			}
		})
	}
}

func TestModuloSane_ShouldPanicForBadModuloInput(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		a int
		b int
	}{
		"zero modulo":     {5, 0},
		"negative modulo": {5, -3},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			defer func() {
				if recover() == nil {
					t.Errorf("expected panic for a=%d, b=%d", test.a, test.b)
				}
			}()

			utils.ModuloSane(test.a, test.b)
		})
	}
}

func FuzzModuloSane(f *testing.F) {
	f.Add(5, 4)
	f.Add(0, 4)
	f.Add(-4, 6)
	f.Fuzz(func(t *testing.T, a, b int) {
		// For invalid modulus, verify it panics
		if b <= 0 {
			defer func() {
				if recover() == nil {
					t.Errorf("expected panic for a=%d, b=%d", a, b)
				}
			}()
			utils.ModuloSane(a, b)
			return
		}

		// For valid modulus, verify properties
		result := utils.ModuloSane(a, b)
		if result < 0 || result >= b {
			t.Errorf("result %d out of range [0, %d) for a=%d, b=%d", result, b, a, b)
		}
		if (a-result)%b != 0 {
			t.Errorf("a=%d and result=%d are not congruent mod %d", a, result, b)
		}
	})
}
