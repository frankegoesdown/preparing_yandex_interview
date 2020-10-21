package main

import "math"

func isReflected(points [][]int) bool {
	m := map[[2]float64]bool{}

	maxX, minX := math.MinInt32, math.MaxInt32

	for _, p := range points {
		m[[2]float64{float64(p[0]), float64(p[1])}] = true
		minX = min(minX, p[0])
		maxX = max(maxX, p[0])
	}

	mid := float64(minX+maxX) / 2.0

	for _, p := range points {
		mirrorX := mid + mid - float64(p[0])
		if !m[[2]float64{mirrorX, float64(p[1])}] {
			return false
		}
	}

	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
