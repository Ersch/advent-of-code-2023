package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	sum := 0
	f, err := os.Open("day01.txt")
	defer f.Close()

	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	reader := bufio.NewReader(f)
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatalf("Error compiling regex: %v", err)
		return
	}

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		cleanString := reg.ReplaceAllString(string(line), "")
		if len(cleanString) >= 2 {
			cleanString = string(cleanString[0]) + string(cleanString[len(cleanString)-1])
		} else {
			cleanString = string(cleanString[0]) + string(cleanString[0])
		}
		number, err := strconv.Atoi(cleanString)
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
			return
		}
		sum += number
	}
	fmt.Println("Magic number: ", sum)
}
