package main

import (
	"math"
)

type coord struct {
	x, y, area int
	infinite   bool
}

func findMinMax(coords []*coord) (minX, maxX, minY, maxY int) {
	minX, maxX = coords[0].x, coords[0].x
	minY, maxY = coords[0].y, coords[0].y
	for _, p := range coords {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	return
}

func manhattanDist(x1, x2, y1, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

func computePart1(coords []*coord) *coord {
	minX, maxX, minY, maxY := findMinMax(coords)

	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			minD := math.MaxInt32
			var closestCoord *coord
			for _, p := range coords {
				d := manhattanDist(j, p.x, i, p.y)
				if d < minD {
					minD = d
					closestCoord = p
				} else if d == minD {
					// disqualify because coord is tied for closest, must be unique
					closestCoord = nil
				}
			}
			if closestCoord != nil && (j == minX || j == maxX || i == minY || i == maxY) {
				// disqualify infinite areas (uniquely closest point to the bounding box)
				closestCoord.infinite = true
			}
			if closestCoord != nil {
				closestCoord.area++
			}
		}
	}

	maxAreaPoint := &coord{}
	for _, p := range coords {
		if p.infinite {
			continue
		}

		if p.area > maxAreaPoint.area {
			maxAreaPoint = p
		}
	}

	return maxAreaPoint
}

func computePart2(coords []*coord, maxDist int) int {
	minX, maxX, minY, maxY := findMinMax(coords)

	answer := 0

	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			var totalDist int
			for _, p := range coords {
				d := manhattanDist(j, p.x, i, p.y)
				totalDist += d
			}

			if totalDist < maxDist {
				answer++
			}
		}
	}

	return answer
}
