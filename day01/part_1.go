package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func RunPart01() {
	sum := 0
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(wd)
	f, err := os.Open("./day01/day01.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	defer f.Close()

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
		cleanString := toTwoDigits(reg, line)
		number, err := strconv.Atoi(cleanString)
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
			return
		}
		sum += number
	}
	fmt.Println("Magic number: ", sum)
}

func toTwoDigits(reg *regexp.Regexp, line []byte) string {
	cleanString := reg.ReplaceAllString(string(line), "")
	if len(cleanString) >= 2 {
		cleanString = string(cleanString[0]) + string(cleanString[len(cleanString)-1])
	} else {
		cleanString = string(cleanString[0]) + string(cleanString[0])
	}
	return cleanString
}
