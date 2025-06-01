# RandomMazeSolver
This project was created as part of the ITCS440 Intelligent Systems course at the University of Bahrain.

#Overview
This program generates random mazes of varying difficulty levels and finds a solution path through them using the A* algorithm. Users can select from predefined difficulty levels or customize their own maze parameters.

#Features
Choose between Easy, Medium, Hard, or Custom maze difficulty
Custom option allows setting maze dimensions, wall density, minimum walls, and minimum path length
Maze is randomly generated based on selected parameters
A* pathfinding algorithm solves the maze, returning the shortest solution path
Visual representation of the maze grid with walls, start, goal, and solution path displayed

#Usage
Upon running the program, the user is prompted to choose a difficulty level:
Easy: 8x8 grid, 20% walls, minimum 10 walls, minimum path length 10
Medium: 12x12 grid, 40% walls, minimum 30 walls, minimum path length 15
Hard: 20x20 grid, 40% walls, minimum 40 walls, minimum path length 35
Custom: User inputs maze size, wall density, minimum walls, and path length

The program then generates a random maze meeting these constraints and prints both the maze and the solution path.

#How it works
Maze is represented as a 2D grid with walls and free spaces
Start point is randomly placed on the first row (first half)
Goal point is randomly placed on the last row (second half)
A* algorithm finds the shortest path from start to goal avoiding walls
Solution path is printed step-by-step and highlighted in the grid
