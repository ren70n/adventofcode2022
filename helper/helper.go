package helper

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFileString(name string)[]string{

	file, err := os.Open(name)

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var data []string

	
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}

func ReadFileInt(name string)[]int{

	file, err := os.Open(name)

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var data []int

	
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		data = append(data, value)
	}

	return data
}