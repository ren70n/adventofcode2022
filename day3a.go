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
				result += tCompare
				res += priority(ch)
				break
			}
		}
	}
	fmt.Println(result, res)
}

func priority(r rune) int{
	if r <= 'Z' {
		return int(r-'A')+27
	}

	return int(r-'a')+1 
}