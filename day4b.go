package main

import (
	"fmt"
	"strings"
	"strconv"
	"adventofcode2022/helper"
)

func main(){
	data := helper.ReadFileString("day4.data")

	res := 0

	for _, row := range data{
		section := strings.Split(row,",")
		secAstr := strings.Split(section[0],"-")
		secBstr := strings.Split(section[1],"-")
		
		a,_ := strconv.Atoi(secAstr[0])
		b,_ := strconv.Atoi(secAstr[1])

		x,_ := strconv.Atoi(secBstr[0])
		y,_ := strconv.Atoi(secBstr[1])

		secA,secB := []int{a,b},[]int{x,y}

		if secA[0]>secB[0]{secA,secB = secB,secA}

		if secB[0]<=secA[1]{
			res++
		}
	}
	fmt.Println(res)
}