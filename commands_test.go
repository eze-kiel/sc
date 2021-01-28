package main

import (
	"reflect"
	"testing"
)

func Test_add(t *testing.T) {
	tests := []struct {
		name string
		s    []float64
		want []float64
	}{
		{name: "long enough", s: []float64{12, 24}, want: []float64{36}},
		{name: "too short", s: []float64{12}, want: []float64{12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sub(t *testing.T) {
	tests := []struct {
		name string
		s    []float64
		want []float64
	}{
		{name: "long enough", s: []float64{30, 10}, want: []float64{20}},
		{name: "too short", s: []float64{12}, want: []float64{12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sub(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mul(t *testing.T) {
	tests := []struct {
		name string
		s    []float64
		want []float64
	}{
		{name: "long enough", s: []float64{30, 10}, want: []float64{300}},
		{name: "too short", s: []float64{12}, want: []float64{12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mul(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_div(t *testing.T) {
	tests := []struct {
		name string
		s    []float64
		want []float64
	}{
		{name: "long enough", s: []float64{30, 10}, want: []float64{3}},
		{name: "too short", s: []float64{12}, want: []float64{12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := div(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swap(t *testing.T) {
	tests := []struct {
		name string
		s    []float64
		want []float64
	}{
		{name: "long enough", s: []float64{30, 10}, want: []float64{10, 30}},
		{name: "too short", s: []float64{12}, want: []float64{12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swap(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pop(t *testing.T) {
	tests := []struct {
		name string
		s    []float64
		want []float64
	}{
		{name: "long enough", s: []float64{30, 10}, want: []float64{30}},
		{name: "too short", s: []float64{}, want: []float64{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pop(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pow(t *testing.T) {
	tests := []struct {
		name string
		s    []float64
		want []float64
	}{
		{name: "long enough", s: []float64{5, 10}, want: []float64{9765625}},
		{name: "too short", s: []float64{5}, want: []float64{5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	tests := []struct {
		name string
		s    []float64
		want []float64
	}{
		{name: "long enough", s: []float64{5, 10, 12}, want: []float64{27}},
		{name: "too short", s: []float64{5}, want: []float64{5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqrt(t *testing.T) {
	tests := []struct {
		name string
		s    []float64
		want []float64
	}{
		{name: "long enough", s: []float64{4}, want: []float64{2}},
		{name: "too short", s: []float64{}, want: []float64{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sqrt(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
