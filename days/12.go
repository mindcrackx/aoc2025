package days

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

func Twelve_1(input io.Reader) (string, error) {
	var resultTotal int

	shapes := make([][][]int, 0, 6)
	toSolve := make([]string, 0)

	inShapes := true
	inAreas := false

	currShapeLines := make([]string, 0, 3)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		if inShapes && strings.Contains(line, "x") {
			inShapes = false
			inAreas = true
		}

		if inShapes {
			if strings.Contains(line, ":") {
				currShapeLines = make([]string, 0)
				continue
			}

			// shape finished!
			if strings.TrimSpace(line) == "" {
				newShape := make([][]int, 3)
				newShape[0] = make([]int, 3)
				newShape[1] = make([]int, 3)
				newShape[2] = make([]int, 3)

				fmt.Printf("%#v\n", currShapeLines)
				for i, l := range currShapeLines {
					for y, r := range l {
						fmt.Println(i, l, y, string(r), r)
						switch r {
						case '.':
							newShape[i][y] = 0
						case '#':
							newShape[i][y] = 1
						}
					}
				}

				shapes = append(shapes, newShape)

				continue
			}

			currShapeLines = append(currShapeLines, line)
		}

		if inAreas {
			// append to area problems
			toSolve = append(toSolve, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	fmt.Printf("toSolve: %#v\n", toSolve)
	fmt.Println()
	fmt.Println("shapes:")
	for _, s := range shapes {
		fmt.Printf("%#v\n\n", s)
	}

	return fmt.Sprintf("%d", resultTotal), nil
}

func Twelve_2(input io.Reader) (string, error) {
	return "", errors.New("not implemented")
}
