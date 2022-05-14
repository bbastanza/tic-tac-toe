package main

import (
	"fmt"
	"strings"
	"strconv"
	"errors"
)

func main() {
	var playAgain bool
	currentStarter := "X"

	for !playAgain {

		playAgainAnswer, nextStarter := gameLoop(currentStarter)

		if (!playAgainAnswer) {
			break
		}

		fmt.Println("Let's Go!")

		currentStarter = nextStarter
	}

	fmt.Println("Have a great day!")

}

func togglePlayer(player string) string {

	if player == "X" {
		return "O"
	}

	return "X"
}

func gameLoop(startingPlayer string) (bool, string) {
	var guess string
	var winningPlayer string
	nextStarter := togglePlayer(startingPlayer)
	currentPlayer := startingPlayer
	board := make([][]string, 3)

	for i := 0; i < 3; i++ {
		board[i] = make([]string, 3)
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}

	}


	for true {
		drawBoard(board)

		fmt.Print("What is your guess? ")

		fmt.Scanf("%s", &guess)

		xValue, yValue, err := parseGuess(guess, board)

		if err == nil {
			board[xValue][yValue] = currentPlayer
		} else {
			fmt.Println("")
			fmt.Println("Nope...", err)
			continue
		}

		win, winner := isWon(board)

		if win {
			winningPlayer = winner
			break
		}

		if isTie(board) {
			break
		}

		currentPlayer = togglePlayer(currentPlayer)
	}

	fmt.Println("")
	fmt.Println("")

	fmt.Println("*****")
	if winningPlayer == "" {
		fmt.Println("It's a tie...")

	} else {
		fmt.Println(winningPlayer, "Won!!!")
	}
	fmt.Println("*****")

	drawBoard(board)

	playAgain := getPlayAgain()

	return playAgain, nextStarter
}

func getPlayAgain() bool {
	var playAgain string

	for playAgain != "y" &&
		playAgain != "n" {

		fmt.Println("Play Again? ")
		fmt.Scanf("%s", &playAgain)
	}

	return playAgain == "y"
}

func isTie(arr [][]string) bool {
	if arr[0][0] == " " {
		return false
	}

	for i, v := range arr {
		for j, _ := range v {
			if arr[i][j] == " " {
				return false
			}
		}
	}

	return true
}

func isWon(arr [][]string) (bool, string) {
	var isWon bool
	var winner string

	for _, v := range arr {
		isWon, winner = isHorizontalWinner(v)
		if (isWon) {
			return true, winner
		}
	}

	isWon, winner = isVerticalWinner(arr)

	if isWon {
		return true, winner
	}

	isWon, winner = isDiagnalWinner(arr)

	if isWon {
		return true, winner
	}


	return false, ""
}

func isHorizontalWinner(arr []string) (bool, string) {
	vOne := arr[0]
	vTwo := arr[1]
	vThree := arr[2]

	if vOne == " " {
		return false, ""
	}

	if vOne == vTwo && vTwo == vThree {
		return true, vOne
	}

	return false, ""
}

func isVerticalWinner(arr [][]string) (bool, string) {

	if arr[0][0] != " " &&
		arr[0][0] == arr[1][0] &&
		arr[0][0] == arr[2][0] {

		return true, arr[0][0]
	}

	if arr[0][1] != " " &&
		arr[0][1] == arr[1][1] &&
		arr[0][1] == arr[2][1] {

		return true, arr[0][1]
	}

	if arr[0][2] != " " &&
		arr[0][2] == arr[1][2] &&
		arr[0][2] == arr[2][2] {

		return true, arr[0][2]
	}

	return false, ""
}

func isDiagnalWinner(arr [][]string) (bool, string) {

	if arr[1][1] == " " {
		return false, ""
	}

	if  arr[0][0] == arr[1][1] &&
		arr[1][1] == arr[2][2] {

		return true, arr[1][1]
	}

	if  arr[2][0] == arr[1][1] &&
		arr[1][1] == arr[0][2] {

		return true, arr[1][1]
	}

	return false, ""
}

func drawBoard(arr [][]string) {
	topDisplay := strings.Join(arr[0], "|")
	midDisplay := strings.Join(arr[1], "|")
	lowDisplay := strings.Join(arr[2], "|")

	fmt.Println("")
	fmt.Println(" ",topDisplay)
	fmt.Println("  -----")
	fmt.Println(" ",midDisplay)
	fmt.Println("  -----")
	fmt.Println(" ",lowDisplay)
	fmt.Println("")
}

func convertLetterToIndex(letter string) int {
	switch letter {
	case "a":
		return 0
	case "b":
		return 1
	default:
		return 2
	}
}

func parseGuess(guess string, board [][]string) (int, int, error) {
	validLetters := []string{"a", "b", "c"}
	validNumbers := []string{"1", "2", "3"}

	if len(guess) != 2{
		return 0, 0, errors.New("Bad guess")
	}

	chars := strings.Split(guess, "")

	invalidLetter := !contains(validLetters, chars[0])
	invalidNumber := !contains(validNumbers, chars[1])

	if invalidLetter || invalidNumber {
		return 0, 0, errors.New("Bad guess")
	}

	if !contains(validNumbers, chars[1]){
		return 0, 0, errors.New("Bad guess")
	}

	convertLetterToIndex(chars[0])

	xValue := convertLetterToIndex(chars[0])
	yValue, _ := strconv.Atoi(chars[1])
	yValue = yValue - 1


	if board[xValue][yValue] != " " {
		return 0, 0, errors.New("Already taken")
	}

	return xValue, yValue, nil
}


func contains(arr []string, value string) bool {
	for _, v := range arr {
		if v == value{
			return true
		}
	}
	return false
}

