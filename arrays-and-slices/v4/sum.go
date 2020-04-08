package main

func Sum(numbers []int) (total int) {
	for _, number := range numbers {
		total += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}
