package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4, 5}
		numbers2 := []int{1, 2, 3, 4, 5}

		got := Sum(numbers1, numbers2)
		want := [] int {2, 4, 6, 8, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})


}