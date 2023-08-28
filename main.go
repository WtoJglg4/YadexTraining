package main

import (
	"fmt"
	set "github/YandexTraining/sets"
)

func main() {

	// linear.Triangle()
	// testing.Testing()

	mySet := set.NewSet(10)
	for i := 50; i < 60; i++ {
		mySet.Add(i)
	}
	for i := 45; i < 65; i++ {
		fmt.Println(mySet.Find(i))
	}
}
