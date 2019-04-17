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

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}

func SumAllTails2(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) > 0 {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}
