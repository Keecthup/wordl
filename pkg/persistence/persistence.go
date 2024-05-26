package persistence

import (
	"bufio"
	"math/rand"
	"mygame/pkg/constants"
	"os"
)

func IsWordValid(word string) bool {
	file, err := os.Open("dictionary.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == word {
			return true
		}
	}
	return false
}

func GetRandomWord() string {
	file, err := os.Open("dictionary.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	randomIndex := rand.Intn(constants.WordQuantity)

	scanner := bufio.NewScanner(file)
	for i := 0; i <= randomIndex; i++ {
		scanner.Scan()
		if i == randomIndex {
			return scanner.Text()
		}
	}
	return ""
}
