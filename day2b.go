package main

import (
	"fmt"
	"strings"
	// "strconv"
	"adventofcode2022/helper"
)

func main(){
	data := helper.ReadFileString("day2.data")
	myPoints := 0
	for _,line := range data{
		x:= strings.Split(line," ")

		if x[0]=="A"{
			switch x[1] {
			case "X":
				myPoints+=(3+0);
			case "Y":
				myPoints+=(1+3);
			case "Z":
				myPoints+=(2+6);
			}
		}

		if x[0]=="B"{
			switch x[1] {
			case "X":
				myPoints+=(1+0);
			case "Y":
				myPoints+=(2+3);
			case "Z":
				myPoints+=(3+6);
			}
		}

		if x[0]=="C"{
			switch x[1] {
			case "X":
				myPoints+=(2+0);
			case "Y":
				myPoints+=(3+3);
			case "Z":
				myPoints+=(1+6);
			}
		}
	}
	fmt.Println(myPoints)
}
