package visualise

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/CSCoursework/tp-astar/astar"
)

var (
	colourBackground = color.RGBA{0, 0, 0, 0xff}
	colourUnused     = color.RGBA{50, 49, 49, 0xff}
	colourUsed       = color.RGBA{190, 52, 58, 0xff}

	cellWidth  = 10 // pixels
	cellHeight = 10 // pixels

	dividerWidth = 1 // pixels
)

func calculateCellBoundaries(cellNumber, cellSize, borderSize int) (int, int) {
	borders := borderSize * cellNumber
	return (cellSize * cellNumber) + borders, (cellSize * (cellNumber + 1)) + borders
}

func Maze(maze astar.Maze, path []astar.Position) {

	mazeHeight := len(maze)
	mazeWidth := len(maze[len(maze)-1])

	imageWidth := (cellWidth * mazeWidth) + (dividerWidth * (mazeWidth - 1))
	imageHeight := (cellHeight * mazeHeight) + (dividerWidth * (mazeHeight - 1))

	// This is used so we can do an O(1) lookup of positions that are on the tracked path
	positionMap := make(map[astar.Position]struct{})
	for _, pos := range path {
		positionMap[pos] = struct{}{}
	}

	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{imageWidth, imageHeight}})

	widthMod := cellWidth + dividerWidth
	heightMod := cellHeight + dividerWidth

	for y := 0; y < imageHeight; y += 1 {
		for x := 0; x < imageWidth; x += 1 {

			// cellBorder variables represent the cell number of the current pixel including the cell and border size
			cellBorderX := x / widthMod
			cellBorderY := y / heightMod

			// Maximum and minimum X and Y values for a cell without the border
			cellMinX, cellMaxX := calculateCellBoundaries(cellBorderX, cellWidth, dividerWidth)
			cellMinY, cellMaxY := calculateCellBoundaries(cellBorderY, cellHeight, dividerWidth)

			pixColour := colourBackground

			if (x >= cellMinX && x < cellMaxX) && (y >= cellMinY && y < cellMaxY) { // this means we're in a cell's space and not a border

				pos := astar.Position{cellBorderX, cellBorderY}

				if maze.GetPosition(pos) { // this is wall
					// walls are the same colour as the background
					// do nothing
				} else if _, onPath := positionMap[pos]; onPath { // this is on the found path
					pixColour = colourUsed
				} else { // this is a random cell that's not used
					pixColour = colourUnused
				}

			}

			img.Set(x, y, pixColour)

		}
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)

}
