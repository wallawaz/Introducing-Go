package main

import "fmt"

func main() {

	// maps need to be initialized
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
	fmt.Println(x["key"])

	check_map()
	shorter_map_creation()
}

func check_map() {
	elem := make(map[string]string)
	elem["C"] = "Carbon"
	elem["He"] = "Helium"

	// if `el` in elem-> ok == true
	if el, ok := elem["C"]; ok {
		fmt.Println(ok, el)
	}
}

func shorter_map_creation() {
	elem := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
	}
	if el, ok := elem["H"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
