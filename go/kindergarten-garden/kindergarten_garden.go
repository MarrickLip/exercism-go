package kindergarten

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Define the Garden type here.

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

type Garden map[string][4]string

var plantTypes = map[string]string{
	"V": "violets",
	"R": "radishes",
	"C": "clover",
	"G": "grass",
}

func hasDuplicates(slice *[]string) bool {
	set := make(map[string]struct{})
	for _, str := range *slice {
		set[str] = struct{}{}
	}

	return len(set) != len(*slice)
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if len(diagram) <= 1 {
		return nil, errors.New("diagram must have at least 1 character")
	}

	if hasDuplicates(&children) {
		return nil, errors.New("children has duplicate names")
	}

	sortedChildren := make([]string, len(children))
	sortedChildren = append(sortedChildren[:0], children...)
	sort.Strings(sortedChildren)

	diagram = diagram[1:]
	rows := strings.Split(diagram, "\n")

	if len(rows) != 2 || len(rows[0]) != len(rows[1]) {
		return nil, errors.New("invalid diagram shape")
	}

	if len(rows[0]) != (len(children) * 2) {
		return nil, errors.New("diagram shape doesn't match number of children")
	}

	garden := make(Garden)

	for i, name := range sortedChildren {
		rawPlants := [4]string{}

		rawPlants[0] = string(rows[0][2*i+0])
		rawPlants[1] = string(rows[0][2*i+1])
		rawPlants[2] = string(rows[1][2*i+0])
		rawPlants[3] = string(rows[1][2*i+1])

		plants := [4]string{}
		for j, rawPlant := range rawPlants {
			plant, ok := plantTypes[rawPlant]
			if !ok {
				return nil, fmt.Errorf("invalid plant type: %v", rawPlant)
			}
			plants[j] = plant
		}

		garden[name] = plants
	}

	return &garden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	if g == nil {
		return nil, false
	}

	val, ok := (*g)[child]

	if ok {
		return val[:], true
	} else {
		return nil, false
	}
}
