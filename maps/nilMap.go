package main

import (
	"fmt"
)

func main() {
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println("aMap:", aMap)
	aMap = nil
	fmt.Println("aMap:", aMap)
	if aMap == nil {
		fmt.Println("nil map!")
	}
	aMap["test"] = 1
	fmt.Println(aMap)
}
