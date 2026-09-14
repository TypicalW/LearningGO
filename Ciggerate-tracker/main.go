package main

import "fmt"

func main() {
	ciggs := Cigarettes{}

	ciggs.add("Kings", "Gold Flake")

	fmt.Println(ciggs)
}
