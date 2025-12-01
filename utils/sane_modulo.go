package utils

import "fmt"

// ModuloSane returns the remainder of dividend/modulus, always in range [0, modulus).
// Unlike Go's % operator, this handles negative dividends correctly for wrapping:
//
//	`ModuloSane(-1, 5) = 4  (Go's -1 % 5 = -1)`
//	`ModuloSane(7, 5)  = 2`
//
// Panics if modulus <= 0 (likely a bug at call site).
func ModuloSane(dividend, modulus int) int {
	if modulus <= 0 {
		panic(fmt.Sprintf("modulus %d <=0 is not sane, probably a bug at the call site", modulus))
	}
	return (dividend%modulus + modulus) % modulus
}
