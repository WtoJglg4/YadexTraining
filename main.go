package main

import (
	"fmt"
	"github/YandexTraining/set"
)

func main() {

	// linear.Triangle()
	// testing.Testing()
	setSize := 10
	mySet := set.NewSet(setSize)
	for i := 50; i < 60; i++ {
		mySet.Add(i)
	}
	for i := 45; i < 65; i++ {
		fmt.Println(mySet.Find(i), i)
	}
	fmt.Println(mySet)
	mySet.Delete(50)
	mySet.Delete(55)
	mySet.Delete(60)
	fmt.Println(mySet)
}
