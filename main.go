package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/CSCoursework/tp-astar/astar"
	"github.com/CSCoursework/tp-astar/visualise"
)

func main() {
	// maze := astar.Maze{{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false},
	// {false, true, false, true, true, true, true, true, true, true, false, true, true, true, true, true, true, true, false, true, true, true, true, true, false, true, true, true, true, true, true, true, true, true, false, true, false},
	// {false, true, false, true, false, false, false, false, false, true, false, true, false, false, false, true, false, false, false, false, false, true, false, false, false, true, false, false, false, true, false, false, false, true, false, true, false},
	// {false, true, true, true, true, true, true, true, false, true, true, true, false, true, false, true, true, true, true, true, false, true, true, true, true, true, false, true, false, true, false, true, true, true, false, true, false},
	// {false, false, false, false, false, true, false, true, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, true, false, true, false, true, false, true, false, false, false, false, false, true, false},
	// {false, true, true, true, false, true, false, true, true, true, true, true, true, true, false, true, true, true, true, true, false, true, false, true, false, true, false, true, false, true, true, true, true, true, false, true, false},
	// {false, true, false, true, false, false, false, false, false, false, false, true, false, true, false, true, false, false, false, false, false, true, false, false, false, true, false, true, false, false, false, false, false, true, false, true, false},
	// {false, true, false, true, true, true, true, true, true, true, false, true, false, true, false, true, true, true, false, true, true, true, true, true, false, true, true, true, true, true, false, true, true, true, false, true, false},
	// {false, true, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, true, false, true, false, false, false, true, false, false, false, false, false, true, false, true, false, false, false, true, false},
	// {false, true, true, true, true, true, false, true, true, true, false, true, false, true, true, true, false, true, false, true, true, true, false, true, true, true, true, true, false, true, false, true, true, true, true, true, false},
	// {false, true, false, false, false, true, false, false, false, true, false, true, false, false, false, true, false, true, false, false, false, false, false, true, false, true, false, true, false, true, false, false, false, false, false, false, false},
	// {false, true, true, true, false, true, true, true, false, true, false, true, false, true, true, true, false, true, true, true, false, true, true, true, false, true, false, true, false, true, true, true, false, true, true, true, false},
	// {false, false, false, false, false, false, false, true, false, true, false, true, false, true, false, false, false, false, false, true, false, true, false, false, false, true, false, true, false, false, false, true, false, false, false, true, false},
	// {false, true, true, true, true, true, false, true, false, true, true, true, true, true, false, true, true, true, false, true, false, true, true, true, false, true, false, true, true, true, false, true, true, true, true, true, false},
	// {false, true, false, false, false, true, false, true, false, false, false, false, false, true, false, false, false, true, false, true, false, false, false, true, false, false, false, true, false, false, false, false, false, false, false, true, false},
	// {false, true, false, true, true, true, false, true, true, true, true, true, false, true, true, true, true, true, false, true, true, true, false, true, false, true, false, true, true, true, true, true, true, true, false, true, false},
	// {false, true, false, true, false, false, false, false, false, true, false, true, false, false, false, true, false, false, false, false, false, true, false, true, false, true, false, false, false, false, false, false, false, true, false, true, false},
	// {false, true, false, true, true, true, true, true, false, true, false, true, false, true, false, true, false, true, true, true, true, true, false, true, true, true, true, true, true, true, false, true, false, true, true, true, false},
	// {false, true, false, false, false, false, false, true, false, true, false, true, false, true, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false, true, false, true, false, false, false, false, false},
	// {false, true, false, true, true, true, false, true, false, true, false, true, true, true, true, true, false, true, true, true, true, true, false, true, false, true, true, true, false, true, true, true, true, true, true, true, false},
	// {false, true, false, true, false, false, false, true, false, false, false, true, false, false, false, true, false, false, false, false, false, true, false, false, false, true, false, true, false, false, false, false, false, true, false, true, false},
	// {false, true, false, true, true, true, true, true, true, true, true, true, false, true, true, true, true, true, false, true, true, true, true, true, true, true, false, true, true, true, true, true, true, true, false, true, false},
	// {false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}}

	// start := astar.Position{22, 1}
	// end := astar.Position{0, 35}

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
	fmt.Println("Found path:", path)

	buf := new(bytes.Buffer)
	visualise.Maze(maze, path, buf)
	fmt.Println("Generated image")

	err := ioutil.WriteFile("image.png", buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Saved image")

}
