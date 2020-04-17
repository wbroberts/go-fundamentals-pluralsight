package main

import "fmt"

var (
	state string
	gpa   float32
)

func main() {
	name := "Will"
	course := "Fundamentals"
	fmt.Println(name, "is enrolled in", course)

	// & references the variable's spot in memory (a pointer)
	changeCourseToLearningPointers(&course)

	fmt.Println(name, "is now enrolled in", course)
}

// * dereferences the pointer
func changeCourseToLearningPointers(course *string) string {
	*course = "Learning Pointers"
	fmt.Println("Attempting to changing your course to", *course)
	return *course
}
