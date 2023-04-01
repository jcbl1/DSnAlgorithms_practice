package meituan2021_3_1

import (
	"fmt"
)

func Main() {
	n := 0
	fmt.Scanln(&n)
	weights := make([]int, n)
	sequence := make([]int, n)
	for i := range weights {
		fmt.Scan(&weights[i])
	}
	for i := range weights {
		fmt.Scan(&sequence[i])
		sequence[i]--
	}

	for i := range sequence {
		x := 0
		weights, x = bigger(weights, sequence[i])
		fmt.Println(x)
	}
}

func bigger(weights []int, takeout int) ([]int, int) {
	addTogether := func(nums []int) int {
		at := 0
		for _, v := range nums {
			at += v
		}
		return at
	}
	weights[takeout] = 0
	fmt.Println(weights, takeout)
	if takeout == len(weights)-1 {
		return weights, addTogether(weights)
	}
	a := weights[:takeout]
	b := weights[takeout+1:]
	aWeight := addTogether(a)
	bWeight := addTogether(b)
	if aWeight >= bWeight {
		return weights, aWeight
	}
	return weights, bWeight
}
