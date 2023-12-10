package node

import (
	"fmt"
	"log/slog"
)

type GhostTown struct {
	NodeMap        map[string]*Node
	DirectionArray []int
}

// Given a starting node for the map. Traverse the tree using direction array until you reach ZZZ
// Output the number of steps it took to get to ZZZ
func (gt *GhostTown) Traverse(startingKey string) int {
	node, exists := gt.NodeMap[startingKey]
	if !exists {
		slog.Error("Invalid starting node. Not found in node map", "nodeMap", gt.NodeMap)
	}
	steps := 0
	directionArrayIndex := 0
	slog.Info("Starting to traverse")
	for node.Root != "ZZZ" {
		// slog.Info("current node being evaluated", "node", node, "direction", gt.DirectionArray[directionArrayIndex])
		if gt.DirectionArray[directionArrayIndex] == 0 {
			node = gt.NodeMap[node.Left]
		} else {
			node = gt.NodeMap[node.Right]
		}

		steps++
		directionArrayIndex++
		// Restart at the beginning of the direction array and continue
		if directionArrayIndex == len(gt.DirectionArray) {
			directionArrayIndex = 0
		}
	}
	return steps
}

type Node struct {
	Root  string
	Left  string
	Right string
}

func (n *Node) Step(nodeMap map[string]*Node, direction int) *Node {
	if direction == 0 {
		return nodeMap[n.Left]
	} else {
		return nodeMap[n.Right]
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("%+v", *n)
}
