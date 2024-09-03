package main

import (
	"testing"
)

// TestInitializeGridSize checks if the grid is correctly initialized with the specified dimensions.
func TestInitializeGridSize(t *testing.T) {
	grid := initializeGrid(5, 5)
	if len(grid) != 5 {
		t.Errorf("Expected grid width of 5, but got %d", len(grid))
	}
	if len(grid[0]) != 5 {
		t.Errorf("Expected grid height of 5, but got %d", len(grid[0]))
	}
}

// TestMarkObstacles checks if obstacles are correctly marked on the grid.
func TestMarkObstacles(t *testing.T) {
	grid := initializeGrid(5, 5)
	markObstacles(grid, 1, 3, 1, 3)
	if !grid[1][1] {
		t.Errorf("Expected grid[1][1] to be true, but got false")
	}
	if !grid[3][3] {
		t.Errorf("Expected grid[3][3] to be true, but got false")
	}
}

// TestIsValidMoveInsideBounds checks valid moves within the grid boundaries.
func TestIsValidMoveInsideBounds(t *testing.T) {
	grid := initializeGrid(5, 5)
	if !isValidMove(1, 1, 5, 5, grid) {
		t.Error("Expected move (1, 1) to be valid, but it was not.")
	}
}

// TestIsValidMoveOutsideBounds checks invalid moves outside the grid boundaries.
func TestIsValidMoveOutsideBounds(t *testing.T) {
	grid := initializeGrid(5, 5)
	if isValidMove(-1, 0, 5, 5, grid) {
		t.Error("Expected move (-1, 0) to be invalid due to being out of bounds, but it was considered valid.")
	}
}

// TestIsValidMoveWithObstacle checks invalid moves into obstacle positions.
func TestIsValidMoveWithObstacle(t *testing.T) {
	grid := initializeGrid(5, 5)
	markObstacles(grid, 2, 2, 2, 2)
	if isValidMove(2, 2, 5, 5, grid) {
		t.Error("Expected move (2, 2) to be invalid due to obstacle, but it was considered valid.")
	}
}

func TestFindMinimumHopsNoSolution(t *testing.T) {
	grid := initializeGrid(3, 3)
	markObstacles(grid, 1, 1, 0, 2)
	markObstacles(grid, 0, 2, 1, 1)
	start := Point{0, 0}
	finish := Point{2, 2}

	// Expect -1 hops because the path is blocked by obstacles
	hops := findMinimumHops(start, finish, 5, 5, grid)
	if hops != -1 {
		t.Errorf("Expected -1 hops (no solution), but got %d hops", hops)
	}
}

func TestFindMinimumHopsValidPath(t *testing.T) {
	grid := initializeGrid(5, 5)
	start := Point{0, 0}
	finish := Point{4, 4}

	// Test with no obstacles for a valid path
	hops := findMinimumHops(start, finish, 5, 5, grid)
	if hops == -1 {
		t.Error("Expected a valid path, but got no solution.")
	}
}
