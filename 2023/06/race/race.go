package race

type BoatRace struct {
	time           int
	distanceRecord int
}

func NewBoatRace(time, distanceRecord int) *BoatRace {
	return &BoatRace{time, distanceRecord}
}

func (br *BoatRace) PossibleWaysToWin() []int {
	var possibleWins []int = make([]int, 0)
	for x := 0; x < br.time; x++ {
		distance := calculateDistance(br.time, x)
		if distance > br.distanceRecord {
			possibleWins = append(possibleWins, distance)
		}
	}
	return possibleWins
}

// The formula to calculate distance traveled: distance = (raceTime - buttonHoldTime) * buttonHoldTime
func calculateDistance(raceTime int, buttonHoldTime int) int {
	return (raceTime - buttonHoldTime) * buttonHoldTime
}
