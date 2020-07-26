package general

import (
	"github.com/sachinbhutani/dispaca/models"
	"github.com/sachinbhutani/sapphire"
)

const (
	Red    int = 0xFF0000
	Green  int = 0x00FF00
	Blue   int = 0x0000FF
	Yellow int = 0xFFFF00
	Orange
	Cyan
)

//Ping Command Reply
func Ping(ctx *sapphire.CommandContext) {
	ctx.Reply("Pong! :llama: dispaca version:  " + models.Version + "\n:heartbeat: Heartbeat ping is " + ctx.Bot.Session.HeartbeatLatency().String())
}

//CommandError - send error message to discord with the input title and description
func CommandError(ctx *sapphire.CommandContext, title string, description string) {
	ctx.ReplyEmbedNoEdit(sapphire.NewEmbed().SetTitle(title).SetColor(Red).SetDescription(description).SetFooter(models.BotName + " " + models.Version).Build())
}

//ShowBool - convert bool to check or cross emojit
func ShowBool(x bool) string {
	if x == true {
		return " :white_check_mark: "
	}
	return " :regional_indicator_x: "
}
