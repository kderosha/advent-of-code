package universe

import (
	"fmt"
	"log/slog"

	"github.com/kderosha/advent-of-code/2023/10/graph"
)

type Universe struct {
	galaxies          []*Galaxy
	galaxiesForRow    []int
	galaxiesForColumn []int
}

type Galaxy struct {
	coord graph.Point
}

func (g *Galaxy) String() string {
	return fmt.Sprintf("%+v", g.coord)
}

func NewGalaxy(x, y int) *Galaxy {
	return &Galaxy{
		coord: graph.Point{x, y},
	}
}

func NewUniverse(lines [][]byte) *Universe {
	galaxiesForRow := make([]int, len(lines))
	galaxiesForColumn := make([]int, len(lines[0]))
	universe := &Universe{
		galaxiesForRow:    galaxiesForRow,
		galaxiesForColumn: galaxiesForColumn,
	}
	galaxies := make([]*Galaxy, 0)
	for y, line := range lines {
		for x, symbol := range line {
			if symbol == '#' {
				galaxiesForColumn[x]++
				galaxiesForRow[y]++
				galaxies = append(galaxies, NewGalaxy(x, y))
			}
		}
	}
	universe.galaxies = galaxies
	slog.Info("Galaxies are processed", "galaxies", galaxies)
	return universe
}

func (u Universe) CalculateSumsOfDistanceBetweenGalaxies() int {
	slog.Info("CalcSums")
	slog.Info("Calculating sum of all distances between pairs of galaxies", "galaxies", u.galaxies)
	sum := 0
	for x, currentGalaxy := range u.galaxies {
		for _, distantGalaxy := range u.galaxies[x+1:] {
			slog.Info("Calculating distances between galaxies", "currentGalaxy", currentGalaxy, "distantGalaxy", distantGalaxy)
			currentGalaxyRealPosition := u.getRealGalaxyPosition(currentGalaxy)
			distantGalaxyRealPosition := u.getRealGalaxyPosition(distantGalaxy)
			slog.Info("Calculating distances between real positions", "currentPosition", currentGalaxyRealPosition, "distantGalaxyPosition", distantGalaxyRealPosition)
			sum += graph.DistanceBetweenPoints(currentGalaxyRealPosition, distantGalaxyRealPosition)
		}
	}
	return sum
}

func (u Universe) getRealGalaxyPosition(g *Galaxy) graph.Point {
	additionalXValue := 0
	additionalYValue := 0
	for x, value := range u.galaxiesForColumn {
		if value == 0 && x < g.coord.X() {
			additionalXValue++
		}
	}
	for y, value := range u.galaxiesForRow {
		if value == 0 && y < g.coord.Y() {
			additionalYValue++
		}
	}
	return graph.Point{(g.coord.X() - additionalXValue) + 1 + (additionalXValue * 1000000), (g.coord.Y() - additionalYValue) + 1 + (additionalYValue * 1000000)}
}
