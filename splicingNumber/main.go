package main

import (
	"fmt"
	"strconv"
)

func SplicingNumber(number *[]int) int {
	str := ""
	for _, v := range *number {
		tmp := strconv.Itoa(v)
		str += tmp
	}

	result, _ := strconv.Atoi(str)
	return result
}

func main() {
	test := []int{11, 21, 33}
	fmt.Println(SplicingNumber(&test))
}
