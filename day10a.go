package main

import (
	"fmt"
	"strings"
	"strconv"
	"adventofcode2022/helper"
)

func main(){
	data := helper.ReadFileString("day10.data")

	cycle := 0
	regX := 1

	inExec := 0
	V := 0

	instruction := []string{"",""}
	currentDataLine := 0

	signalStrength := 0

	for true{
		cycle++
		// fmt.Println(cycle, regX)
		if cycle==20 || (cycle-20)%40==0{
			signalStrength += regX*cycle
		}

		if inExec <= 0{
			if currentDataLine>=len(data){
				break
			}

			instruction = strings.Split(data[currentDataLine]," ")
			currentDataLine++

			switch instruction[0]{
			case "noop":
				inExec = 0
			case "addx":
				V,_ = strconv.Atoi(instruction[1])

				inExec = 1
			}

		}else{
			inExec--

			if inExec <= 0{
				switch instruction[0]{
				case "addx":
					regX+=V
				}
			}
		}
	}
	fmt.Println(signalStrength)
}
