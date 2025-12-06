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

/*
[
	[1,2,3],
	[1,2,3],
	[1,2,3],
	[1,2,3],
	[+,*,+],
]

123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +


 46 15  823 321
 32 783  46 54
413 512  89 6
  +   *   +   *

*/

func RotateArray2DLeft[T any](aa [][]T) [][]T {
	rowLength := len(aa)
	colLength := len(aa[0])

	bb := make([][]T, 0, colLength)

	for i := 0; i < colLength; i++ {
		tmp := make([]T, 0, rowLength)
		for j := 0; j < rowLength; j++ {
			tmp = append(tmp, aa[j][colLength-1-i])
		}
		bb = append(bb, tmp)
	}

	return bb
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

	ss := RotateArray2DLeft(lines)

	var result int
	tmp := make([]string, 0)
	for _, line := range ss {
		s := string(line)
		if strings.TrimSpace(s) == "" {
			if len(s) == 0 {
				continue
			}
			// calc
			tmpStr := tmp[len(tmp)-1]
			op := tmpStr[len(tmpStr)-1]
			tmp[len(tmp)-1] = tmpStr[0 : len(tmpStr)-1]

			switch string(op) {
			case "+":
				result = 0
			case "*":
				result = 1
			}

			for i := range tmp {
				switch string(op) {
				case "+":
					result += utils.MustAtoi(strings.TrimSpace(tmp[i]))
				case "*":
					result *= utils.MustAtoi(strings.TrimSpace(tmp[i]))
				}
			}

			resultTotal += result
			// slog.Info("", "result", result, "tmp", tmp)
			tmp = make([]string, 0)
			continue
		}
		tmp = append(tmp, s)
	}
	// last one
	tmpStr := tmp[len(tmp)-1]
	op := tmpStr[len(tmpStr)-1]
	tmp[len(tmp)-1] = tmpStr[0 : len(tmpStr)-1]

	switch string(op) {
	case "+":
		result = 0
	case "*":
		result = 1
	}

	for i := range tmp {
		switch string(op) {
		case "+":
			result += utils.MustAtoi(strings.TrimSpace(tmp[i]))
		case "*":
			result *= utils.MustAtoi(strings.TrimSpace(tmp[i]))
		}
	}

	resultTotal += result

	return fmt.Sprintf("%d", resultTotal), nil
}
