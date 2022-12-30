package main

import (
	"fmt"
	"math"
	"adventofcode2022/helper"
)

func main(){
	data := helper.ReadFileString("day25.data")

	var ret int64
	for _,row := range data{
		val := SNAFUToDecimal(row)
		ret += val
	}
	fmt.Println(DecimalToSNAFU(ret))
}

func SNAFUToDecimal(in string)int64{
	var ret int64
	pow := 0

	for i:=len(in)-1;i>=0;i--{
		var value int64
		multiplier := int64(math.Pow(float64(5),float64(pow)))
		pow++

		char := in[i:i+1]

		if char == "-"{
			value = -1
		}else if char == "="{
			value = -2
		}else{
			value = int64(char[0]-'0')
		}
		ret += value*multiplier

	}
	return ret
}

func DecimalToSNAFU(in int64)string{

	if in == 0 {return ""}

	decRem := in%5
	SNAFUDigits := []string{"0","1","2","=","-"}

	newDec := (in+2)/5

	return DecimalToSNAFU(newDec)+SNAFUDigits[decRem]
}
