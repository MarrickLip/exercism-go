package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type TeamResults struct {
	Name   string
	Wins   int
	Losses int
	Draws  int
}

func (teamResults *TeamResults) CalculatePoints() int {
	return teamResults.Wins*3 + teamResults.Draws
}

func (teamResults *TeamResults) CalculateMatchesPlayed() int {
	return teamResults.Wins + teamResults.Draws + teamResults.Losses
}

func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)

	teamResults := map[string]TeamResults{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			continue
		}

		_, ok := teamResults[parts[0]]
		if !ok {
			teamResults[parts[0]] = TeamResults{Name: parts[0], Wins: 0, Losses: 0, Draws: 0}
		}

		_, ok = teamResults[parts[1]]
		if !ok {
			teamResults[parts[1]] = TeamResults{Name: parts[1], Wins: 0, Losses: 0, Draws: 0}
		}

		teamA := teamResults[parts[0]]
		teamB := teamResults[parts[1]]

		switch parts[2] {
		case "win":
			teamA.Wins++
			teamB.Losses++
		case "loss":
			teamA.Losses++
			teamB.Wins++
		case "draw":
			teamA.Draws++
			teamB.Draws++
		}

		teamResults[parts[0]] = teamA
		teamResults[parts[1]] = teamB
	}

	allResults := []TeamResults{}
	for _, results := range teamResults {
		allResults = append(allResults, results)
	}

	sort.Slice(allResults, func(i, j int) bool {
		return allResults[i].CalculatePoints() > allResults[j].CalculatePoints()
	})

	fmt.Fprintf(writer, "Team                            | MP |  W |  D |  L |  P\n")
	for _, result := range allResults {
		fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n", result.Name, result.CalculateMatchesPlayed(), result.Wins, result.Losses, result.Draws, result.CalculatePoints())
	}

	return nil
}
