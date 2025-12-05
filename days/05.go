package days

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"log/slog"
	"slices"
	"strings"

	"github.com/mindcrackx/aoc2025/utils"
)

func Five_1(input io.Reader) (string, error) {
	var resultTotal int
	scanner := bufio.NewScanner(input)

	freshRanges := make([]struct {
		First  int
		Second int
	}, 0)
	freshDone := false
	spoiledChecks := make([]int, 0)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			freshDone = true
			continue
		}

		if !freshDone {
			splt := strings.Split(text, "-")
			freshRanges = append(freshRanges, struct {
				First  int
				Second int
			}{
				First:  utils.MustAtoi(splt[0]),
				Second: utils.MustAtoi(splt[1]),
			})
			continue
		}

		spoiledChecks = append(spoiledChecks, utils.MustAtoi(text))
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	for _, num := range spoiledChecks {
		for _, freshRange := range freshRanges {
			if num >= freshRange.First && num <= freshRange.Second {
				resultTotal++
				// slog.Info("fresh!", "num", num, "freshRange", freshRange, "resultTotal", resultTotal)
				break
			}
		}
	}

	return fmt.Sprintf("%d", resultTotal), nil
}

func Five_2(input io.Reader) (string, error) {

	type Range struct {
		First  int
		Second int
	}

	freshRanges := make([]Range, 0)

	var resultTotal int
	scanner := bufio.NewScanner(input)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		text := scanner.Text()

		if text == "" {
			break
		}

		splt := strings.Split(text, "-")
		freshRanges = append(freshRanges, Range{
			First:  utils.MustAtoi(splt[0]),
			Second: utils.MustAtoi(splt[1]),
		})
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	slices.SortFunc(freshRanges, func(a, b Range) int {
		cmp1 := cmp.Compare(a.First, b.First)
		if cmp1 != 0 {
			return cmp1
		}
		return cmp.Compare(a.Second, b.Second)
	})

	contained := false
	newRanges := make([]Range, 0)
	for i := range len(freshRanges) - 1 {
		contained = false
		for y := i + 1; y < len(freshRanges); y++ {
			if freshRanges[i].First >= freshRanges[y].First && freshRanges[i].Second <= freshRanges[y].Second {
				// slog.Info("already fully contained in", "i", freshRanges[i], "y", freshRanges[y])
				contained = true
				break
			}
		}
		if !contained {
			newRanges = append(newRanges, freshRanges[i])
		}
	}
	newRanges = append(newRanges, freshRanges[len(freshRanges)-1])

	slog.Info("", "lenFreshRanges", len(freshRanges), "lenNewRanges", len(newRanges))
	freshRanges = newRanges

	newRanges = make([]Range, 0)
	var start, end int
	for i := 0; i < len(freshRanges)-1; i++ {
		start = freshRanges[i].First
		end = freshRanges[i].Second
		newIndex := i

		for y := i + 1; y < len(freshRanges); y++ {
			if freshRanges[y].First <= end && freshRanges[y].First >= start {
				// use max here, so we only ever extend the merged range, never shrink it!
				end = max(end, freshRanges[y].Second)
				newIndex = y
			}
		}
		newRanges = append(newRanges, Range{
			First:  start,
			Second: end,
		})
		//fmt.Println(i, newIndex)
		i = newIndex
	}

	// last one
	if freshRanges[len(freshRanges)-1].Second > newRanges[len(newRanges)-1].Second {
		newRanges = append(newRanges, freshRanges[len(freshRanges)-1])
	}

	// fmt.Println(freshRanges)
	// fmt.Println(newRanges)

	for _, r := range newRanges {
		fmt.Println(r, resultTotal, r.Second-r.First+1)
		resultTotal += r.Second - r.First + 1
	}

	return fmt.Sprintf("%d", resultTotal), nil
}
