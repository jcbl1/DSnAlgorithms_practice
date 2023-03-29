package sololearn_new_drivers_license

import (
	"fmt"
)

func Main() {
	myName, agents, others := getInput()
	names := append(others, myName)
	names = sortAlpha(names)

	fmt.Println(timeConsumption(names, agents, myOrder(myName, names)))
}

func getInput() (myName string, agents int, others []string) {
	others = make([]string, 4)
	fmt.Scanln(&myName)
	fmt.Scanln(&agents)
	for i := range others {
		fmt.Scan(&others[i])
	}
	return
}

func sortAlpha(names []string) []string {
	alphabet := "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	rank := func(c rune) int {
		for i, v := range alphabet {
			if c == v {
				return i / 2
			}
		}
		return -1
	}
	aBiggerThanB := func(a, b string) bool {
		switch {
		case len(a) >= len(b):
			for i := range b {
				if rank(rune(a[i])) > rank(rune(b[i])) {
					return true
				} else if rank(rune(a[i])) < rank(rune(b[i])) {
					return false
				} else {
					continue
				}
			}
			return true
		case len(a) < len(b):
			for i := range a {
				if rank(rune(a[i])) > rank(rune(b[i])) {
					return true
				} else if rank(rune(a[i])) < rank(rune(b[i])) {
					return false
				} else {
					continue
				}
			}
			return false
		}
		return false
	}

	for i := range names {
		for j := i + 1; j < len(names); j++ {
			if aBiggerThanB(names[i], names[j]) {
				k := names[j]
				names[j] = names[i]
				names[i] = k
			}
		}
	}
	return names
}

func myOrder(myName string, names []string) int {
	for i, v := range names {
		if myName == v {
			return i
		}
	}
	return -1
}

func timeConsumption(names []string, agents, myOrder int) (tc int) {
	tc = (myOrder/agents + 1) * 20
	return
}
