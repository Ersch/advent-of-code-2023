package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func RunPart02() {
	var games []Game

	f, err := os.Open("./day02/day02.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		newGame := lineToGamePart02(string(line))
		games = append(games, newGame)
	}

	sumOfCubePerGame := 0

	for _, game := range games {
		for _, round := range game.rounds {
			sumOfCubePerGame += round.Blue * round.Red * round.Green
		}
	}
	fmt.Println("Sum of power of cubes: ", sumOfCubePerGame)
}

func lineToGamePart02(line string) Game {
	round := strings.Split(line, ":")
	re := regexp.MustCompile(`(\d+)`)
	match := re.FindStringSubmatch(round[0])
	gameID, _ := strconv.Atoi(match[1])
	line = strings.Replace(line, round[0], "", -1)
	line = strings.Replace(line, ":", "", -1)
	game := Game{}
	game.id = gameID
	rounds := strings.Split(line, ";")
	newRound := Round{Red: 0, Blue: 0, Green: 0}
	for _, round := range rounds {
		colors := strings.Split(round, ",")
		for _, color := range colors {
			quantity, color := parseQuantityAndColor(color)
			switch color {
			case "red":
				if quantity > newRound.Red {
					newRound.Red = quantity
				}
			case "blue":
				if quantity > newRound.Blue {
					newRound.Blue = quantity
				}
			case "green":
				if quantity > newRound.Green {
					newRound.Green = quantity
				}
			}
		}
	}
	game.rounds = append(game.rounds, newRound)
	return game
}
