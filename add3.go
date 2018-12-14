package main

import "fmt"

func getIndex(param int) int {
	if param <= 0 {
		return 0
	} else if param == 1 {
		return 1
	} else if param == 2 {
		return 2
	}
	return getIndex(param - 1) + getIndex(param - 2)
}

func getResult(param int) int {
	if param <= 0 {
		return 0
	} else if param == 1 {
		return 1
	} else if param == 2 {
		return 3
	}
	return getResult(param - 1) + getIndex(param)
}

func main() {
	num := 6
	fmt.Printf("number index: %d\nthe number is: %d\nsumup result is: %d\n", num, getIndex(num), getResult(num))
}
