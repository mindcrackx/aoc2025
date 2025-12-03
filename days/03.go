package days

import (
	"bufio"
	"fmt"
	"io"
)

func Three_1(input io.Reader) (string, error) {
	var resultTotal int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		nums := make([]int, 0, len(line))
		for _, r := range line {
			// optimize to not use utils.MustAtoI(string(r)), saves like 40% execution time
			// So if r = '5':
			// - int(r) → 53 (the ASCII code point)
			// - int(r - '0') → 53 - 48 → 5 (the actual digit)
			nums = append(nums, int(r-'0'))
		}

		first := -1
		second := -1
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] <= first {
				continue
			}
			first = nums[i]
			tmpSecond := -1

			for y := i + 1; y < len(nums); y++ {
				if nums[y] <= tmpSecond {
					continue
				}
				tmpSecond = nums[y]
			}
			second = tmpSecond
		}
		resultTotal += first*10 + second
		// slog.Info("", "resultTotal", resultTotal, "first", first, "second", second)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", resultTotal), nil
}

func Three_2(input io.Reader) (string, error) {
	var resultTotal int

	nextBest := func(nums []int, currentI, remaining int) (int, int) {
		currBest := -1
		currI := -1
		for i := currentI + 1; i < len(nums)-remaining+1; i++ {
			if nums[i] <= currBest {
				continue
			}
			currBest = nums[i]
			currI = i
		}
		return currI, currBest
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		nums := make([]int, 0, len(line))
		for _, r := range line {
			// optimize to not use utils.MustAtoI(string(r)), saves like 40% execution time
			// So if r = '5':
			// - int(r) → 53 (the ASCII code point)
			// - int(r - '0') → 53 - 48 → 5 (the actual digit)
			nums = append(nums, int(r-'0'))
		}

		bestNums := make([]int, 0, 12)
		var currBest = -1
		var currI = -1

		for i := range 12 {
			currI, currBest = nextBest(nums, currI, 12-i)
			bestNums = append(bestNums, currBest)
			// slog.Info("", "i", i, "currI", currI, "currBest", currBest, "bestNums", bestNums)
		}

		var total int
		for _, n := range bestNums {
			total = total*10 + n
		}
		resultTotal += total
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", resultTotal), nil
}
