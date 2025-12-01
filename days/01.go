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

		// calculate how often we land on "0"
		if rotation == 0 {
			countZeroRotation++
		}

		// slog.Info("", "countZeroRotation", countZeroRotation, "line", line, "rotation", rotation)
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

		// calculate how often we passed "0"
		switch rotaDirection {
		case 'L':
			if !(rotation > rotaNum) {
				if rotation == 0 {
					countZeroRotation += rotaNum / 100
				} else {
					countZeroRotation += 1 + (rotaNum-rotation)/100
				}
			}
			rotation = utils.ModuloSane(rotation-rotaNum, 100)
		case 'R':
			countZeroRotation += (rotation + rotaNum) / 100
			rotation = utils.ModuloSane(rotation+rotaNum, 100)
		}

		// slog.Info("", "countZeroRotation", countZeroRotation, "line", line, "rotation", rotation)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", countZeroRotation), nil
}
