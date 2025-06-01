package main

import "container/heap"

// Point represents a coordinate in the maze grid (x = row, y = column)
type Point struct{ x, y int }

// Node represents a cell in A* with costs (g, h, f), parent for backtracking, and index for heap
type Node struct {
	pt      Point // Coordinate of the node
	g, h, f int   // g = cost from start, h = heuristic to goal, f = g + h
	parent  *Node // Pointer to parent node (used to reconstruct path)
	index   int   // Index in the priority queue (required by heap.Interface)
}



// PriorityQueue is a min-heap of *Node ordered by f value (lowest first)
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].f < pq[j].f
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	node.index = len(*pq)
	*pq = append(*pq, node)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.index = -1
	*pq = old[:n-1]
	return node
}
func (pq *PriorityQueue) update(node *Node, g, h int, parent *Node) {
	node.g, node.h, node.f, node.parent = g, h, g+h, parent
	heap.Fix(pq, node.index)
}
