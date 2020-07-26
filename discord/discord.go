package discord

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/sachinbhutani/dispaca/discord/commands"
	"github.com/sachinbhutani/dispaca/models"
	"github.com/sachinbhutani/sapphire"
)

//DG - discord go session
//exported in case direct access is required to discordgo functions
var DG *discordgo.Session

//BOT - Dispaca BOT object
var BOT *sapphire.Bot

//Init - Initialize the discord client
func Init() {
	var err error
	//dicordgo requires a prefix of "Bot " before the token
	DG, err = discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(models.BotName + ": Error initializing discord client, please check your token")
	}
	BOT = sapphire.New(DG)
	BOT.SetPrefix(os.Getenv("BOT_PREFIX"))
	BOT.LoadBuiltins() // Loads builtin commands.
	fmt.Print(models.BotName + ": connecting to discord..")
	err = BOT.Connect()
	if err != nil {
		log.Fatalln(models.BotName+": ERROR while connecting to discord: ", err)
	}
	fmt.Println("...connected!")
	// we will wait based on term signal or web client in the main program. bot.Wait is not required.
	fmt.Println(models.BotName + ": initializing commands")
	commands.Init(BOT)
}
