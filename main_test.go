package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"positive size", 10},
		{"zero size", 0},
		{"negative size", -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)

			expectedSize := tt.size
			if expectedSize < 0 {
				expectedSize = 0
			}

			if len(result) != expectedSize {
				t.Errorf("expected length %d, got %d",
					expectedSize, len(result))
			}

			for _, value := range result {
				if value < 1 || value > 1_000_000 {
					t.Errorf("invalid random value: %d", value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"one element", []int{5}, 5},
		{"several elements", []int{1, 8, 3, 5}, 8},
		{"maximum at beginning", []int{10, 5, 3}, 10},
		{"maximum at end", []int{2, 4, 9}, 9},
		{"negative numbers", []int{-5, -2, -10}, -2},
		{"equal elements", []int{7, 7, 7}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)

			if got != tt.want {
				t.Errorf("expected %d, got %d", tt.want, got)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"one element", []int{5}, 5},
		{"less than eight elements", []int{1, 9, 3, 4}, 9},
		{"eight elements", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"more than eight elements", []int{1, 5, 3, 9, 2, 7, 4, 6, 10}, 10},
		{"maximum at last element", []int{1, 2, 3, 4, 5, 6, 7, 8, 100}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)

			if got != tt.want {
				t.Errorf("expected %d, got %d", tt.want, got)
			}
		})
	}
}
