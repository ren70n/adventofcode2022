package main

import (
	"fmt"
	"adventofcode2022/helper"
)

func main(){
	data := helper.ReadFileString("day6.data")
	
	for _, row := range data{

		for i:=0;i<len(row)-4;i++{
			if checkWindow(row[i:i+4]){
				fmt.Println(i+4)
				break
			}
		}
	}
}

func checkWindow(window string)bool{
	mapOfChars := make(map[rune]int)
	
	for _,b := range window{
		mapOfChars[b]++
		if mapOfChars[b] > 1{return false}
	}
	
	return true
}