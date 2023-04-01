package pkg_regexp

import (
	"fmt"
	"regexp"
)

func Main() {
	re1, err := regexp.Compile("d[aeiou]g")
	checkErr(err)
	fmt.Println(re1.FindString("kldhfdigdslkhfdsdog"))
}
func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
