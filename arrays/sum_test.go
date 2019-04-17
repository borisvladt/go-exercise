package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of five numbers as array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,2}, []int{0,9})
	want := []int{3, 9}
	checkSums(t, got, want)
}

func TestSumAll2(t *testing.T) {
	got := SumAll2([]int{1,2}, []int{0,2,5,8,13}, []int{5,11,25})
	want := []int{3, 28, 41}
	checkSums(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of slices", func(t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{0,2,5,8,13}, []int{5,11,25})
		want := []int{2, 28, 36}
		checkSums(t, got, want)
	})

	t.Run("sum with an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{0,2,5,8,13}, []int{})
		want := []int{2, 28, 0}
		checkSums(t, got, want)
	})

	t.Run("sum with an empty slice", func(t *testing.T) {
		got := SumAllTails2([]int{1,2}, []int{0,2,5,8,13}, []int{})
		want := []int{2, 28}
		checkSums(t, got, want)
	})
}

func checkSums(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
