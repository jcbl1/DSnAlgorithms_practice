package main

import (
	"fmt"

	"github.com/jcbl1/DSnAlgorithms_practice/meituan_train_inNout"
)

var ents []string

func main() {
	ents = []string{
		"美团笔试-火车入站出站",
	}
	welcomStr := ""
	for i, v := range ents {
		welcomStr += fmt.Sprintf("%d) %s\n", i+1, v)
	}
	welcomStr += "0) 退出\n"
	welcomStr += "请输入数字："
	fmt.Print(welcomStr)
	selection := 0
	fmt.Scan(&selection)

	switch selection {
	case 0:
		return
	case 1:
		meituan_train_inNout.Main()
	}
}
