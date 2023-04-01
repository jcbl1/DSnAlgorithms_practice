package meituan_train_inNout

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var randInt int

type StationStack []int

func (s *StationStack) Push(x any) {
	*s = append(*s, x.(int))
}
func (s *StationStack) Pop() any {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

func Main() {
	randInt = rand.Intn(60000)

	in, out := getInOut()
	fmt.Println(checkInOutV2(in, out))
}

func getInOut() (in, out []int) {
	scnr := bufio.NewScanner(os.Stdin)
	ints := make([][]int, 2)
	fmt.Println("请输入入站和出站队列（两行、整数、以空格隔开）:")
	for i := 0; i < 2; i++ {
		scnr.Scan()
		splited := strings.Split(scnr.Text(), " ")
		for _, v := range splited {
			n, _ := strconv.Atoi(v)
			ints[i] = append(ints[i], n)
		}
	}
	in = ints[0]
	out = ints[1]
	return
}

func checkInOutV2(in, out []int) bool {
	if len(in) != len(out) {
		return false
	}
	sh := new(StationStack)
	sh.Push(randInt + 1e9)
	count := 0
	for i := range in {
		sh.Push(in[i])
		for j := count; j < len(in); j++ {
			x := sh.Pop()
			if out[j] == x {
				count++
			} else {
				sh.Push(x)
				break
			}
		}
	}
	return count == len(in)
}
