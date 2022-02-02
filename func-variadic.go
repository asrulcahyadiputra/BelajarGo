package main

import "fmt"

func sumAll(numbers ...int) int{
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main(){
	total := sumAll(10,10,10,10)
	fmt.Println(total)
	// menggunakan slice
	numbersSlice := []int{10,10,10}
	total = sumAll(numbersSlice...)
	fmt.Println(total)
}