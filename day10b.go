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
		drawCRT(cycle,regX)
		cycle++

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
}

func drawCRT(cycle int, spriteCenter int){
	column := cycle%40

	if column == 0{
		fmt.Println()
	}
	if spriteCenter >= column-1 && spriteCenter <= column+1{
		fmt.Print("#")	
	}else{
		fmt.Print(".")
	}
	
}
