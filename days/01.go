package days

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mindcrackx/aoc2025/utils"
)

func One_1(input io.Reader) (string, error) {

	rotation := 50
	countZeroRotation := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		rotaDirection := line[0]
		rotaNum := utils.MustAtoi(line[1:])

		switch rotaDirection {
		case 'L':
			rotation = utils.ModuloSane(rotation-rotaNum, 100)
		case 'R':
			rotation = utils.ModuloSane(rotation+rotaNum, 100)
		}

		// slog.Info("", "countZeroRotation", countZeroRotation, "line", line, "rotation", rotation)

		if rotation == 0 {
			countZeroRotation++
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", countZeroRotation), nil
}

func One_2(input io.Reader) (string, error) {

	rotation := 50
	countZeroRotation := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		rotaDirection := line[0]
		rotaNum := utils.MustAtoi(line[1:])

		switch rotaDirection {
		case 'L':
			for i := 0; i < rotaNum; i++ {
				rotation--
				rotation = utils.ModuloSane(rotation, 100)
				if rotation == 0 {
					countZeroRotation++
				}
			}
		case 'R':
			for i := 0; i < rotaNum; i++ {
				rotation++
				rotation = utils.ModuloSane(rotation, 100)
				if rotation == 0 {
					countZeroRotation++
				}
			}
		}

		// slog.Info("", "countZeroRotation", countZeroRotation, "line", line, "rotation", rotation)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", countZeroRotation), nil
}
