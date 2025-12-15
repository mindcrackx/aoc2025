package days

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"slices"
	"strings"

	"github.com/mindcrackx/aoc2025/utils"
)

type EightPoint struct {
	X int
	Y int
	Z int
}

func (p EightPoint) Distance(pp EightPoint) float64 {
	return math.Sqrt(
		math.Pow(float64(p.X)-float64(pp.X), 2) + math.Pow(float64(p.Y)-float64(pp.Y), 2) + math.Pow(float64(p.Z)-float64(pp.Z), 2))
}

func Eight_1(input io.Reader) (string, error) {
	var resultTotal int

	points := make([]EightPoint, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		fields := strings.FieldsFunc(scanner.Text(), func(r rune) bool { return r == ',' })
		p := EightPoint{}
		p.X, p.Y, p.Z = utils.MustAtoi(fields[0]), utils.MustAtoi(fields[1]), utils.MustAtoi(fields[2])
		// fmt.Println(p)
		points = append(points, p)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	parent := make([]int, len(points))
	for i := range parent {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		pa, pb := find(a), find(b)
		if pa != pb {
			parent[pa] = pb
		}
	}

	connections := make([][]int, 0, len(points))

	for len(connections) < len(points) {

		lowestDistanceIdx1 := -1
		lowestDistanceIdx2 := -1
		lowestDistance := math.MaxFloat64
		lowestDistanceNew := math.MaxFloat64
		for i := 0; i < len(points)-1; i++ {
			for y := i + 1; y < len(points); y++ {
				lowestDistanceNew = math.Abs(points[i].Distance(points[y]))
				if lowestDistanceNew < lowestDistance {
					found := false
					for _, c := range connections {
						if c[0] == min(i, y) && c[1] == max(i, y) {
							found = true
							break
						}
					}
					if !found {
						lowestDistance = lowestDistanceNew
						lowestDistanceIdx1 = i
						lowestDistanceIdx2 = y

						/*
							slog.Info("new lowest Distance",
								"distance", lowestDistance, "p1_idx", lowestDistanceIdx1, "p1", points[lowestDistanceIdx1],
								"p2_idx", lowestDistanceIdx2, "p2", points[lowestDistanceIdx2],
							)
						*/
					}
				}
			}
		}
		/*
			slog.Info("===lowest Distance",
				"distance", lowestDistance, "p1_idx", lowestDistanceIdx1, "p1", points[lowestDistanceIdx1],
				"p2_idx", lowestDistanceIdx2, "p2", points[lowestDistanceIdx2],
			)
		*/
		connections = append(connections, []int{min(lowestDistanceIdx1, lowestDistanceIdx2), max(lowestDistanceIdx1, lowestDistanceIdx2)})
		union(lowestDistanceIdx1, lowestDistanceIdx2)
	}

	// fmt.Println(connections)

	circuitSizes := make(map[int]int)
	for i := range points {
		root := find(i)
		circuitSizes[root]++
	}

	sizes := make([]int, 0, len(circuitSizes))
	for _, size := range circuitSizes {
		sizes = append(sizes, size)
	}
	slices.Sort(sizes)
	slices.Reverse(sizes)

	resultTotal = 1 * sizes[0] * sizes[1] * sizes[2]

	return fmt.Sprintf("%d", resultTotal), nil
}

func Eight_2(input io.Reader) (string, error) {
	return "", errors.New("not implemented")
}
