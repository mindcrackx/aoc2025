package days

import (
	"bufio"
	"fmt"
	"io"
)

func Four_1(input io.Reader) (string, error) {

	grid := make([][]int, 0, 0)

	scanner := bufio.NewScanner(input)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, make([]int, len(line)))

		for i, r := range line {
			if r == '@' {
				grid[lineNum][i] = 1
			}
		}

		lineNum++
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	gridNew := make([][]int, 0, len(grid))
	for x := range grid {
		gridNew = append(gridNew, make([]int, len(grid[x])))
		copy(gridNew[x], grid[x])
	}

	var result int
	for x := range grid {
		maxX := len(grid) - 1
		for y := range grid[x] {
			if grid[x][y] == 0 {
				// slog.Info("skipped", "x", x, "y", y, "obj", grid[x][y])
				continue
			}

			maxY := len(grid[x]) - 1

			found := 0

			if x > 0 && y > 0 && grid[x-1][y-1] == 1 {
				found++
			}
			if x > 0 && grid[x-1][y] == 1 {
				found++
			}
			if x > 0 && y < maxY && grid[x-1][y+1] == 1 {
				found++
			}

			if y > 0 && grid[x][y-1] == 1 {
				found++
			}
			if y < maxY && grid[x][y+1] == 1 {
				found++
			}

			if x < maxX && y > 0 && grid[x+1][y-1] == 1 {
				found++
			}
			if x < maxX && grid[x+1][y] == 1 {
				found++
			}
			if x < maxX && y < maxY && grid[x+1][y+1] == 1 {
				found++
			}

			// slog.Info("", "x", x, "y", y, "obj", grid[x][y], "found", found)
			if found < 4 {
				result++
				gridNew[x][y] = 2
			}
		}
	}

	// PrintGrid(gridNew)

	return fmt.Sprintf("%d", result), nil
}

func PrintGrid(grid [][]int) {
	for x := range grid {
		for y := range grid[x] {
			fmt.Print(grid[x][y])
		}
		fmt.Println()
	}
}

func Four_2(input io.Reader) (string, error) {

	grid := make([][]int, 0, 0)

	scanner := bufio.NewScanner(input)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, make([]int, len(line)))

		for i, r := range line {
			if r == '@' {
				grid[lineNum][i] = 1
			}
		}

		lineNum++
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	solve := func(grid [][]int) (int, [][]int) {

		gridNew := make([][]int, 0, len(grid))
		for x := range grid {
			gridNew = append(gridNew, make([]int, len(grid[x])))
			copy(gridNew[x], grid[x])
		}

		var result int
		for x := range grid {
			maxX := len(grid) - 1
			for y := range grid[x] {
				if grid[x][y] == 0 {
					// slog.Info("skipped", "x", x, "y", y, "obj", grid[x][y])
					continue
				}

				maxY := len(grid[x]) - 1

				found := 0

				if x > 0 && y > 0 && grid[x-1][y-1] == 1 {
					found++
				}
				if x > 0 && grid[x-1][y] == 1 {
					found++
				}
				if x > 0 && y < maxY && grid[x-1][y+1] == 1 {
					found++
				}

				if y > 0 && grid[x][y-1] == 1 {
					found++
				}
				if y < maxY && grid[x][y+1] == 1 {
					found++
				}

				if x < maxX && y > 0 && grid[x+1][y-1] == 1 {
					found++
				}
				if x < maxX && grid[x+1][y] == 1 {
					found++
				}
				if x < maxX && y < maxY && grid[x+1][y+1] == 1 {
					found++
				}

				// slog.Info("", "x", x, "y", y, "obj", grid[x][y], "found", found)
				if found < 4 {
					result++
					gridNew[x][y] = 0
				}
			}
		}
		return result, gridNew
	}

	var resultTotal int

	result, gridNew := solve(grid)

	resultTotal += result
	// fmt.Println(resultTotal, result)

	for result != 0 {
		// do it again

		// PrintGrid(gridNew)

		result, gridNew = solve(gridNew)

		resultTotal += result
		// fmt.Println(resultTotal, result)
	}

	// PrintGrid(gridNew)

	return fmt.Sprintf("%d", resultTotal), nil
}
