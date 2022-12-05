package main

import (
	"fmt"
	"strings"
	"adventofcode2022/helper"
)

func main(){
	data := helper.ReadFileString("day3.data")
	group := 0
	oldDict := ""

	res := 0

	for _, val := range data{

		dict := makeDict(val)
		if group == 0 {oldDict = dict}else{
			oldDict = findSingleChar(dict,oldDict)
			
			if group == 2{
				res += priority(rune(oldDict[0]))
			}
		}

		group ++
		group%=3
	}
	fmt.Println(res)
}

func priority(r rune) int{
	if r <= 'Z' {
		return int(r-'A')+27
	}
	return int(r-'a')+1 
}

func makeDict(s string)string{
	mapOfChars := make(map[string]int)
	dict := ""
	for _,b := range s{
		ch := string(rune(b))
		mapOfChars[ch]++
		if mapOfChars[ch] == 1{ dict+=ch }
	}
	return dict
}
func findSingleChar(input,dict string)string{
	res := ""
	for _,ch := range dict{
		letter := string(rune(ch))
		if strings.Contains(input,letter){res+=letter}

	}
	return res
}