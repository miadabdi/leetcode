package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSumIteration3(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3}, 7, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
		{[]int{1}, 1, []int{-1, -1}},
	}

	for _, tt := range tests {
		got := TwoSumIteration3(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("TwoSumIteration3(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

func BenchmarkTwoSumIteration3(b *testing.B) {
	short := []int{2, 7, 11, 15}

	b.Run("short", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TwoSumIteration3(short, 9)
		}
	})

	long := make([]int, 10000)
	for i := range long {
		long[i] = i
	}
	long[2689] = -1
	long[9876] = -2
	answer := -3

	b.Run("long", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TwoSumIteration3(long, answer)
		}
	})
}
