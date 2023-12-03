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

/*
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
*/

type Round struct {
	Red   int
	Blue  int
	Green int
}

type Game struct {
	id     int
	rounds []Round
}

func RunPart01() {
	var games []Game

	f, err := os.Open("./day02/day02.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	if err != nil {
		log.Fatalf("Error compiling regex: %v", err)
		return
	}

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		newGame := lineToGame(string(line))
		games = append(games, newGame)
	}

	nbBlue := 14
	nbRed := 12
	nbGreen := 13
	sumOfPossibleGames := 0

	for _, game := range games {
		possible := false
		for _, round := range game.rounds {
			if round.Blue > nbBlue || round.Red > nbRed || round.Green > nbGreen {
				possible = false
				break
			} else {
				possible = true
			}
		}

		if possible {
			sumOfPossibleGames += game.id
		}
	}

	fmt.Println("Sum of possible games: ", sumOfPossibleGames)
}

func lineToGame(line string) Game {
	round := strings.Split(line, ":")
	re := regexp.MustCompile(`(\d+)`)
	match := re.FindStringSubmatch(round[0])
	gameID, _ := strconv.Atoi(match[1])
	line = strings.Replace(line, round[0], "", -1)
	line = strings.Replace(line, ":", "", -1)
	game := Game{}
	game.id = gameID
	rounds := strings.Split(line, ";")
	for _, round := range rounds {
		newRound := Round{Red: 0, Blue: 0, Green: 0}
		colors := strings.Split(round, ",")
		for _, color := range colors {
			quantity, color := parseQuantityAndColor(color)
			switch color {
			case "red":
				newRound.Red = quantity
			case "blue":
				newRound.Blue = quantity
			case "green":
				newRound.Green = quantity
			}
		}
		game.rounds = append(game.rounds, newRound)
	}
	return game
}

func parseQuantityAndColor(line string) (int, string) {
	re := regexp.MustCompile(`(\d+) (\w+)`)
	match := re.FindStringSubmatch(line)
	quantity, _ := strconv.Atoi(match[1])
	color := match[2]
	return quantity, color
}
