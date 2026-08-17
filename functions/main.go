package main

import ("fmt")

type transformFunction func(int)int

// type complexSignatureFunction func(int, []string, map[string][]int) ([]int, string)

func main () {
	fmt.Println("Functions")
	numbers := []int{1,2,3,4,5}
	moreNumbers := []int{6,7,8,9,10}

	// doubled[0] = 100

	transformerFn1 := getTranformerFunction(&numbers)
	transformerFn2 := getTranformerFunction(&moreNumbers)
	doubled := transformNumbers(&numbers, transformerFn1)
	tripled := transformNumbers(&numbers, transformerFn2)

	fmt.Println(numbers)
	fmt.Println(doubled)
	fmt.Println(tripled)

	double := createTransformer(2)
	triple := createTransformer(3)

	doubled = transformNumbers(&numbers, double)
	tripled = transformNumbers(&numbers, triple)

	fmt.Println("after createTransformer:", doubled)
	fmt.Println("after createTransformer:", tripled)

	fmt.Println("Factorial", factorial(5))
}

func createTransformer(factor int) func(int) int {
	// factory function
	return func (number int) int {
		return number * factor
	}
}

func getTranformerFunction(numbers *[]int) transformFunction {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func transformNumbers(numbers *[]int, transform transformFunction) []int {
	// in this form, its a new array
	dNumbers := []int{}
	for _ , v := range *numbers {
		dNumbers = append(dNumbers, transform(v))
	}
	return dNumbers
}

func double (value int) int {
	return value * 2
}
func triple (value int) int {
	return value * 3
}

func factorial (number int) int {
	if number == 0 {
		return 1
	}
	return factorial(number - 1) * number
}

// func doubleNumbers(numbers *[]int) []int {
// 	// in this form, its a new array
// 	dNumbers := []int{}
// 	for _ , v := range *numbers {
// 		dNumbers = append(dNumbers, v * 2)
// 	}
// 	return dNumbers
// }

// func doubleNumbers(numbers []int) []int {
	// in this form, i'm copying everytime
// 	for k, v := range numbers {
// 		numbers[k] = v * 2
// 	}
// 	return numbers
// }