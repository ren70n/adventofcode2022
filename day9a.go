package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"adventofcode2022/helper"
)

func main(){
	data := helper.ReadFileString("day9.data")

	head := []int{0,0}
	headPath := make([][]int,0)
	headPath = append(headPath,head)
	tail := []int{0,0}

	worldMap := make(map[string]int)
	worldMap[fmt.Sprintf("%d:%d",tail[0],tail[1])]++

	for _, line :=range data {

		row := strings.Split(line," ")
		dir,strStep := row[0],row[1]

		step,_ := strconv.Atoi(strStep)

		for i:=0;i<step;i++{
			dist := countDistance(head,tail)
			if dist >=1.5{
				tail = headPath[len(headPath)-2]
				worldMap[fmt.Sprintf("%d:%d",tail[0],tail[1])]++

			}
			newHead := moveHead(dir,head)

			headPath = append(headPath,newHead)
			head = headPath[len(headPath)-1]

		}
	}
	fmt.Println(len(worldMap))
}

func moveHead(dir string,pos []int)[]int{
	ret := []int{pos[0],pos[1]}

	switch dir{
	case "R":
		ret[0]++
	case "L":
		ret[0]--
	case "U":
		ret[1]++
	case "D":
		ret[1]--		
	}
	return ret
}

func countDistance(A,B []int)float64{
	_a , _b := float64(A[0]-B[0]),  float64(A[1]-B[1])
	dist := math.Sqrt(math.Pow(_a,2)+math.Pow(_b,2))

	return dist
}

// func printMap(_map map[string]int){

// 	fmt.Println(_map)

// 	for y:=0;y<6;y++{
// 		for x:=0;x<6;x++{
// 			// fmt.Println(_map[fmt.Sprintf("%d:%d",x,y)])
// 			if _map[fmt.Sprintf("%d:%d",x,y)]>0{
// 				fmt.Print("#")
// 			}else{
// 				fmt.Print("_")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }
