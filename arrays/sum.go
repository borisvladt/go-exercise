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
