package day01

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func RunPart02() {
	numToName := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	sum := 0
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
		number := transform(string(line), reg, numToName)
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
			return
		}
		sum += number
	}
	println("Magic number: ", sum)
}

func transform(line string, reg *regexp.Regexp, numToName map[int]string) int {
	for stringContainsValueFromMap(line, numToName) {
		for index := 0; index < len(line); index++ {
			for key, name := range numToName {
				str := string(line[index:])
				if strings.HasPrefix(str, name) {
					nameWithoutLastChar := name[:len(name)-1]
					line = strings.Replace(line, nameWithoutLastChar, strconv.Itoa(key), -1)
				}
			}
		}
	}

	cleanString := reg.ReplaceAllString(string(line), "")
	if len(cleanString) >= 2 {
		cleanString = string(cleanString[0]) + string(cleanString[len(cleanString)-1])
	} else {
		cleanString = string(cleanString[0]) + string(cleanString[0])
	}
	num, err := strconv.Atoi(cleanString)

	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
		return 0
	}
	return num
}

func stringContainsValueFromMap(str string, numToName map[int]string) bool {
	for _, name := range numToName {
		if strings.Contains(str, name) {
			return true
		}
	}
	return false
}

func insertAtIndex(source string, insert string, index int) string {
	start := source[:index]
	end := source[index:]
	return start + insert + end
}
