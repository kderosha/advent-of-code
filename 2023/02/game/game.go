package game

type Game struct {
	id int
	rounds []Round
}

func NewGame(game string) (Game, error) {
	colonIndex := strings.IndexByte(game, ':')
	gameSubstring := game[:colonIndex]
	roundsSubstring := game[colonIndex + 1:]
	slog.Info("game processed", "game", game, "colonIndex", colonIndex, "gameSubstring", gameSubstring, "roundsSubstring", roundsSubstring)

	// get id
	id, err := getGameId(gameSubstring)
	if err != nil {
		return Game{}, err
	}

	// create new rounds
	rounds := createRounds(roundsSubstring)

	return Game{
		id: id,
		rounds: rounds,
	}, nil
}

// Create rounds from a substring of format N color, N color, N color...;
func createRounds(roundsSubstring string) []Round {
	// split the string into an array of strings by semi colon
	rounds := strings.Split(roundsSubstring, ";")
	numOfRounds := len(rounds)
	slog.Info("Processed rounds substring", "rounds", rounds, "numOfRounds", numOfRounds)
	newRounds := make([]Round, 0)
	for _, round := range rounds {
		newRounds = append(newRounds, NewRound(round))
	}
	return newRounds
}

// Parse out the game id from the game string
func getGameId(game string) (int, error) {
	// split string
	if idString, found := strings.CutPrefix(game, "Game "); found {
		if id, err := strconv.Atoi(idString); err != nil {
			slog.Error("Error parsing game id", "error", err.Error(), "idString", idString)
			return 0, err
		} else {
			return id, nil
		}
	}
	return 0, fmt.Errorf("Game id not found for game %s", game)
}

type Round struct {
	roundString string
	colors map[string]int
}

func (r Round) Possible(p Puzzle) bool {
	// Get red color from Round compare to red limit
	if redColor, ok := r.colors["red"] ; ok {
		if redColor > p.redLimit {
			return false
		}
	}

	if blueColor, ok := r.colors["blue"]; ok {
		if blueColor > p.blueLimit {
			return false
		}
	}

	if greenColor, ok := r.colors["green"] ; ok {
		if greenColor > p.greenLimit {
			return false
		}
	}

	return true
}
