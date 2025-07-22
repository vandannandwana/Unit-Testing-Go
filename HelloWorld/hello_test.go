package main

import (
	"testing"

)

func assertCorrectMessage(t testing.TB, got, want string){

	t.Helper()
	if got != want{
		t.Errorf("got %q want %q", got, want)
	}

}


func TestHello(t *testing.T) {
	
	t.Run("Saying Hello to Vandan", func (t * testing.T){
		got := Hello("Vandan")
		want := "Hello, Vandan"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("When input string is empty", func (t * testing.T){
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})



}