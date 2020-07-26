package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sachinbhutani/dispaca/alpaca"
	"github.com/sachinbhutani/dispaca/discord"
	"github.com/sachinbhutani/dispaca/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(models.BotName+": Starting "+models.BotName+" version: ", models.Version)
	alpaca.Init()
	discord.Init()
	fmt.Println(models.BotName + ": ready for action!")
	discord.BOT.Wait()
}
