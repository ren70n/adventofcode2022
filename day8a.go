package main

import (
	"fmt"
	"adventofcode2022/helper"
)

func main(){
	data := helper.ReadFileString("day8.data")
	
	field := make([][]int,0)
	
	for _,line := range data{
		row := make([]int,0)
		for _, char := range line{
			row = append(row,int(char-'0'))
		}
		field = append(field,row)
	}

	boolField := make([][]bool,0)
	
	topBottomRow := make([]bool,len(field[0]))
	
	for i,_ := range topBottomRow{
		topBottomRow[i] = true
	}
	boolField = append(boolField,topBottomRow)	
	for y:=1 ; y<len(field)-1 ; y++{
		row := field[y]
		boolField = append(boolField,checkRow(row))
		
	}
	boolField = append(boolField,topBottomRow)

	// columns
	
	for x,_ := range field[0]{
		col := make([]int,len(field))
		for y:=0;y<len(field);y++{
			col[y]=field[y][x]
		}
		boolCol := checkRow(col)
		
		boolField = boolColMerge(boolField,boolCol,x)
	}
	
	for _,r:=range(boolField){
		fmt.Println(r)
	}
		
	fmt.Println(sumUp(boolField))
	
}
func sumUp(input [][]bool)int{
	res := 0
	
	for _,line:=range input{
		for _,val := range line{
			if val{res++}
		}
	}
	
	return res
}

func checkRow(row []int)[]bool{
	max := row[0]
	
	tall := make([]bool,len(row))
	tall[0]=true
	tall[len(row)-1]=true
	
	for i,x := range row{
		if x>max{
			tall[i] = true
			max = x
		}
	}
	max = row[len(row)-1]
	
	for i := len(row)-1;i>0;i--{
		x := row[i]
		if x>max{
			tall[i] = true
			max = x
		}
	}
	return tall

}
func boolColMerge(bTable [][]bool,col []bool,cNum int)[][]bool{
	for y,val := range col{
		bTable[y][cNum] = bTable[y][cNum] || val
	}
	
	return bTable
}
