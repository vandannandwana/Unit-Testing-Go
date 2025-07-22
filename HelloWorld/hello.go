package main

import "fmt"

const prefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return prefix + name
}

func main() {

	s := "Vandan, Nandwana\n"

	fmt.Printf("%s", s)

}
