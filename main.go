package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Can not load .env: %v", err)
		os.Exit(1)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY environment variable is not set")
		fmt.Println("Please set your OpenAI API key: export OPENAI_API_KEY=your_api_key_here")
		os.Exit(1)
	}

	// Init OpenAI Client
	client := openai.NewClient(apiKey)

}
