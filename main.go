package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	numberSet    = "0123456789"
	lowerCharSet = "abcdefghijqlmnopqrstuvwxyz"
	upperCharSet = "ABCDEFGHIJQLMNOPQRSTUVWXYZ"
	symbolSet    = "!@#$%&*"
	allCharSet   = numberSet + lowerCharSet + upperCharSet + symbolSet
)

var (
	password                                    strings.Builder
	result                                      string
	length, digit, lowerChar, upperChar, symbol int
)

func main() {

	fmt.Printf("\n Wellcome To The CLI Password Generator Application 😎   \n")

	command := flag.String("command", "0", "Password Generation")
	flag.Parse()

	for {
		runCommands(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Please enter another command: ")
		fmt.Println("commands to use => 1) create-password    2) exit")
		scanner.Scan()
		*command = scanner.Text()
	}

}

func generateRandomPassword() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("enter the length of password(enter number) : ")
	scanner.Scan()
	length, _ = strconv.Atoi(scanner.Text())

	if length != 0 {
		generateInput()

		total := digit + lowerChar + upperChar + symbol

		if total <= length {

			// set number
			createRandomChar(digit, numberSet)

			// set lower character
			createRandomChar(lowerChar, lowerCharSet)

			// set upper character
			createRandomChar(upperChar, upperCharSet)

			// set symbols
			createRandomChar(symbol, symbolSet)

			// remaining length
			remainingLength := length - (digit + lowerChar + upperChar + symbol)
			createRandomChar(remainingLength, allCharSet)

			// shuffle of character password
			result = shuffleString()
			fmt.Printf("🔑 Your Password is :  %v \n\n", result)

		} else {
			fmt.Println("The sum of the entered numbers must be equal length or lower then length 😒 ")
		}

	} else {
		fmt.Println("Please enter a valid number ❌ ")
	}
}

func shuffleString() string {
	passwordRune := []rune(password.String())
	rand.Shuffle(len(passwordRune), func(i, j int) {
		passwordRune[i], passwordRune[j] = passwordRune[j], passwordRune[i]
	})
	return string(passwordRune)
}

func createRandomChar(char int, charTypeSets string) strings.Builder {
	for i := 0; i < char; i++ {
		random := rand.Intn(len(charTypeSets))
		password.WriteString(string(charTypeSets[random]))
	}
	return password
}

func runCommands(command string) {

	switch command {
	case "create-password":
		generateRandomPassword()

	case "exit":
		os.Exit(0)

	default:
		fmt.Println("command is not valid ❌ ")
	}
}

func generateInput() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digit :  ")
	scanner.Scan()
	digit, _ = strconv.Atoi(scanner.Text())

	fmt.Println("Lower Characters : ")
	scanner.Scan()
	lowerChar, _ = strconv.Atoi(scanner.Text())

	fmt.Println("Upper Characters : ")
	scanner.Scan()
	upperChar, _ = strconv.Atoi(scanner.Text())

	fmt.Println("Symbol Characters : ")
	scanner.Scan()
	symbol, _ = strconv.Atoi(scanner.Text())
}
