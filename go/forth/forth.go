package forth

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func getUnaryOperand(stack *[]int) (int, error) {
	if len(*stack) >= 1 {
		num := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		return num, nil
	} else {
		err := fmt.Errorf("can't perform a unary operation, stack only has %v elements", len(*stack))
		return 0, err
	}
}

func getBinaryOperands(stack *[]int) ([2]int, error) {
	if len(*stack) >= 2 {
		num1 := (*stack)[len(*stack)-2]
		num2 := (*stack)[len(*stack)-1]

		*stack = (*stack)[:len(*stack)-2]

		return [2]int{num1, num2}, nil
	} else {
		err := fmt.Errorf("can't perform a binary operation, stack only has %v elements", len(*stack))
		return [2]int{}, err
	}
}

// arithmetic
func add(stack *[]int) error {
	nums, err := getBinaryOperands(stack)
	if err == nil {
		*stack = append(*stack, nums[0]+nums[1])
		return nil
	} else {
		return err
	}
}

func sub(stack *[]int) error {
	nums, err := getBinaryOperands(stack)
	if err == nil {
		*stack = append(*stack, nums[0]-nums[1])
		return nil
	} else {
		return err
	}
}

func mul(stack *[]int) error {
	nums, err := getBinaryOperands(stack)
	if err == nil {
		*stack = append(*stack, nums[0]*nums[1])
		return nil
	} else {
		return err
	}
}

func div(stack *[]int) error {
	nums, err := getBinaryOperands(stack)
	if err == nil {
		if nums[1] == 0 {
			return errors.New("can't divide by 0")
		}
		*stack = append(*stack, nums[0]/nums[1])
		return nil
	} else {
		return err
	}
}

// stack operations
func dup(stack *[]int) error {
	num, err := getUnaryOperand(stack)
	if err == nil {
		*stack = append(*stack, num, num)
		return nil
	} else {
		return err
	}
}

func drop(stack *[]int) error {
	_, err := getUnaryOperand(stack)
	return err
}

func swap(stack *[]int) error {
	nums, err := getBinaryOperands(stack)
	if err == nil {
		*stack = append(*stack, nums[1], nums[0])
		return nil
	} else {
		return err
	}
}

func over(stack *[]int) error {
	nums, err := getBinaryOperands(stack)
	if err == nil {
		*stack = append(*stack, nums[0], nums[1], nums[0])
		return nil
	} else {
		return err
	}
}

func tryParseNum(input string) (int, bool) {
	num, err := strconv.ParseInt(input, 10, 64)
	if err == nil {
		return int(num), true
	} else {
		return 0, false
	}
}

func Forth(input []string) ([]int, error) {
	stack := []int{}
	words := map[string][]string{}

	wordPattern := regexp.MustCompile(`^: (\S*) (.*) ;$`)
	wordDefinitions := input[:len(input)-1]
	for _, wordDefinition := range wordDefinitions {
		matches := wordPattern.FindAllStringSubmatch(wordDefinition, -1)
		if len(matches[0]) != 3 {
			return stack, fmt.Errorf("invalid word definition: '%v'", wordDefinition)
		}

		wordName := strings.ToLower(matches[0][1])

		if _, ok := tryParseNum(wordName); ok {
			return stack, fmt.Errorf("can't redefine %v", matches[0][1])
		}

		words[wordName] = strings.Split(matches[0][2], " ")
	}

	fmt.Printf("words: %v\n", words)

	rawOps := strings.Split(input[len(input)-1], " ")
	for _, rawOp := range rawOps {
		ops, ok := words[strings.ToLower(rawOp)]
		if !ok {
			ops = []string{rawOp}
		}
		fmt.Println(rawOp, ok, ops)
		for _, op := range ops {
			if num, ok := tryParseNum(rawOp); ok {
				stack = append(stack, num)
				continue
			}

			var err error = nil
			switch strings.ToLower(op) {
			case "+":
				err = add(&stack)
			case "-":
				err = sub(&stack)
			case "*":
				err = mul(&stack)
			case "/":
				err = div(&stack)
			case "dup":
				err = dup(&stack)
			case "drop":
				err = drop(&stack)
			case "swap":
				err = swap(&stack)
			case "over":
				err = over(&stack)
			default:
				err = fmt.Errorf("unhandled opearation: %v", op)
			}

			if err != nil {
				return stack, err
			}
		}

	}

	return stack, nil
}
