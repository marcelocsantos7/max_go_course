package main

import (
	"fmt"
)

// func Keyword: Begins the method declaration.
// Receiver (r ReceiverType): Sitting between func and the method name, this
// 		defines which custom type owns the method.
// Method Name: Follows Go's visibility rules (Capitalized = Public/Exported,
//  	Lowercase = Package-Private).
// Parameters & Return Types: Works identically to standard Go functions,
// 		supporting multiple return values.

type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)
}

func main() {
	fmt.Println("Maps, Slices and Arrays")
	userNames := make([]string, 2, 5)
	// userNames := []string{}

	userNames[0] = "Lili"
	// userNames[1] = "Malu"
	userNames = append(userNames, "Marcelo")
	userNames = append(userNames, "Max")
	// fmt.Println(userNames)

	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.8
	courseRatings["react"] = 4.6
	courseRatings["angular"] = 4.9
	// fmt.Println(courseRatings)
	// courseRatings.output()

	for index, value := range userNames {
		fmt.Println(index, value)
	}
	for key, value := range courseRatings {
		fmt.Println(key, value)
	}
}
