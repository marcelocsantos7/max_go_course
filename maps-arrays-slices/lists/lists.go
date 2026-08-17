package lists

import (
	"fmt"
)

// type Product struct {
// 	id    int
// 	title string
// 	price float64
// }

// // Practice
// func main() {
// 	var hobbies [3]string = [3]string{"Guitar", "Programing", "Gaming"}
// 	fmt.Println(hobbies)

// 	fmt.Println(hobbies[0])
// 	fmt.Println(hobbies[1:]) // [1:3]

// 	bestHobbies := hobbies[0:2]
// 	bestHobbies2 := hobbies[:2]

// 	fmt.Println(bestHobbies)
// 	fmt.Println(bestHobbies2)

// 	bestHobbies = hobbies[:]
// 	fmt.Println(bestHobbies)
// 	fmt.Println(cap(bestHobbies))

// 	myGoals := []string{"Learn the basics,", "Learn Everything"}
// 	fmt.Println(myGoals)
// 	fmt.Println(cap(myGoals))

// 	myGoals[1] = "Learn almost everything"
// 	fmt.Println(myGoals)

// 	myGoals = append(myGoals, "Learn the core concepts")
// 	fmt.Println(myGoals)

// 	var products []Product = []Product{
// 		{1, "Produto1", 10.00},
// 		{2, "Produto2", 11.00},
// 	}
// 	fmt.Println(products)

// 	products = append(products, Product{3, "Product3", 12.00})
// 	fmt.Println(products)
// }

// ############################################################

func main() {
	prices := []float64{10.99, 9.99}
	fmt.Println(prices[1])

	prices = append(prices, 5.99, 6.99, 7.99, 8.99) // adding N values separated by commas
	fmt.Println(prices)

	discountPrices := []float64{2.99, 3.99, 4.99}
	prices = append(prices, discountPrices...) // like a spread operator in javascrip

	fmt.Println(prices)
}

// ############################################################

// func main() {
// 	fmt.Println("Slices, Maps and Arrays")

// 	var productNames [4]string = [4]string{"A Book"}

// 	prices := [4]float64{11.1, 22.2, 33.33, 44.44}
// 	fmt.Println("prices original ", prices)

// 	productNames[2] = "A Carpet"

// 	fmt.Println(prices[0])
// 	fmt.Println(productNames)

// 	pricesSlice := prices[1:3] // slice: first element included, last excluded
// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[1:]

// 	fmt.Println("prices slice prices[1:3]:", pricesSlice)

// 	fmt.Println("Featured ", featuredPrices)
// 	fmt.Println("prices updated", prices)
// 	fmt.Println("Highlighted", highlightedPrices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// }
