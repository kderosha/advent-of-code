package main

import (
	"bytes"
	"log/slog"
	"os"
	"regexp"

	"github.com/kderosha/advent-of-code/2023/08/node"
)

var nodeMatcher *regexp.Regexp = regexp.MustCompile(`(.+) = \((.+), (.+)\)`)

func main() {
	// Read in the file bytes
	fileBytes, err := os.ReadFile("puzzle_input.txt")
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(fileBytes, []byte("\n"))

	// Parse first line into directions
	directions := make([]int, len(lines[0]))
	for x, direction := range lines[0] {
		directionString := string(direction)
		if directionString == "L" {
			directions[x] = 0
		} else {
			directions[x] = 1
		}
	}
	slog.Info("Directional array has been processed", "array", directions)
	nodeLines := lines[2:]
	nodeMap := make(map[string]*node.Node, len(nodeLines))
	// Process the nodes from now on
	for _, nodeLine := range nodeLines {
		// Parse out the lines with regex
		matched := nodeMatcher.FindSubmatch(nodeLine)
		slog.Info("Processed node line", "Original string", string(matched[0]), "first group", string(matched[1]), "second group", string(matched[2]), "third matched", string(matched[3]))
		// New node added to map
		nodeMap[string(matched[1])] = &node.Node{
			Root:  string(matched[1]),
			Left:  string(matched[2]),
			Right: string(matched[3]),
		}
	}

	ghostTown := &node.GhostTown{
		DirectionArray: directions,
		NodeMap:        nodeMap,
	}
	slog.Info("Part 1 answer", "value", ghostTown.Traverse("AAA"))
}
