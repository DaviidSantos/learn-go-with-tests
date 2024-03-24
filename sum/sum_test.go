package sum

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	numSlice1 := []int{1, 2}
	numSlice2 := []int{0, 9}
	got := SumAll(numSlice1, numSlice2)
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v, given %v and %v", got, want, numSlice1, numSlice2)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("sum all tails", func(t *testing.T) {
		numSlice1 := []int{1, 2}
		numSlice2 := []int{0, 9}
		got := SumAllTails(numSlice1, numSlice2)
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("sum all tails when passed an empty slice", func(t *testing.T) {
		numSlice1 := []int{1, 2}
		numSlice2 := []int{}
		got := SumAllTails(numSlice1, numSlice2)
		want := []int{2, 0}

		checkSums(t, got, want)
	})
}
