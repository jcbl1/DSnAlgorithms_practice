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

type Node struct {
	Val  int
	Next *Node
}

func (head *Node) insert(x int) {
	if head.Val == 1e9+randInt {
		head.Val = x
		return
	}
	node := new(Node)
	node.Val = head.Val
	node.Next = head.Next
	head.Val = x
	head.Next = node
}

func (head *Node) pop() (x int) {
	x = head.Val
	if head.Next == nil {
		head.Val = 1e9 + randInt
	} else {
		*head = *head.Next
	}
	return
}

func Main() {
	randInt = rand.Intn(60000)

	in, out := getInOut()
	fmt.Println(checkInOut(in, out))
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

func checkInOut(in, out []int) bool {
	if len(in) != len(out) {
		return false
	}
	head := Node{Val: 1e9 + randInt}
	count := 0
	for _, v := range in {
		head.insert(v)
		for j := count; j < len(out); j++ {
			if out[j] == head.Val {
				head.pop()
				count++
			} else {
				break
			}
		}
	}
	if count == len(in) {
		return true
	}
	return false
}
