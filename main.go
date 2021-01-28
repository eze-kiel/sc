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

	for {
		showStack(stack)

		input, err := getInput(r)
		if err != nil && err == io.EOF {
			os.Exit(0)
		} else if err != nil {
			logrus.Errorf("error parsing input: %s", err)
			continue
		}

		for _, in := range input {
			switch in {
			case "+", "a", "add":
				stack = add(stack)
			case "-", "s", "sub":
				stack = sub(stack)
			case "*", "m", "mul":
				stack = mul(stack)
			case "/":
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
			default:
				n, err := strconv.ParseFloat(in, 64)
				if err != nil {
					logrus.Errorf("error casting input to float: %s", err)
					continue
				}

				stack = append(stack, n)
			}
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
	for idx, val := range s {
		fmt.Printf("%d:\t%f\n", idx, val)
	}
	fmt.Println()
}
