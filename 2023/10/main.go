package main

import (
	"bytes"
	"log/slog"
	"os"

	"github.com/kderosha/advent-of-code/2023/10/graph"
)

func main() {

	fileBytes, err := os.ReadFile("puzzle_input.txt")
	if err != nil {
		slog.Error("Error reading file")
		panic(err)
	}
	lines := bytes.Split(fileBytes, []byte("\n"))

	// Matrix points are referenced y,x
	nodeMatrix := make([][]*graph.Node, len(lines))
	var startingNode *graph.Node
	// Make the nodeMatrix
	for y, line := range lines {
		nodeMatrix[y] = make([]*graph.Node, len(line))
		for x, lineByte := range line {
			// Create new graph from coordinate and byte. byte determines the direction of the edges
			newNode := graph.NewNode(x, y, rune(lineByte))
			nodeMatrix[y][x] = newNode
			if newNode.IsStartingNode() {
				startingNode = newNode
			}
		}
	}
	slog.Info("node matrix has been created", "matrix", nodeMatrix, "startingNode", startingNode)
	visitedNodes := make([]*graph.Node, 0)
	visitedNodes = append(visitedNodes, startingNode)
	startingNode.Visit()
	node1 := nodeMatrix[startingNode.Point().North().Y()][startingNode.Point().North().X()]
	visitedNodes = append(visitedNodes, node1)
	node1.Visit()
	previousNode := startingNode
	for !node1.Equal(startingNode) {
		nextNodePoint := node1.Move(previousNode)
		previousNode = node1
		node1 = nodeMatrix[nextNodePoint.Y()][nextNodePoint.X()]
		node1.Visit()
		visitedNodes = append(visitedNodes, node1)
		slog.Info("Calculated next node1 based on the current node and the previous known node1", "next node", node1)
	}
	slog.Info("Steps to get to the same node", "answer", len(visitedNodes)/2, "visitedNodes", visitedNodes)

	// Ray tracing solution
	area := 0
	for _, line := range nodeMatrix {
		inside := false
		for _, node := range line {
			if node.Visited {
				// flip the inside
				inside = !inside
			}
			if inside {
				area++
			}
		}
	}

	slog.Info("Completed analysis of the second part.", "area", area)
}
