package sololearn_its_a_sigh

import "fmt"


func Main(){
	signs:=make([]string,4)
	for i:=range signs{
		fmt.Scanln(&signs[i])
		if isPal(signs[i]){
			fmt.Println("Open")
			return
		}
	}
}

func isPal(s string)bool{
	halfway:=len(s)/2
	for i:=0;i<=halfway;i++{
		j:=len(s)-1-i
		if s[i]!=s[j]{
			return false
		}
	}
	return true
}
