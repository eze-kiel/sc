package main

import "math"

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
