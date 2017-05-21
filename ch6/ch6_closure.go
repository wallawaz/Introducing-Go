package main

import "fmt"

func main() {
	even := makeEvenGenerator()
	for i := 1; i <= 10; i++ {
		fmt.Println(even())
	}

}

//returning a function
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
