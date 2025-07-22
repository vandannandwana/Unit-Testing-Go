package main

import "testing"

func TestParameter(t *testing.T){

	t.Run("Getting Parameter for Rectangle", func(t *testing.T){
		var rect Rectangle
		got := rect.Parameter(4,5)
		var want float32 = 18

		if got != want{
			t.Errorf("Got: %f Want: %f", got, want)
		}

	})
	t.Run("Getting Parameter for Circle", func(t *testing.T){
		var circ Circle
		got := circ.Parameter(4,4)
		var want float32 = 50.24
		if got != want{
			t.Errorf("Got: %f Want: %f", got, want)
		}

	})

}