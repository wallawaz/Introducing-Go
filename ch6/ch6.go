package main

import "fmt"

func main() {
	numbers := []float64{100.0, 200.0, 300.0, 444.32, 343.11}
	fmt.Println(average(numbers))

	//python * equivilent
	total := variadic_add(numbers...)
	fmt.Println("total", total)
}

//(xs <param><type>) <return type>
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	if total > 0.0 {
		return total / float64(len(xs))
	}
	return total
}

func variadic_add(args ...float64) float64 {
	total := 0.0
	for _, v := range args {
		total += v
	}
	return total
}
