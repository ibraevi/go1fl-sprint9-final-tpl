package main

import (
	"fmt"
	"math"
	"slices"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-10, 0, 10, 1000} {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			data := generateRandomElements(size)
			wantLength := max(size, 0)
			if len(data) != wantLength {
				t.Fatalf("len(data) = %d, want %d", len(data), wantLength)
			}
			for i, value := range data {
				if value <= 0 {
					t.Errorf("data[%d] = %d, want a positive integer", i, value)
				}
			}
		})
	}
}

func TestMaximumAlgorithms(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil", data: nil, want: 0},
		{name: "empty", data: []int{}, want: 0},
		{name: "single", data: []int{100500}, want: 100500},
		{name: "zero", data: []int{0}, want: 0},
		{name: "single_negative", data: []int{-100500}, want: -100500},
		{name: "first", data: []int{9, 3, 1}, want: 9},
		{name: "middle", data: []int{1, 9, 3}, want: 9},
		{name: "last", data: []int{1, 3, 9}, want: 9},
		{name: "duplicates", data: []int{5, 5, 5, 5}, want: 5},
		{name: "negative_short", data: []int{-9, -2, -5}, want: -2},
		{name: "negative_chunks", data: []int{-9, -8, -7, -6, -5, -4, -3, -2, -1}, want: -1},
		{name: "integer_limits", data: []int{math.MinInt, 0, math.MaxInt}, want: math.MaxInt},
		{name: "minimum_integer", data: []int{math.MinInt}, want: math.MinInt},
	}
	algorithms := []struct {
		name string
		find func([]int) int
	}{
		{name: "maximum", find: maximum},
		{name: "maxChunks", find: maxChunks},
	}

	for _, algorithm := range algorithms {
		t.Run(algorithm.name, func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					data := slices.Clone(test.data)
					if got := algorithm.find(data); got != test.want {
						t.Errorf("maximum = %d, want %d", got, test.want)
					}
					if !slices.Equal(data, test.data) {
						t.Errorf("input changed: got %v, want %v", data, test.data)
					}
				})
			}
		})
	}
}

func TestMaxChunksAllPositions(t *testing.T) {
	for size := 1; size <= 33; size++ {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			data := make([]int, size)
			for position := range data {
				data[position] = 100500
				if got := maxChunks(data); got != 100500 {
					t.Fatalf("maximum at position %d: got %d, want 100500", position, got)
				}
				data[position] = 0
			}
		})
	}
}

func TestMaximumLargeSlice(t *testing.T) {
	data := make([]int, 100_003)
	for i := range data {
		data[i] = -1
	}
	data[len(data)-1] = math.MaxInt
	if got := maximum(data); got != math.MaxInt {
		t.Errorf("maximum = %d, want %d", got, math.MaxInt)
	}
	if got := maxChunks(data); got != math.MaxInt {
		t.Errorf("maxChunks = %d, want %d", got, math.MaxInt)
	}
}
