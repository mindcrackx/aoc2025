package days

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/mindcrackx/aoc2025/utils"
)

func Six_1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var resultTotal int

	lines := make([][]string, 0)

	firstLine := true
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Fields(line)
		if firstLine {
			for range len(fields) {
				lines = append(lines, make([]string, 0))
			}
			firstLine = false
		}

		for i := range fields {
			lines[i] = append(lines[i], fields[i])
		}

	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	var result int
	for i := range lines {
		op := lines[i][len(lines[i])-1]
		switch op {
		case "+":
			result = 0
			for y := 0; y < len(lines[i])-1; y++ {
				result += utils.MustAtoi(lines[i][y])
			}
		case "*":
			result = 1
			for y := 0; y < len(lines[i])-1; y++ {
				result *= utils.MustAtoi(lines[i][y])
			}
		}

		resultTotal += result
	}

	return fmt.Sprintf("%d", resultTotal), nil
}

func Six_2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var resultTotal int

	lines := make([][]byte, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		lines = append(lines, []byte(line))

	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	matrix := utils.RotateMatrix2DCounterClockwise(lines)

	calculate := func(arr []string) int {
		var result int
		tmpStr := arr[len(arr)-1]
		op := tmpStr[len(tmpStr)-1]
		lastNum := tmpStr[0 : len(tmpStr)-1] // don't mutate arr

		switch op {
		case '+':
			result = 0
			for i := 0; i < len(arr)-1; i++ {
				result += utils.MustAtoi(strings.TrimSpace(arr[i]))
			}
			result += utils.MustAtoi(strings.TrimSpace(lastNum))
		case '*':
			result = 1
			for i := 0; i < len(arr)-1; i++ {
				result *= utils.MustAtoi(strings.TrimSpace(arr[i]))
			}
			result *= utils.MustAtoi(strings.TrimSpace(lastNum))
		}
		return result
	}

	tmp := make([]string, 0)
	for _, line := range matrix {
		s := string(line)
		if len(s) > 0 && strings.TrimSpace(s) == "" {
			resultTotal += calculate(tmp)
			tmp = tmp[:0] // reuse existing slice
			continue
		}
		tmp = append(tmp, s)
	}
	// last one
	resultTotal += calculate(tmp)

	return fmt.Sprintf("%d", resultTotal), nil
}
