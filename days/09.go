package days

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/mindcrackx/aoc2025/utils"
)

func Nine_1(input io.Reader) (string, error) {
	/*
		makeGrid := func(points []Point) [][]int {
			var maxX, maxY int
			for _, p := range points {
				if p.X > maxX {
					maxX = p.X
				}
				if p.Y > maxY {
					maxY = p.Y
				}
			}
			maxX++
			maxY++

			grid := make([][]int, maxX)
			for i := 0; i < maxX; i++ {
				grid[i] = make([]int, maxY)
			}

			for i, p := range points {
				grid[p.X][p.Y] = i + 1
			}

			return grid
		}

		printGrid := func(grid [][]int) {
			for col := 0; col < len(grid[0]); col++ {
				for row := 0; row < len(grid); row++ {
					switch grid[row][col] {
					case 0:
						fmt.Print(".")
					case -1:
						fmt.Print("#")
					// case 1:
					// fmt.Print("#")
					// case 2:
					// fmt.Print("0")
					default:
						fmt.Print(grid[row][col])
					}
				}
				fmt.Println()
			}
		}
	*/

	var resultTotal int

	type Point struct {
		X int
		Y int
	}

	points := make([]Point, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		p := Point{
			X: utils.MustAtoi(parts[0]),
			Y: utils.MustAtoi(parts[1]),
		}
		points = append(points, p)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	for i, p := range points[0 : len(points)-1] {
		for _, o := range points[i+1:] {
			left := min(p.X, o.X)
			right := max(p.X, o.X)
			top := min(p.Y, o.Y)
			bottom := max(p.Y, o.Y)
			area := (right + 1 - left) * (bottom + 1 - top)

			if area > resultTotal {
				resultTotal = area
			}
		}
	}

	return fmt.Sprintf("%d", resultTotal), nil
}

func Nine_2(input io.Reader) (string, error) {
	return "", errors.New("not implemented")
}
