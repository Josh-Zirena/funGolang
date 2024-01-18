package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var result = addNumbers(3, 5)
	fmt.Println("The answer is", result)

	var rangeOfEvenNumbers = evenNumbers(12, 20)
	fmt.Println("Even numbers in our range is", rangeOfEvenNumbers)

}

func addNumbers(a int, b int) int {
	return a + b
}

func evenNumbers(a int, b int) []int {
	var rangeNums []int
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			rangeNums = append(rangeNums, i)
		}
	}

	// fmt.Printf("%v", rangeNums)
	return rangeNums
}
