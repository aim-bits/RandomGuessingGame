package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

)
 

func main()  {
	fmt.Println("Guess a number between 0 to 6")

	COMEBACK:
	rand.Seed(time.Now().UnixNano())
	var	diceNumber int = rand.Intn(6) + 1

	
	guessReader := bufio.NewReader(os.Stdin)
	guessInput, _ := guessReader.ReadString('\n')

	castedInput, err := strconv.ParseInt(strings.TrimSpace(guessInput), 8, 8)
	if err != nil {
		panic(err)
	}

	fmt.Println("you guessed: ", castedInput)
	fmt.Println("Random Number is: ", diceNumber)
	var resReceipt []int

	resReceipt = append(resReceipt, int(castedInput), diceNumber)
	fmt.Println(resReceipt)
	
	if castedInput == int64(diceNumber) {
		fmt.Println("Correct!!")
	} else if castedInput != int64(diceNumber) {
		fmt.Println("Wrong!! Try again")
		goto COMEBACK
	} else {
		fmt.Println("You selected out of range")
	}



}

