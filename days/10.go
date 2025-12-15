package days

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"slices"
	"strings"
	"sync"

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
	var resultTotal int

	sem := make(chan struct{}, 14)
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(input)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		idx1 := strings.Index(line, "]")
		idx2 := strings.Index(line, "{")

		// lightsStr := strings.TrimSpace(line[0 : idx1+1])
		buttonsStr := strings.TrimSpace(line[idx1+1 : idx2-1])
		joltageStr := strings.TrimSpace(line[idx2+1 : len(line)-1])

		buttons := make([][]int, 0)
		for i, bStr := range strings.Fields(buttonsStr) {
			buttonsStrs := strings.Split(bStr[1:len(bStr)-1], ",")
			buttons = append(buttons, make([]int, len(buttonsStrs)))
			for ii, b := range buttonsStrs {
				buttons[i][ii] = utils.MustAtoi(b)
			}
		}

		joltages := make([]int, 0)
		for _, j := range strings.Split(joltageStr, ",") {
			joltages = append(joltages, utils.MustAtoi(j))
		}

		wg.Add(1)
		go func(ln int) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			result := solveTenPart2(joltages, buttons)
			resultTotal += result
			fmt.Printf("ln %d |->  %d\n", ln, result)
		}(lineNum)

	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	wg.Wait()

	return fmt.Sprintf("%d", resultTotal), nil
}

func solveTenPart2(target []int, buttons [][]int) int {
	best := math.MaxInt

	lastBtnIdxForCounter := make([]int, len(target))
	for i := range lastBtnIdxForCounter {
		lastBtnIdxForCounter[i] = -1
	}
	for btnIdx, btn := range buttons {
		for _, counterIdx := range btn {
			lastBtnIdxForCounter[counterIdx] = btnIdx
		}
	}

	var solve func(btnIdx int, current []int, pressesTotal int)
	solve = func(btnIdx int, current []int, pressesTotal int) {
		if pressesTotal >= best {
			return // we already found a better solution
		}

		// pruning 1
		for i := range current {
			if lastBtnIdxForCounter[i] < btnIdx {
				// no remaining btn affects counter i, impossible to reach target
				if current[i] != target[i] {
					return
				}
			}
		}

		// pruning 2
		// estimate min presses still needed
		minRemaining := 0
		for i := range current {
			diff := target[i] - current[i]
			if diff > minRemaining {
				minRemaining = diff
			}
		}
		// each btn press can only increase by 1!
		if pressesTotal+minRemaining >= best {
			return
		}

		if btnIdx == len(buttons) {
			if slices.Equal(current, target) {
				if pressesTotal < best {
					best = pressesTotal
				}
			}
			return
		}

		var exceeded bool
		for presses := 0; ; presses++ {
			exceeded = false
			for i := range current {
				if current[i] > target[i] {
					exceeded = true
					break
				}
			}
			if exceeded {
				break
			}

			next := make([]int, len(current))
			copy(next, current)

			// recurse
			solve(btnIdx+1, next, pressesTotal+presses)

			// press btn again for next iteration
			for _, idx := range buttons[btnIdx] {
				current[idx]++
			}
		}
	}

	solve(0, make([]int, len(target)), 0)
	return best
}
