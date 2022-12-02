package main

import (
	"fmt"
	"strconv"
	"adventofcode2022/helper"
)

type elf struct{
	foodCallories []int
	max int
}

func main(){
	data := helper.ReadFileString("day1.data")
	sElf := elf{foodCallories: make([]int,0)}
	elfArr := make([]elf,0)

	// max := 0
	for _,line:=range data{
		val,err := strconv.Atoi(line)

		if err==nil{
			sElf.foodCallories = append(sElf.foodCallories, val)
			sElf.max += val
		}else{
			elfArr = append(elfArr,sElf)
			sElf = elf{foodCallories: make([]int,0)}
		}
	}
	if len(sElf.foodCallories)>0{
		elfArr = append(elfArr,sElf)
	}

	max := 0
	for _,e := range elfArr{
		if e.max > max{
			max = e.max
		}
	}
	fmt.Println(max)
}
