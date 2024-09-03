package main

type Point struct {
	x, y int
}

type State struct {
	position Point
	velocity Point
	hops     int
}

func findMinimumHops(start, finish Point, width, height int, grid [][]bool) int {
	directions := []int{-1, 0, 1}

	visited := make(map[Point]map[Point]bool)
	initialState := State{start, Point{0, 0}, 0}
	queue := []State{initialState}

	if visited[start] == nil {
		visited[start] = make(map[Point]bool)
	}
	visited[start][Point{0, 0}] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.position == finish {
			return current.hops
		}

		for _, dx := range directions {
			for _, dy := range directions {
				newVelocity := Point{current.velocity.x + dx, current.velocity.y + dy}
				newPosition := Point{current.position.x + newVelocity.x, current.position.y + newVelocity.y}

				if isValidMove(newPosition.x, newPosition.y, width, height, grid) {
					if visited[newPosition] == nil {
						visited[newPosition] = make(map[Point]bool)
					}
					if !visited[newPosition][newVelocity] {
						visited[newPosition][newVelocity] = true
						newState := State{newPosition, newVelocity, current.hops + 1}
						queue = append(queue, newState)
					}
				}
			}
		}
	}

	return -1 //If there is no solution
}

func isValidMove(x, y, width, height int, grid [][]bool) bool {
	return x >= 0 && x < width && y >= 0 && y < height && !grid[x][y]
}
