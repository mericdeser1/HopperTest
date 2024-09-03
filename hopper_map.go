package main

func initializeGrid(width, height int) [][]bool {
	grid := make([][]bool, width)
	for i := range grid {
		grid[i] = make([]bool, height)
	}
	return grid
}

func markObstacles(grid [][]bool, x1, x2, y1, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			grid[x][y] = true
		}
	}
}
