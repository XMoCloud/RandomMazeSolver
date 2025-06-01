package main

import "container/heap"

// heuristic calculates the Manhattan distance between two points
func heuristic(a, b Point) int {
	dx := a.x - b.x
	if dx < 0 {
		dx = -dx
	}
	dy := a.y - b.y
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

// aStar implements the A* pathfinding algorithm
func aStar(grid [][]int, start, goal Point) ([]Point, bool) {
	n, m := len(grid), len(grid[0])
	open := make(map[Point]*Node)  // Nodes to be evaluated
	closed := make(map[Point]bool) // Already evaluated nodes
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Initialize start node
	startNode := &Node{pt: start, g: 0, h: heuristic(start, goal)}
	startNode.f = startNode.h
	heap.Push(pq, startNode)
	open[start] = startNode

	// Directions for movement (up, down, left, right)
	dirs := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for pq.Len() > 0 {
		// Get the node with the lowest f score
		curr := heap.Pop(pq).(*Node)

		// If goal reached, reconstruct the path
		if curr.pt == goal {
			var path []Point
			for c := curr; c != nil; c = c.parent {
				path = append(path, c.pt)
			}
			// Reverse the path to start from the beginning
			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}
			return path, true
		}

		// Mark current node as evaluated
		delete(open, curr.pt)
		closed[curr.pt] = true

		// Explore all valid neighbors
		for _, d := range dirs {
			nx, ny := curr.pt.x+d.x, curr.pt.y+d.y
			p := Point{nx, ny}

			// Skip if out of bounds, wall, or already evaluated
			if nx < 0 || nx >= n || ny < 0 || ny >= m || grid[nx][ny] == 1 || closed[p] {
				continue
			}

			ng := curr.g + 1 // cost from start to neighbor

			// If neighbor is not in open list, create and add it
			if node, ok := open[p]; !ok {
				nh := heuristic(p, goal)
				node = &Node{pt: p, g: ng, h: nh, f: ng + nh, parent: curr}
				heap.Push(pq, node)
				open[p] = node
			} else if ng < node.g {
				// If found a shorter path to the neighbor, update it
				pq.update(node, ng, node.h, curr)
			}
		}
	}
	return nil, false // No path found
}
