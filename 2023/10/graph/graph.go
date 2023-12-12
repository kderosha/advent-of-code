package graph

import "math"

type Node struct {
	Symbol       rune
	coordPoint   Point   // Stores the coordinate for this node
	edges        []Point // linked to a different points in the matrix
	startingNode bool
	groundNode   bool
	Visited      bool
}

func NewNode(x, y int, symbol rune) *Node {
	graphNode := &Node{
		Symbol:     symbol,
		coordPoint: Point{x, y},
	}
	//calculate edge coordinates
	if symbol == '|' {
		// edgeCoordinates are north and south
		graphNode.edges = append(graphNode.edges, graphNode.coordPoint.North())
		graphNode.edges = append(graphNode.edges, graphNode.coordPoint.South())
	} else if symbol == '-' {
		graphNode.edges = append(graphNode.edges, graphNode.coordPoint.East())
		graphNode.edges = append(graphNode.edges, graphNode.coordPoint.West())
	} else if symbol == 'L' {
		graphNode.edges = append(graphNode.edges, graphNode.coordPoint.North())
		graphNode.edges = append(graphNode.edges, graphNode.coordPoint.East())
	} else if symbol == 'J' {
		graphNode.edges = append(graphNode.edges, graphNode.coordPoint.North())
		graphNode.edges = append(graphNode.edges, graphNode.coordPoint.West())
	} else if symbol == '7' {
		graphNode.edges = append(graphNode.edges, graphNode.coordPoint.South())
		graphNode.edges = append(graphNode.edges, graphNode.coordPoint.West())
	} else if symbol == 'F' {
		graphNode.edges = append(graphNode.edges, graphNode.coordPoint.South())
		graphNode.edges = append(graphNode.edges, graphNode.coordPoint.East())
	} else if symbol == 'S' {
		// Starting node
		graphNode.startingNode = true
	} else {
		// ground node
		graphNode.groundNode = true
	}
	return graphNode
}

// Two graphs are considered equal if they are at the same coordinate.
func (n *Node) Equal(another *Node) bool {
	return n.coordPoint.Equals(another.coordPoint)
}

func (n *Node) IsStartingNode() bool {
	return n.startingNode
}

func (n *Node) String() string {
	return string(n.Symbol)
}

func (n *Node) Point() Point {
	return n.coordPoint
}

func (n *Node) Move(previousNode *Node) Point {
	for _, edge := range n.edges {
		if !edge.Equals(previousNode.coordPoint) {
			return edge
		}
	}
	panic("Error couldn't find edge to follow that didn't lead to the previous node")
}

func (n *Node) Visit() {
	n.Visited = true
}

type Point [2]int

func (p Point) North() Point {
	return Point{p.X(), p.Y() - 1}
}

func (p Point) East() Point {
	return Point{p.X() + 1, p.Y()}
}

func (p Point) South() Point {
	return Point{p.X(), p.Y() + 1}
}

func (p Point) West() Point {
	return Point{p.X() - 1, p.Y()}
}

func (p Point) Equals(another Point) bool {
	return p.X() == another.X() && p.Y() == another.Y()
}

// Get the x coordinate for the Point
func (p Point) X() int {
	return p[0]
}

func (p Point) Y() int {
	return p[1]
}

func DistanceBetweenPoints(p1, p2 Point) int {
	xDistance := int(math.Abs(float64(p2.X() - p1.X())))
	yDistance := int(math.Abs(float64(p2.Y() - p1.Y())))
	return xDistance + yDistance
}
