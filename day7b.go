package main

import (
	"fmt"
	"strings"
	"strconv"
	"adventofcode2022/helper"
)

type file struct{
	name string
	size int
}

type dir struct{
	name string
	files []file
	dirs []dir
	size int
	dirUp *dir
}

func main(){
	data := helper.ReadFileString("day7.data")
	root := dir{name: "/"}

	currentDir := &root
	appendFiles := false
	
	for _, line := range data{
		prompt := strings.Split(line," ")

		switch prompt[0]{
		case "$":
			switch prompt[1]{
			case "cd":
				switch prompt[2]{
				case "/":
					currentDir = &root
				case "..":
					if currentDir.dirUp != nil{
						currentDir = currentDir.dirUp
					}
				default:
					dirIndex := checkIfDirExists(currentDir,prompt[2])
			
					if dirIndex < 0{
						fmt.Println("directory",prompt[2],"does not exists")
						
					}else{
						currentDir = &currentDir.dirs[dirIndex]
					}
				}
				appendFiles = false
			case "ls":
				appendFiles = true
			}
		default:
			if appendFiles{

				if prompt[0] == "dir"{
					newDir := newDir(prompt[1],currentDir)
					currentDir.dirs = append(currentDir.dirs,newDir)
				}else{
					fileSize,_ := strconv.Atoi(prompt[0])
					newFile:= file{name:prompt[1],size:fileSize}

					currentDir.files = append(currentDir.files,newFile)
				}
			}
		}
	}
	sizeTab := make([]int,0)
	fillDirSizes(&root,&sizeTab)

	fmt.Println(findProperDirSize(sizeTab,root.size))
}

func findProperDirSize(sizeTab []int,rootSize int) int{
	min := rootSize
	minIndex := len(sizeTab)-1

	discSize := 70000000
	freeSpace := 30000000

	for i,v := range sizeTab{
		temp := discSize - rootSize + v
		if temp > freeSpace{
			if temp < min{
				min = temp
				minIndex = i
			}
		}
	}

	return sizeTab[minIndex]
}

func fillDirSizes(_dir *dir,_sizes *[]int){
	dirs := _dir.dirs

	for _,d :=range dirs{
		if d.size==0{
			fillDirSizes(&d,_sizes)
		}
		_dir.size+=d.size
	}

	for _,f :=range _dir.files{
		_dir.size += f.size
	}

	if len(dirs)==0{
		// sum the files
		fileSize := 0
		for	_,_file := range _dir.files{
			fileSize += _file.size
		}
		_dir.size = fileSize
	}
	*_sizes = append(*_sizes,_dir.size)
}

func newDir(name string, dirUp *dir)dir{
	newDir := dir{
		name: name,
		dirUp: dirUp}
	return newDir
}

func checkIfDirExists(currentDir *dir, dirName string)int{
	for i,d := range currentDir.dirs{
		if d.name == dirName{
			return i
		}
	}
	return -1
}