package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	var stack []float64
	r := bufio.NewReader(os.Stdin)

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// The input is piped
		var rawInput []rune
		for {
			raw, _, err := r.ReadRune()
			if err != nil && err == io.EOF {
				break
			}
			rawInput = append(rawInput, raw)
		}
		strInput := strings.Replace(string(rawInput), "\n", "", -1)
		input := strings.Split(strInput, " ")
		stack = rangeOverCommands(stack, input)

		// We only display one element.
		// As it is piped, we expect to have one result
		fmt.Println(stack[len(stack)-1])
	} else {

		for {
			showStack(stack)

			input, err := getInput(r)
			if err != nil && err == io.EOF {
				os.Exit(0)
			} else if err != nil {
				logrus.Errorf("error parsing input: %s", err)
				continue
			}

			stack = rangeOverCommands(stack, input)
		}
	}
}

func getInput(r *bufio.Reader) ([]string, error) {
	rawInput, err := r.ReadString('\n')
	if err != nil {
		return []string{}, err
	}

	rawInput = strings.Trim(rawInput, "\n")
	return strings.Split(rawInput, " "), nil
}

func showStack(s []float64) {
	for _, val := range s {
		fmt.Printf("\t%f\n", val)
	}
	fmt.Println()
}

func rangeOverCommands(stack []float64, input []string) []float64 {
	for _, in := range input {
		switch in {
		case "+", "a", "add":
			stack = add(stack)
		case "-", "s", "sub":
			stack = sub(stack)
		case "*", "m", "mul":
			stack = mul(stack)
		case "/", "d", "div":
			stack = div(stack)
		case "sw", "swap":
			stack = swap(stack)
		case "p", "pop":
			stack = pop(stack)
		case "q", "quit":
			os.Exit(0)
		case "c", "clear":
			stack = []float64{}
		case "^", "pow":
			stack = pow(stack)
		case "sum":
			stack = sum(stack)
		default:
			n, err := strconv.ParseFloat(in, 64)
			if err != nil {
				logrus.Errorf("error casting input to float: %s", err)
				continue
			}

			stack = append(stack, n)
		}
	}

	return stack
}
