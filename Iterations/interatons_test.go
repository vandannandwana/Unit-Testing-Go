package main

import "testing"

func BenchmarkRepeat(b *testing.B){

	for b.Loop(){
		Repeat("a")
	}

}