package wordy

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func pop(slice *[]string) (string, error) {
	if len(*slice) == 0 {
		return "", errors.New("can't pop from empty slice")
	}
	next := (*slice)[0]
	*slice = (*slice)[1:]
	return next, nil
}

func popNum(slice *[]string) (int, error) {
	next, err := pop(slice)
	if err != nil {
		return 0, err
	}

	num, err := strconv.ParseInt(next, 10, 64)
	return int(num), err
}

func Answer(question string) (int, bool) {
	// extract raw question
	re := regexp.MustCompile(`What is (.*)\?`)
	matches := re.FindAllStringSubmatch(question, -1)

	if len(matches) == 0 || len(matches[0]) == 0 {
		return 0, false
	}

	question = matches[0][1]
	question = strings.ReplaceAll(question, "multiplied by", "multiplied")
	question = strings.ReplaceAll(question, "divided by", "divided")

	parts := strings.Split(question, " ")

	firstPart, err := pop(&parts)
	if err != nil {
		return 0, false
	}

	firstNum, err := strconv.ParseInt(firstPart, 10, 32)
	if err != nil {
		return 0, false
	}

	result := int(firstNum)

	for len(parts) != 0 {
		nextPart, _ := pop(&parts)

		switch nextPart {
		case "plus":
			nextNum, err := popNum(&parts)
			if err != nil {
				return 0, false
			}
			result = result + nextNum
		case "minus":
			nextNum, err := popNum(&parts)
			if err != nil {
				return 0, false
			}
			result = result - nextNum
		case "multiplied":
			nextNum, err := popNum(&parts)
			if err != nil {
				return 0, false
			}
			result = result * nextNum
		case "divided":
			nextNum, err := popNum(&parts)
			if err != nil {
				return 0, false
			}
			result = result / nextNum
		default:
			return 0, false
		}
	}

	return result, true
}
