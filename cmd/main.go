package main

import "github.com/joho/godotenv"

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}
}
