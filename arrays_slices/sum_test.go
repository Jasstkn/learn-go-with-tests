package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T)  {
	t.Run("collection of any length", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("sum of 2 slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given", got, want )
		}
	})

	t.Run("sum of 1 slices", func(t *testing.T) {
		got := SumAll([]int{1, 1, 1})
		want := []int{3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given", got, want)
		}
	})

	t.Run("sum of tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v given", got, want)
		}
	})

	t.Run("sum of tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v given", got, want)
		}
	})
}