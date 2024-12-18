package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func makeGuess(correctNumber int, scanner *bufio.Scanner, chances int) {
	for i := 0; i < chances; i++ {
		fmt.Printf("You have %d chances left\n", chances-i)
		fmt.Print("Enter your guess: ")

		if !scanner.Scan() {
			fmt.Println("Failed to read input")
			return
		}

		guess, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input, please enter a number")
			continue
		}

		if guess < correctNumber {
			fmt.Println("Your guess is too low")
		} else if guess > correctNumber {
			fmt.Println("Your guess is too high")
		} else {
			fmt.Println("Congratulations! You guessed the correct number")
			return
		}
	}
	fmt.Println("Sorry, you've run out of chances. The correct number was", correctNumber)
}

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with the current time
	correctNumber := rand.Intn(100) + 1
	chances := 10

	scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Choose a number between 1 and 100")
    fmt.Println("1. Easy (10 chances)")
    fmt.Println("2. Medium (5 chances)")
    fmt.Println("3. Hard (3 chances)")
    fmt.Println("4. Exit")
    fmt.Print("Enter your choice: ")

  if !scanner.Scan() {
    fmt.Println("Failed to read input")
    return
  }

  choice, err := strconv.Atoi(scanner.Text())
  if err != nil {
    fmt.Println("Invalid input, please enter a number")
    return
  }

  switch choice {
    case 1:
      chances = 10
      makeGuess(correctNumber, scanner, chances)
    case 2:
      chances = 5
      makeGuess(correctNumber, scanner, chances)
    case 3:
      chances = 3
      makeGuess(correctNumber, scanner, chances)
    case 4:
      fmt.Println("Exiting the game")
    default:
      fmt.Println("Invalid choice, exiting the game")
  }
}




