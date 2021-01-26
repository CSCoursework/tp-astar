package astar

var mooreNeighbourhood = []Position{{0, -1}, {0, 1}, {-1, 0}, {1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

type Maze [][]bool
type Position [2]int

func (m Maze) GetPosition(pos Position) bool {
	return m[pos[0]][pos[1]]
}

type node struct {
	parent   *node
	position Position

	g, h, f int
}

func newNode(parent *node, position Position) *node {
	return &node{
		parent:   parent,
		position: position,
	}
}

func (n *node) equalTo(other *node) bool {
	return n.position == other.position
}

// Astar returns a list of Positions that represent a path from a given start and end point in a given maze
func Astar(maze Maze, start, end Position) []Position {

	startNode := newNode(nil, start)
	endNode := newNode(nil, end)

	var (
		openList   = []*node{startNode}
		closedList []*node
	)

	for len(openList) > 0 {
		// find the current node 
		currentNode := openList[0]
		currentIndex := 0
		for index, item := range openList {
			if item.f < currentNode.f {
				currentNode = item
				currentIndex = index
			}
		}

		// remove the current node from the open list and add it to the closed list
		{
			// https://github.com/golang/go/wiki/SliceTricks#delete
			copy(openList[currentIndex:], openList[currentIndex+1:])
			openList[len(openList)-1] = nil
			openList = openList[:len(openList)-1]
		}
		closedList = append(closedList, currentNode)

		if currentNode.equalTo(endNode) { // found the target node
			var path []Position
			current := currentNode
			for current != nil {
				path = append(path, current.position)
				current = current.parent
			}
			// this crazy thing reverses path - https://stackoverflow.com/a/28058324
			for i, j := 0, len(path) - 1; i < j; i, j = i + 1, j - 1 {
				path[i], path[j] = path[j], path[i]
			}
			return path
		}

		// generate children
		var children []*node
		for _, delta := range mooreNeighbourhood {
			nodePosition := Position{currentNode.position[0] + delta[0], currentNode.position[1] + delta[1]}

			// ensure the new position is in range
			if nodePosition[0] > len(maze)-1 || nodePosition[0] < 0 || 
				nodePosition[1] > len(maze[len(maze)-1])-1 || nodePosition[1] < 0 {
				
				continue
			}

			// make sure nodePosition isn't a wall
			if maze.GetPosition(nodePosition) {
				continue
			}

			// create the new node and append that
			children = append(children, newNode(currentNode, nodePosition))
		}

		for _, child := range children {
			// child is on the closed list
			for _, cChild := range closedList {
				if child.equalTo(cChild) {
					continue
				}
			}

			// set f, g and h
			child.g = currentNode.g + 1
			ax := child.position[0] - endNode.position[0]
			ay := child.position[1] - endNode.position[1]
			child.h = (ax*ax) + (ay*ay)
			child.f = child.h + child.g

			// child is already in the open list
			for _, oNode := range openList {
				if child.equalTo(oNode) && child.g > oNode.g {
					continue
				}
			}

			// add the child to the open list
			openList = append(openList, child)
		}

	}

	return []Position{}
}
