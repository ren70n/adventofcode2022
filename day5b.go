package main

import (
	"fmt"
	"strings"
	"errors"
	"strconv"
	"adventofcode2022/helper"
)

func main(){
	data := helper.ReadFileString("day5.data")

	stacks, endOfInput ,stacksCnt := initStacks(data)
	
	for _,row := range data[endOfInput:]{ //endOfInput+5
		line := strings.Split(row," ")
		howmuch, _ := strconv.Atoi(line[1])
		from, _ := strconv.Atoi(line[3])
		to, _ := strconv.Atoi(line[5])

		from-=1
		to-=1

		crates,rest := pull(howmuch,stacks[from])
		stacks[from] = rest

		stackTo := append(crates, stacks[to]...)
		stacks[to] = stackTo
	}
	res := ""
	for i:=0;i<=stacksCnt;i++{
		res+=stacks[i][0]
	}
	fmt.Println(res)
}
func initStacks(data []string)(map[int][]string,int,int){
	endOfInput := 0
	stacksCnt := 0

	stacks := make(map[int][]string)

	for i,row := range data {

		crate:=""
		for x := 0; x < len(row)-4 ; x += 4{
			crate,e := crateTrim(row[x:x+4])

			if e == nil{
				stackNum := x/4
				if stackNum > stacksCnt{stacksCnt = stackNum}
				stacks[stackNum] = append(stacks[stackNum],crate)
			}else{
				if crate == "eol"{
					break 
				}
			}
		}
		if crate == "eol"{
			endOfInput = i
			break
		}

		if len(row)>3{
			crate,e := crateTrim((row[len(row)-4:]))
			if e == nil{
				stackNum := len(row)/4
				if stackNum > stacksCnt{stacksCnt = stackNum}
				stacks[stackNum] = append(stacks[stackNum],crate)
			}
		}else{
			endOfInput = i
			break
		}
	}
	return stacks,endOfInput+1,stacksCnt
}

func crateTrim(crate string)(string,error){
	crate = strings.TrimSpace(crate)
	if len(crate)==0{
		return "",errors.New("empty crate")
	}
	_,e := strconv.Atoi(crate)
	if e==nil{
		return "eol",errors.New("eol")
	}
	crate = strings.ReplaceAll(crate,"[","")
	crate = strings.ReplaceAll(crate,"]","")

	return crate,nil
}

func pull(howmuch int,from []string)(taken, rest []string){

	crates := make ([]string,howmuch)
	copy(crates,from[:howmuch])
	stackFrom := make([]string,len(from)-howmuch)
	copy(stackFrom,from[howmuch:])

	return crates,stackFrom
}
