package main

import (
	"fmt"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxProfit(prices []int, k int) int {
	n := len(prices)

	if n < 2 {
		return 0
	}

	if k == 0 {
		return 0
	}

	if k >= n/2 {
		maxProfit := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				maxProfit += (prices[i] - prices[i-1])
				// fmt.Println("max profit :", maxProfit)
			}
		}
		return maxProfit
	}

	result := make([][]int, k+1)
	// fmt.Println(result)
	for i := range result {
		result[i] = make([]int, n)
	}
	// fmt.Println(result)

	for j := 1; j <= k; j++ {
		maxDiff := -prices[0]
		for x := 1; x < n; x++ {
			maxDiff = Max(maxDiff, result[j-1][x-1]-prices[x])
			result[j][x] = Max(result[j][x-1], maxDiff+prices[x])
			// fmt.Println("maxDiff :", maxDiff)
		}

	}

	// fmt.Println(result)

	return result[k][n-1]
}

func main() {
	fmt.Println(MaxProfit([]int{2, 7, 10, 20, 30, 0, 15}, 3)) // 43
	fmt.Println(MaxProfit([]int{2, 7, 10, 20, 30, 0, 15}, 1)) // 28
	fmt.Println(MaxProfit([]int{2, 4, 1}, 10))                // 2
	fmt.Println(MaxProfit([]int{2, 4, 1}, 2))                 // 2
	fmt.Println(MaxProfit([]int{3, 2, 6, 5, 0, 30}, 2))       //34
	fmt.Println(MaxProfit([]int{3, 2, 6, 5, 0, 30}, 1))       //30
	fmt.Println(MaxProfit([]int{3, 2, 6, 5, 0, 3}, 2))        //7
}
