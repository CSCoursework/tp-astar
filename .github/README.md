# tp-astar

An implementation of the A* pathfinding algorithm (translated from the code supplied in [this](https://medium.com/@nicholas.w.swift/easy-a-star-pathfinding-7e6689c7f7b2) article) in Go, with an additional image generator.

(I'm just a little bit proud of the fact that I came up with the relevant maths in my head and got the code working first try for the image)

![Example visualisation](https://raw.githubusercontent.com/CSCoursework/tp-astar/master/.github/image.png)

## To run:

1. Create a maze in `main.go` and define a start and end point. `true` denotes a wall, `false` denotes an empty cell.
2. Run `main.go`.
3. Observe the path printed to the console and the generated image (blue is the start, green is the end and red is the path taken).