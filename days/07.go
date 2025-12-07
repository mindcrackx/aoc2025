package days

import (
	"bufio"
	"fmt"
	"io"
)

func Seven_1(input io.Reader) (string, error) {
	var resultTotal int

	firstLine := true
	current := make([]int, 0)
	next := make([]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		if firstLine {
			firstLine = false
			current = make([]int, len(line))
			next = make([]int, len(line))
		}
		current, next = next, current
		clear(next)

		for i, val := range line {
			switch val {
			case '.':
				next[i] += current[i]
			case 'S':
				next[i]++
			case '^':
				if current[i] != 0 {
					resultTotal++
				}
				next[i-1] += current[i]
				next[i+1] += current[i]
			}
		}
		// fmt.Println(line, current, next)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", resultTotal), nil
}

func Seven_2(input io.Reader) (string, error) {
	var resultTotal int

	firstLine := true
	current := make([]int, 0)
	next := make([]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		if firstLine {
			firstLine = false
			current = make([]int, len(line))
			next = make([]int, len(line))
		}
		current, next = next, current
		clear(next)

		for i, val := range line {
			switch val {
			case '.':
				next[i] += current[i]
			case 'S':
				next[i]++
			case '^':
				next[i-1] += current[i]
				next[i+1] += current[i]
			}
		}
		// fmt.Println(line, current, next)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	for _, result := range next {
		resultTotal += result
	}
	return fmt.Sprintf("%d", resultTotal), nil
}
