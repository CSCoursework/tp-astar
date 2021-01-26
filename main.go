package main

import (
	"fmt"
	"github.com/CSCoursework/tp-astar/astar"
	"github.com/CSCoursework/tp-astar/visualise"
)

func main() {
	maze := astar.Maze{{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false}}

	start := astar.Position{0, 0}
	end := astar.Position{7, 6}

	path := astar.Astar(maze, start, end)

	visualise.Maze(maze, path)

	fmt.Println(path)
}
