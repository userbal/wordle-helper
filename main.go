package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type letterWithLocation struct {
	letter   string
	location int
}

func parseColoredLetters(letters []string) []letterWithLocation {
	x := []letterWithLocation{}
	for i := 0; i < len(letters)-1; i++ {
		location, err := strconv.Atoi(string(letters[i+1]))
		if err != nil {
			panic("Something went wrong when reading your letters with a location. Make sure it follows this format a 0 b 3")
		}
		x = append(x, letterWithLocation{string(letters[i]), location})
		i++
	}
	return x
}

func getLetters() (greenLetters, yellowLetters []letterWithLocation, grayLetters []string) {
	fmt.Println("Type the green letters with their position number. Example: a1b3c4")
	var rawGreenLetters string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		rawGreenLetters = scanner.Text()
	}

	fmt.Println("Type the yellow letters with their position number. Example: a5b2cr2")
	var rawYellowLetters string
	if scanner.Scan() {
		rawYellowLetters = scanner.Text()
	}

	fmt.Println("Type the grey letters. Example: abc")
	var rawGrayLetters string
	if scanner.Scan() {
		rawGrayLetters = scanner.Text()
	}

	greenLetters = parseColoredLetters(strings.Split(rawGreenLetters, ""))
	yellowLetters = parseColoredLetters(strings.Split(rawYellowLetters, ""))
	grayLetters = strings.Split(rawGrayLetters, "")

	return greenLetters, yellowLetters, grayLetters
}

func main() {

	wordsString, err := os.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(wordsString), "\n")

	greenLetters, yellowLetters, grayLetters := getLetters()

	newWords := []string{}
	for _, w := range words {
		qualifies := true
		for _, l := range grayLetters {
			if strings.Contains(w, l) {
				qualifies = false
				break
			}
		}
		if qualifies && w != "" {
			newWords = append(newWords, w)
		}
	}
	words = newWords

	newWords = []string{}
	for _, w := range words {
		qualifies := true
		for _, l := range greenLetters {
			if string(w[l.location]) != l.letter {
				qualifies = false
				break
			}
		}
		if qualifies {
			newWords = append(newWords, w)
		}

	}
	words = newWords

	newWords = []string{}
	for _, w := range words {
		qualifies := true
		for _, l := range yellowLetters {
			if !strings.Contains(w, l.letter) {
				qualifies = false
				break
			}
			if string(w[l.location]) == l.letter {
				qualifies = false
				break
			}
		}
		if qualifies {
			newWords = append(newWords, w)
		}
	}

	fmt.Println(newWords)
}
