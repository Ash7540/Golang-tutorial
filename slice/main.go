package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slice Study")
	// Slices are like arrays but with dynamic length
	var fruitList = []string{"apple", "peach", "mango"}
	fmt.Printf("Type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "banana", "orange")
	// fmt.Println("Fruit List: ", fruitList)

	fruitList = append(fruitList[1:3])
	// fmt.Println("Fruit List: ", fruitList)

	// method II
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777
	// highScores = append(highScores, 555, 666, 777)
	// fmt.Println("High Scores: ", highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	// fmt.Println("Sorted Scores: ", highScores)

	// removing the value from slice based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("Courses: ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses: ", courses)
}

