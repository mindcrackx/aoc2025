package days

import (
	"fmt"
	"io"
	"strings"

	"github.com/mindcrackx/aoc2025/utils"
)

func Two_1(input io.Reader) (string, error) {
	data, err := io.ReadAll(input)
	if err != nil {
		return "", err
	}

	result := 0

	parts := strings.Split(strings.TrimSpace(string(data)), ",")
	for _, part := range parts {
		ids := strings.Split(part, "-")
		for i := utils.MustAtoi(ids[0]); i <= utils.MustAtoi(ids[1]); i++ {
			istr := fmt.Sprintf("%d", i)
			if len(istr)%2 != 0 {
				continue
			}

			if istr[0:len(istr)/2] == istr[len(istr)/2:] {
				// slog.Info("found", "istr", istr)
				result += i
			}
		}
	}

	return fmt.Sprintf("%d", result), nil
}

func Two_2(input io.Reader) (string, error) {
	data, err := io.ReadAll(input)
	if err != nil {
		return "", err
	}

	result := 0

	parts := strings.Split(strings.TrimSpace(string(data)), ",")
	for _, part := range parts {
		ids := strings.Split(part, "-")
		for i := utils.MustAtoi(ids[0]); i <= utils.MustAtoi(ids[1]); i++ {
			istr := fmt.Sprintf("%d", i)

			for x := range len(istr) {
				if x == 0 {
					continue
				}
				repeat := strings.Repeat(istr[0:x], len(istr)/(x))
				if istr == repeat {
					// slog.Info("found", "istr", istr)
					result += i
					break
				}
			}

		}
	}

	return fmt.Sprintf("%d", result), nil
}
