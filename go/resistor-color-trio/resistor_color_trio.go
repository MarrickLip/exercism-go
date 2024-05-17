package resistorcolortrio

import (
	"fmt"
	"strconv"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").

var colourToValue = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func calculateBase(colors []string) int {
	firstDigit := colourToValue[colors[0]]
	secondDigit := colourToValue[colors[1]]
	return firstDigit*10 + secondDigit
}

func Label(colors []string) string {
	base := calculateBase(colors)
	exponent := colourToValue[colors[2]]

	if base == 0 {
		return "0 ohms"
	}

	// if second digit is "0"
	if colors[1] == "black" {
		exponent += 1
		base /= 10
	}

	unit := "ohms"
	if exponent >= 9 {
		unit = "gigaohms"
		exponent -= 9
	} else if exponent >= 6 {
		unit = "megaohms"
		exponent -= 6
	} else if exponent >= 3 {
		unit = "kiloohms"
		exponent -= 3
	}

	fmt.Println(colors, base, exponent, unit)

	result := strconv.Itoa(base)
	for i := 0; i < exponent; i++ {
		result += "0"
	}
	result = result + " " + unit
	return result

}
