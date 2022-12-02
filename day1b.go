package main

import (
	"fmt"
	"strconv"
	"sort"
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

// now we have to sort it to get 3 most 

	sort.Sort(elfList(elfArr))
	fmt.Println(elfArr[0].max+elfArr[1].max+elfArr[2].max)
}

type elfList []elf

func (e elfList) Len() int {
	return len(e)
}

func (e elfList) Less(i, j int) bool {
    return e[i].max > e[j].max
}

func (e elfList) Swap(i, j int) {
    e[i], e[j] = e[j], e[i]
}
