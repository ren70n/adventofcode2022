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

	maxSize := 100000
	fmt.Println(fillDirSizes(&root,maxSize))
}

func fillDirSizes(_dir *dir,maxSize int)int{
	dirs := _dir.dirs

	ret := 0

	for _,d :=range dirs{
		if d.size==0{
			ret += fillDirSizes(&d, maxSize)
		}
		_dir.size+=d.size
	}

	for _,f :=range _dir.files{
		_dir.size += f.size
	}

	// this is final directory
	if len(dirs)==0{
		// sum the files
		fileSize := 0
		for	_,_file := range _dir.files{
			fileSize += _file.size
		}
		_dir.size = fileSize
	}

	if _dir.size <= maxSize{
		return _dir.size + ret
	}

	return ret
}

func printDir(d dir){
	fmt.Println(">",d.name)
	for _,_dir := range d.dirs{
		fmt.Println("dir",_dir.name)
	}
	for _,_file := range d.files{
		fmt.Println(_file.name,_file.size)
	}
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