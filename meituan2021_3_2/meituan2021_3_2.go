package meituan2021_3_2

import (
	"fmt"
)

type Order struct {
	idx   int
	v, w  int
	total int
}

func Main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)
	ors := make([]Order, n)
	for i := range ors {
		fmt.Scan(&ors[i].v, &ors[i].w)
		ors[i].idx = i
		ors[i].total = getEarning(ors[i].v, ors[i].w)
	}

	sortByTotal(&ors)

	m1, m2 := m, m
	for ors[m2-1].total == ors[m2].total {
		m2++
	}
	for m1-2 >= 0 && ors[m1-2].total == ors[m1-1].total {
		m1--
	}

	tail := ors[m1-1 : m2]
	sortByIdx(&tail)
	topMs := append(ors[:m1-1], tail[:m-m1+1]...)

	sortByIdx(&topMs)
	for i := range topMs {
		fmt.Printf("%d ", topMs[i].idx+1)
	}
	fmt.Println()

}

func getEarning(v, w int) int {
	return v + w*2
}
func getEarnings(v, w []int) []int {
	earnings := make([]int, len(v))
	for i := range v {
		earnings[i] = getEarning(v[i], w[i])
	}
	return earnings
}
func sortByTotal(ors *[]Order) {
	for i := range *ors {
		for j := i + 1; j < len(*ors); j++ {
			if (*ors)[i].total < (*ors)[j].total {
				k := (*ors)[j]
				(*ors)[j] = (*ors)[i]
				(*ors)[i] = k
			}
		}
	}
}

func sortByIdx(ors *[]Order) {
	for i := range *ors {
		for j := i + 1; j < len(*ors); j++ {
			if (*ors)[i].idx > (*ors)[j].idx {
				k := (*ors)[j]
				(*ors)[j] = (*ors)[i]
				(*ors)[i] = k
			}
		}
	}
}
