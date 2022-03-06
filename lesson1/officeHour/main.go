package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)


var dictionary = []string {
	"azure",
	"bagpipes",
	"bandwagon",
	"banjo",
	"bayou",
	"beekeeper",
	"bikini",
	"blitz",
	"blizzard",
	"boggle",
	"bookworm",
	"boxcar",
	"boxful",
	"jaywalk",
	"jazziest",
	"jazzy",
	"jelly",
	"jigsaw",
	"jinx",
	"jiujitsu",
	"jockey",
	"jogging",
	"joking",
	"joyful",
	"juicy",
	"kayak",
	"kazoo",
	"keyhole",
	"khaki",
	"kilobyte",
	"lengths",
	"lucky",
	"luxury",
	"lymph",
	"marquis",
	"matrix",
	"megahertz",
	"microwave",
	"frizzled",
	"fuchsia",
	"funny",
	"gabby",
	"galaxy",
	"quizzes",
	"quorum",
	"razzmatazz",
	"rhubarb",
	"rhythm",
	"staff",
	"strength",
	"strengths",
	"stretch",
	"stronghold",
	"transcript",
	"transgress",
	"transplant",
	"twelfth",
	"twelfths",
	"unknown",
	"unworthy",
	"unzip",
	"uptown",
	"vaporize",
	"vixen",
	"vodka",
	"wave",
	"wavy",
	"waxy",
	"wellspring",
	"wheezy",
	"whiskey",
	"youthful",
	"yummy",
	"zephyr",
	"zigzag",
}

var state = 0

func main(){
	wordToGuess := GiveWordForGame()
	playerWord := make([]byte, len(wordToGuess))
	for i := range playerWord {
		playerWord[i] = '_'
	}
	fmt.Println(wordToGuess)
	fmt.Println("**** Welcome To Hangman Game ****")

	for state != 9 {
		fmt.Println(string(playerWord))
		text := ReadAndValidateUserInput()
		valRes, playerWord := CheckLetterForWord(wordToGuess, text, playerWord)
		if valRes == false {
			fmt.Printf("%s is not contained in the word.\n", text)
			state++
		}
		PrintStateOfGame(state)
		if !strings.Contains(string(playerWord), "_") {
			fmt.Println(strings.ToUpper(wordToGuess))
			fmt.Println("--------- You Won !!!! ------------")
			os.Exit(0)
		}
	}
	if state == 9 {
		fmt.Println(strings.ToUpper(wordToGuess))
		fmt.Println("-------- You Lost !! -----------")
	}
}

func GiveWordForGame() string {
	rand.Seed(time.Now().UnixNano())
	return dictionary[rand.Intn(len(dictionary))]
}

func ReadAndValidateUserInput() string{
	var text string
	result := false
	for result != true {
		text = ReadLetterFromPlayer()
		result = ValidateInputString(text)
		if result != true {
			fmt.Printf("%s is not valid. Please enter valid input.\n", text)
		}
	}
	return text
}

func ReadLetterFromPlayer() string{
	fmt.Printf("\nPlease enter a letter : ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	sz := len(text)
	if sz > 0 {
		text = text[:sz-2]
	}
	return text
}

func ValidateInputString(input string) bool {
	if len(input) > 1 {
		return false
	}
	for _, c := range input {
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return false
		}
	}
	return true
}

func CheckLetterForWord(word string, letter string, userWord []byte) (bool, []byte) {
	index := strings.Index(word, letter)
	if index == -1 {
		return false, userWord
	}
	for i, char := range word {
		if string(char) == letter {
			userWord[i] = byte(char)
		}
	}
	return true, userWord
}

func PrintStateOfGame(s int){
	state := fmt.Sprintf("states/hangman%d",s)
	f, err := os.Open(state)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
