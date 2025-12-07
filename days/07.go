package days

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"
)

func Seven_1(input io.Reader) (string, error) {
	var resultTotal int

	firstLine := true
	positionsNext := make([]int, 0)
	positionsNextNext := make([]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		positionsNext = slices.Clone(positionsNextNext)
		positionsNextNext = make([]int, 0)

		// first find start
		if firstLine {
			firstLine = false
			positionsNextNext = append(positionsNextNext, strings.Index(line, "S"))
		}

		for _, pos := range positionsNext {
			switch line[pos] {
			case '.':
				if !slices.Contains(positionsNextNext, pos) {
					positionsNextNext = append(positionsNextNext, pos)
				}
			case '^':
				// counting
				resultTotal++

				if pos > 0 && !slices.Contains(positionsNextNext, pos-1) {
					positionsNextNext = append(positionsNextNext, pos-1)
				}
				if pos < len(line)-1 && !slices.Contains(positionsNextNext, pos+1) {
					positionsNextNext = append(positionsNextNext, pos+1)
				}
			}
		}
		// fmt.Println(positionsNext, positionsNextNext)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", resultTotal), nil
}

func Seven_2(input io.Reader) (string, error) {
	var resultTotal int

	firstLine := true
	positionsNext := make([]int, 0)
	positionsNextNext := make([]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		positionsNext = slices.Clone(positionsNextNext)
		if firstLine {
			firstLine = false
			positionsNext = make([]int, len(line))
		}
		positionsNextNext = make([]int, len(line))

		for i, val := range line {
			switch val {
			case '.':
				positionsNextNext[i] += positionsNext[i]
			case 'S':
				positionsNextNext[i]++
			case '^':
				positionsNextNext[i-1] += positionsNext[i]
				positionsNextNext[i+1] += positionsNext[i]
			}
		}
		// fmt.Println(line, positionsNext, positionsNextNext)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	for _, result := range positionsNextNext {
		resultTotal += result
	}
	return fmt.Sprintf("%d", resultTotal), nil
}
