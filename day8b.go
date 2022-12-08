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

	boolField := make([][]int,0)
	
	boolField = append(boolField,field[0])
	
	for y:=1 ; y<len(field)-1 ; y++{
		row := field[y]
		boolField = append(boolField,checkRow(row))
		
	}
	boolField = append(boolField,field[len(field)-1])

	for x,_ := range field[0]{
		col := make([]int,len(field))
		for y:=0;y<len(field);y++{
			col[y]=field[y][x]
		}
		boolCol := checkRow(col)
		
		boolField = boolColMerge(boolField,boolCol,x)
	}
	
	res := make([][]int,len(field))
	
	for y,row := range boolField{
		res[y] = make([]int,len(field[0]))
		
		for x,_ := range row{
			if boolField[y][x] > 0{
				countDist(boolField,x,y)
			}
		}
	}
	
	for _,r:=range(boolField){
		fmt.Println(r)
	}
}

func countDist(field [][]int, x,y int)int{
	res := 1
	baseVal := field[y][x]
	// left
	cntl := 0
	for i:=x-1;i>=0;i--{
		cntl++
		if field[y][i]>=baseVal{
			break
		}
	}
	// right
	cntr := 0
	for i:=x+1;i<len(field[y]);i++{
		cntr++
		if field[y][i]>=baseVal{
			break
		}
	}
	fmt.Println(baseVal,cntr,cntl)
	// up
	// down
	
	return res
}

func checkRow(row []int)[]int{
	max := row[0]
	
	tall := make([]int,len(row))
	tall[0] = row[0]
	tall[len(row)-1]=row[len(row)-1]
	
	for i,x := range row{
		if x>max{
			tall[i] = x
			max = x
		}
	}
	max = row[len(row)-1]
	
	for i := len(row)-1;i>0;i--{
		x := row[i]
		if x>max{
			tall[i] = x
			max = x
		}
	}
	return tall

}
func boolColMerge(bTable [][]int,col []int,cNum int)[][]int{
	for y,val := range col{
		if bTable[y][cNum] == 0 { bTable[y][cNum] = val }
	}
	
	return bTable
}
