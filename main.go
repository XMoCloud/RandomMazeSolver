package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Select difficulty level
	var choice int
	fmt.Println("Choose difficulty level:")
	fmt.Println("1. Easy")
	fmt.Println("2. Medium")
	fmt.Println("3. Hard")
	fmt.Println("4. Custom")
	fmt.Print("Enter choice (1-4): ")
	fmt.Scanln(&choice)

	n, m, density, minWalls, minPathLen := param(choice)

	var grid [][]int
	var path []Point
	var found bool
	var start, goal Point

	startTime := time.Now()
	// Try generating a valid maze within 3 seconds
	for time.Since(startTime) < 3*time.Second {
		// Create empty grid and fill with random walls
		grid = make([][]int, n)
		wallCount := 0
		for i := 0; i < n; i++ {
			grid[i] = make([]int, m)
			for j := 0; j < m; j++ {
				if rand.Float64() < density {
					grid[i][j] = 1
					wallCount++
				}
			}
		}

		// Random start in the first half of the first row
		start = Point{0, rand.Intn(m / 2)}

		// Random goal in the second half of the last row
		goal = Point{n - 1, rand.Intn(m/2) + m/2}

		grid[start.x][start.y] = 0
		grid[goal.x][goal.y] = 0

		// Try finding a path
		path, found = aStar(grid, start, goal)
		if wallCount < minWalls || !found || len(path) < minPathLen {
			found = false
			continue
		}
		break
	}

	// If no valid maze generated in time, alert the user
	if !found {
		fmt.Println("\nInput error: Could not generate a valid maze within 3 seconds. Try adjusting the input values.")
		return
	}

	// Print maze and solution
	fmt.Printf("\nGenerated %dx%d Maze:\n", n, m)
	printGrid(grid, start, goal, path)

	fmt.Print("\nSolution path: ")
	for i, p := range path {
		if i > 0 {
			fmt.Print(" - ")
		}
		fmt.Print(label(p))
	}
	fmt.Println()
}
