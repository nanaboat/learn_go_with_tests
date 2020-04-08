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

func SumAllTails(numbersToSum ...[]int) (sumTails []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sumTails = append(sumTails, 0)
		} else {
			sumTails = append(sumTails, Sum(numbers[1:]))
		}
	}
	return
}
