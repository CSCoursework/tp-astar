package main

import (
	"fmt"
	"github.com/CSCoursework/tp-astar/astar"
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
	fmt.Println(path)
}
