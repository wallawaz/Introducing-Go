package main

import "fmt"
import "ch8/math"

func main() {
	xs := []float64{1, 3, 4, 45}
	avg := math.Average(xs)
	fmt.Println(avg)
}
