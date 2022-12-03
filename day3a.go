package main

import (
	"fmt"
	"strings"
	"adventofcode2022/helper"
)

func main(){
	data := helper.ReadFileString("day3.data")
	result := ""
	res := 0

	for _, val := range data{
		comp1, comp2 := val[:len(val)/2], val[len(val)/2:]

		for _,ch := range comp1{
			tCompare := string(rune(ch))
			if strings.Contains(comp2, tCompare){
				result+=tCompare
				if ch<='Z'{
					res += int(ch-'A')+27
				}else{
					res += int(ch-'a')+1
				}
				break
			}
		}
	}
	fmt.Println(result, res)
}