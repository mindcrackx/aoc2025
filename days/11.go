package days

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Eleven_1(input io.Reader) (string, error) {
	var resultTotal int

	m := make(map[string][]string)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return "", fmt.Errorf("expected 2 parts for line %q but got %d", line, len(parts))
		}

		machine := parts[0]
		outputs := strings.Fields(strings.TrimSpace(parts[1]))
		m[machine] = outputs
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	// solve
	var solve func(machine string, m map[string][]string) int
	solve = func(machine string, m map[string][]string) int {
		result := 0
		for _, o := range m[machine] {
			if o == "out" {
				return 1
			}
			result += solve(o, m)
		}
		return result
	}

	resultTotal += solve("you", m)

	return fmt.Sprintf("%d", resultTotal), nil
}

func Eleven_2(input io.Reader) (string, error) {
	var resultTotal int

	m := make(map[string][]string)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return "", fmt.Errorf("expected 2 parts for line %q but got %d", line, len(parts))
		}

		machine := parts[0]
		outputs := strings.Fields(strings.TrimSpace(parts[1]))
		m[machine] = outputs
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	// solve
	solveCache := make(map[string]int)
	var solve func(machine string, m map[string][]string, dacFound bool, fftFound bool) int
	solve = func(machine string, m map[string][]string, dacFound bool, fftFound bool) int {

		key := fmt.Sprintf("%s-%t-%t", machine, dacFound, fftFound)
		if val, ok := solveCache[key]; ok {
			return val
		}

		result := 0
		for _, o := range m[machine] {
			if o == "out" {
				if dacFound && fftFound {
					solveCache[key] = 1
					return 1
				}
				solveCache[key] = 0
				return 0
			}
			result += solve(o, m, dacFound || o == "dac", fftFound || o == "fft")
		}

		solveCache[key] = result
		return result
	}

	resultTotal += solve("svr", m, false, false)
	return fmt.Sprintf("%d", resultTotal), nil
}
