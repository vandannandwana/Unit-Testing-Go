package main

import "testing"

func TestAdd(t *testing.T) {

	got := Add(2, 5)
	want := 7

	if got != want{
		t.Errorf("Required %d Get %d", want, got)
	}

}