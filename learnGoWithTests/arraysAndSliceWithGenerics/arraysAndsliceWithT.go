package arraysandslicewithgenerics

func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

func Reduce[A any](collection []A, f func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

func Find[A any](collection []A, f func(A) bool) (value A, found bool) {
	for _, x := range collection {
		if f(x) {
			return x, true
		}
	}
	return
}

func SumAll(numberss ...[]int) []int {
	var sums []int
	for _, numbers := range numberss {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(numbers, sumTail, []int{})
}

// func SumAllTails(numberss ...[]int) []int {
// 	var sums []int
// 	for _, numbers := range numberss {
// 		if len(numbers) == 0 {
// 			sums = append(sums, 0)
// 			continue
// 		} else {
// 			sums = append(sums, Sum(numbers[1:]))
// 		}
// 	}
// 	return sums
// }
