package days

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/mindcrackx/aoc2025/utils"
)

func Ten_1(input io.Reader) (string, error) {
	var resultTotal int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		idx1 := strings.Index(line, "]")
		idx2 := strings.Index(line, "{")

		lightsStr := strings.TrimSpace(line[0 : idx1+1])
		buttonsStr := strings.TrimSpace(line[idx1+1 : idx2-1])

		lights := make([]int, len(lightsStr)-2)
		for i, r := range lightsStr[1 : len(lightsStr)-1] {
			if r == '#' {
				lights[i] = 1
			}
		}

		buttons := make([][]int, 0)
		for i, bStr := range strings.Fields(buttonsStr) {
			buttonsStrs := strings.Split(bStr[1:len(bStr)-1], ",")
			buttons = append(buttons, make([]int, len(buttonsStrs)))
			for ii, b := range buttonsStrs {
				buttons[i][ii] = utils.MustAtoi(b)
			}
		}

		solution := tenSolveBFS(lights, buttons)
		resultTotal += len(solution)

	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", resultTotal), nil
}

func tenSolveBFS(target []int, buttons [][]int) []int {
	targetStr := fmt.Sprint(target)
	type state struct {
		current []int
		pressed []int
	}

	start := make([]int, len(target))
	visited := map[string]bool{fmt.Sprint(start): true}
	queue := []state{{
		current: start,
		pressed: nil,
	}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if fmt.Sprint(cur.current) == targetStr {
			return cur.pressed
		}

		for btnIdx, button := range buttons {
			next := make([]int, len(cur.current))
			copy(next, cur.current)

			for _, i := range button {
				next[i] ^= 1 // toggle
			}

			key := fmt.Sprint(next)
			if !visited[key] {
				visited[key] = true
				newPressed := append([]int{}, cur.pressed...)
				newPressed = append(newPressed, btnIdx)
				queue = append(queue, state{next, newPressed})
			}
		}
	}

	return nil
}

func Ten_2(input io.Reader) (string, error) {
	return "", errors.New("not implemented")
}
