package main

import (
	"fmt"
	"adventofcode2022/helper"
)

func main(){
	data := helper.ReadFileString("day6.data")
	
	markerSize := 14
	
	for _, row := range data{

		for i:=0;i<len(row)-markerSize;i++{
			if checkWindow(row[i:i+markerSize]){
				fmt.Println(i+markerSize)
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