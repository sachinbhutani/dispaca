package commands

import (
	"github.com/sachinbhutani/dispaca/discord/commands/account"
	"github.com/sachinbhutani/dispaca/discord/commands/general"
	"github.com/sachinbhutani/sapphire"
)

//Init - Initialize bot commands
func Init(bot *sapphire.Bot) {
	//general commands
	//ping command
	bot.AddCommand(sapphire.NewCommand("ping", "General", general.Ping).SetDescription("Ping Pong!"))

	//account related commands
	accountInfo := sapphire.NewCommand("account", "Account", account.Info).SetDescription("Get alpaca account information")
	accountInfo.AddAliases("accountinfo", "accinfo")
	bot.AddCommand(accountInfo)
	mode := sapphire.NewCommand("mode", "Account", account.Mode).SetDescription("Account is in Paper trading or Live trading mode")
	bot.AddCommand(mode)
	balance := sapphire.NewCommand("balance", "Account", account.Balance).SetDescription("Account Balances")
	balance.AddAliases("bal")
	bot.AddCommand(balance)

}
