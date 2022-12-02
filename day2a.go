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
		fmt.Println(x[0],x[1])

		switch x[1] {
		case "X":
			myPoints+=1;
		case "Y":
			myPoints+=2;
		case "Z":
			myPoints+=3;
		}

		if x[0]=="A"{
			switch x[1] {
			case "X":
				myPoints+=3;
			case "Y":
				myPoints+=6;
			case "Z":
				myPoints+=0;
			}
		}
		if x[0]=="B"{
			switch x[1] {
			case "X":
				myPoints+=0;
			case "Y":
				myPoints+=3;
			case "Z":
				myPoints+=6;
			}
		}
		if x[0]=="C"{
			switch x[1] {
			case "X":
				myPoints+=6;
			case "Y":
				myPoints+=0;
			case "Z":
				myPoints+=3;
			}
		}
	}
	fmt.Println(myPoints)
}
