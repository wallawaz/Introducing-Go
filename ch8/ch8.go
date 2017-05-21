package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	maxVal := flag.Int("max", 6, "the max value")
	flag.Parse()
	randomVal := rand.Intn(*maxVal)
	fmt.Println(randomVal)
}

//when importing other packages, only UpperCase values are accessible from
//other programs/packages.
