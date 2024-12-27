package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	answers := []string{
		"It is certain.",
		"It is decidedly so.",
		"Without a doubt.",
		"Yes – definitely.",
		"You may rely on it.",
		"As I see it, yes.",
		"Most likely.",
		"Outlook good.",
		"Yes.",
		"Signs point to yes.",
		"Reply hazy, try again.",
		"Ask again later.",
		"Better not tell you now.",
		"Cannot predict now.",
		"Concentrate and ask again.",
		"Don't count on it.",
		"My reply is no.",
		"My sources say no.",
		"Outlook not so good.",
		"Very doubtful.",
	}

	fmt.Println("Magic 8-Ball Simulator")
	fmt.Println("========================")
	fmt.Println("Ask any yes-or-no question or type 'exit' to quit.\n")

	for {
		fmt.Print("What is your question? ")
		var question string
		fmt.Scanln(&question)

		if question == "exit" {
			fmt.Println("Goodbye! May your future be bright!")
			break
		}

		// Generate a random response
		randomIndex := rand.Intn(len(answers))
		fmt.Printf("Magic 8-Ball says: %s\n\n", answers[randomIndex])
	}
}
