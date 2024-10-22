package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func generateRandomNumber() int {
	randomNumber := rand.Intn(100)

	return randomNumber
}

func readUserInput() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)

	if err != nil {
		stdin.ReadString('\n')
	}

	return n, err
}

func checkUserInputSame(userInput, random int) bool {
	return userInput == random
}

func main() {
	randomNumber := generateRandomNumber()

	fmt.Println(randomNumber)

	for {
		fmt.Println("숫자를 입력하세요.")
		userNumber, err := readUserInput()

		if (err != nil) {
			fmt.Println("올바른 숫자를 입력해주세요.")
			continue
		}

		fmt.Println("입력한 숫자:", userNumber)
		result := checkUserInputSame(userNumber, randomNumber)

		if result {
			fmt.Println("맞았습니다.")
			break
		}
	}
}