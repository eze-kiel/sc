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

type stack []float64

func main() {
	var stack []float64
	r := bufio.NewReader(os.Stdin)

	for {
		showStack(stack)

		input, err := getInput(r)
		if err != nil && err == io.EOF {
			os.Exit(0)
		} else if err != nil {
			logrus.Errorf("error parsing output: %s", err)
			continue
		}

		switch input {
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
			quit()
		case "c", "clear":
			stack = []float64{}
		default:
			n, err := strconv.ParseFloat(input, 64)
			if err != nil {
				logrus.Errorf("error casting input to float: %s", err)
				continue
			}

			stack = append(stack, n)
		}
	}
}

func getInput(r *bufio.Reader) (string, error) {
	rawInput, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Trim(rawInput, "\n"), nil
}

func showStack(s []float64) {
	fmt.Println("\n-- stack")
	if len(s) == 0 {
		fmt.Println("   empty")
	} else {
		for idx, val := range s {
			fmt.Printf("%d:\t%f\n", idx, val)
		}
	}
	fmt.Printf("-- end\n\n")
}
