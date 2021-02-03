package main

import (
	"fmt"
	"math"
)

func add(s []float64) []float64 {
	if len(s) > 1 {
		x, y := s[len(s)-1], s[len(s)-2]
		s = s[:len(s)-2]
		s = append(s, x+y)
	}
	return s
}

func sub(s []float64) []float64 {
	if len(s) > 1 {
		x, y := s[len(s)-1], s[len(s)-2]
		s = s[:len(s)-2]
		s = append(s, y-x)
	}
	return s
}

func mul(s []float64) []float64 {
	if len(s) > 1 {
		x, y := s[len(s)-1], s[len(s)-2]
		s = s[:len(s)-2]
		s = append(s, x*y)
	}
	return s
}

func div(s []float64) []float64 {
	if len(s) > 1 {
		x, y := s[len(s)-1], s[len(s)-2]
		s = s[:len(s)-2]
		s = append(s, y/x)
	}
	return s
}

func swap(s []float64) []float64 {
	if len(s) > 1 {
		x, y := s[len(s)-1], s[len(s)-2]
		s[len(s)-1], s[len(s)-2] = y, x
	}
	return s
}

func pop(s []float64) []float64 {
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

func pow(s []float64) []float64 {
	if len(s) > 1 {
		x, y := s[len(s)-1], s[len(s)-2]
		s = s[:len(s)-2]
		s = append(s, math.Pow(y, x))
	}
	return s
}

func sum(s []float64) []float64 {
	if len(s) > 1 {
		var sum float64
		for _, v := range s {
			sum += v
		}
		return []float64{sum}
	}
	return s
}

func sqrt(s []float64) []float64 {
	if len(s) > 0 && s[len(s)-1] != 0 {
		x := s[len(s)-1]
		s = s[:len(s)-1]
		s = append(s, math.Sqrt(x))
	}
	return s
}

func help() {
	fmt.Println(`
--- Help

Basic:
  +, a, add
  	add last 2 values of the stack
  
  -, s, sub
  	subtract last 2 values of the stack
  
  *, m, mul
  	multiply last 2 values of the stack

  /, d, div
  	divide last 2 values of the stack

Advanced:
  ^, pow
  	power of the last 2 values of the stack

  sum
  	sum all the stack

  sqrt
  	square root the last value of the stack

Functions:
  sw, swap
	swap last two values of the stack

  p, pop
  	pop last value of the stack

  c, clear
	clear the stack

  q, quit
  	quit the program
  
  h, help
	display help
---

`)
}
