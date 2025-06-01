package main

import "fmt"

// label converts a Point to a readable label like A1, B2, etc.
func label(p Point) string {
	return fmt.Sprintf("%c%d", 'A'+p.x, p.y+1)
}

// printGrid displays the maze with walls, path, start (Ｓ), and goal (Ｇ)
func printGrid(grid [][]int, start, goal Point, path []Point) {
    n, m := len(grid), len(grid[0]) // Get the dimensions of the grid
    pathSet := make(map[Point]bool) // Create a set to store points in the solution path
    for _, p := range path {        // Populate the set with points from the path
        pathSet[p] = true
    }

    // Print column headers (1, 2, 3, ...)
    fmt.Print("   ")
    for j := 0; j < m; j++ {
        fmt.Printf("%2d ", j+1)
    }
    fmt.Println()

    // Iterate through each row of the grid
    for i := 0; i < n; i++ {
        // Print row label (A, B, C, ...)
        fmt.Printf("%c  ", 'A'+i)
        for j := 0; j < m; j++ {
            p := Point{i, j} // Current point in the grid
            switch {
            case p == start: // If the point is the start position
                fmt.Print("Ｓ")
            case p == goal: // If the point is the goal position
                fmt.Print("Ｇ")
            case pathSet[p]: // If the point is part of the solution path
                // Green colored path using ANSI escape codes
                fmt.Print("\033[32m░░░\033[0m")
            case grid[i][j] == 1: // If the point is a wall
                fmt.Print("███")
            default: // If the point is an empty space
                fmt.Print("   ")
            }
        }
        fmt.Println() // Move to the next row
    }
}

func param(choice int) (int, int, float64, int, int) {
    var n, m int
    var density float64
    var minWalls, minPathLen int

    // Determine maze parameters based on difficulty level
    switch choice {
    case 1: // Easy difficulty
        n, m = 8, 8                // Grid size: 8x8
        density = 0.2              // Wall density: 20%
        minWalls = 10              // Minimum number of walls
        minPathLen = 10            // Minimum path length
    case 2: // Medium difficulty
        n, m = 12, 12              // Grid size: 12x12
        density = 0.4              // Wall density: 40%
        minWalls = 30              // Minimum number of walls
        minPathLen = 15            // Minimum path length
    case 3: // Hard difficulty
        n, m = 20, 20              // Grid size: 20x20
        density = 0.4              // Wall density: 40%
        minWalls = 40              // Minimum number of walls
        minPathLen = 35            // Minimum path length
    case 4: // Custom difficulty
        // Prompt user for custom maze dimensions
        for {
            fmt.Print("Enter maze width (columns, 4 to 26): ")
            fmt.Scanln(&m)
            if m >= 4 && m <= 26 { // Ensure width is within valid range
                break
            }
            fmt.Println("Invalid input. Columns must be between 4 and 26.")
        }
        for {
            fmt.Print("Enter maze height (rows, 4 to 26): ")
            fmt.Scanln(&n)
            if n >= 4 && n <= 26 { // Ensure height is within valid range
                break
            }
            fmt.Println("Invalid input. Rows must be between 1 and 26.")
        }

        // Prompt user for custom wall density
        fmt.Print("Enter wall density (0.0 - 0.6): ")
        fmt.Scanln(&density)
        if density < 0.0 { // Clamp density to minimum of 0.0
            density = 0.0
        } else if density > 0.6 { // Clamp density to maximum of 0.6
            density = 0.6
        }

        // Prompt user for minimum number of walls
        fmt.Print("Enter minimum number of walls: ")
        fmt.Scanln(&minWalls)

        // Prompt user for minimum path length
        fmt.Print("Enter minimum path length: ")
        fmt.Scanln(&minPathLen)
    default:
        // Handle invalid choice by defaulting to Easy difficulty
        fmt.Println("Invalid choice. Defaulting to Easy.")
        n, m = 8, 8                // Grid size: 8x8
        density = 0.2              // Wall density: 20%
        minWalls = 10              // Minimum number of walls
        minPathLen = 10            // Minimum path length
    }

    // Return the determined maze parameters
    return n, m, density, minWalls, minPathLen
}
