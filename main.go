package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	if !scanner.Scan() {
		fmt.Println("Error reading number of test cases.")
		return
	}
	testCases := 0
	fmt.Sscanf(scanner.Text(), "%d", &testCases)

	for t := 0; t < testCases; t++ {
		if !scanner.Scan() {
			fmt.Println("Error reading grid dimensions.")
			return
		}
		var width, height int
		fmt.Sscanf(scanner.Text(), "%d %d", &width, &height)

		if !scanner.Scan() {
			fmt.Println("Error reading start and finish points.")
			return
		}
		var startX, startY, finishX, finishY int
		fmt.Sscanf(scanner.Text(), "%d %d %d %d", &startX, &startY, &finishX, &finishY)

		grid := initializeGrid(width, height)

		if !scanner.Scan() {
			fmt.Println("Error reading number of obstacles.")
			return
		}
		obstacles := 0
		fmt.Sscanf(scanner.Text(), "%d", &obstacles)

		// Mark obstacles on the grid
		for i := 0; i < obstacles; i++ {
			if !scanner.Scan() {
				fmt.Println("Error reading obstacle.")
				return
			}
			var x1, x2, y1, y2 int
			fmt.Sscanf(scanner.Text(), "%d %d %d %d", &x1, &x2, &y1, &y2)
			markObstacles(grid, x1, x2, y1, y2)
		}

		start := Point{startX, startY}
		finish := Point{finishX, finishY}

		// Find the minimum hops from start to finish
		result := findMinimumHops(start, finish, width, height, grid)

		if result == -1 {
			fmt.Printf("No solution.")
			fmt.Fprintln(outputFile, "No solution.")
		} else {
			fmt.Printf("Optimal solution takes %d hops.\n", result)
			fmt.Fprintf(outputFile, "Optimal solution takes %d hops.\n", result)
		}
	}
}
