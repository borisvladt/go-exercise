package arrays

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	//for i:=0; i<5; i++ {
	//	sum += numbers[i]
	//}

	for key, value := range numbers {
		sum += value
		fmt.Print(key)
	}

	fmt.Println("")
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfArgs := len(numbersToSum)
	sums := make([]int, lengthOfArgs)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func SumAll2(numbersToSum ...[]int) []int {
	var sums []int
	for _, value := range numbersToSum {
		sums = append(sums, Sum(value))
	}

	return sums
}
