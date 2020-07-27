package account

import (
	"strconv"

	"github.com/sachinbhutani/dispaca/alpaca"
	"github.com/sachinbhutani/dispaca/discord/commands/general"
	"github.com/sachinbhutani/sapphire"
)

//Info command - return basic account information, no balances or amounts are returned with this command
func Info(ctx *sapphire.CommandContext) {
	acct, err := alpaca.Client.GetAccount()
	if err != nil {
		general.CommandError(ctx, "Error while fetching account information from Alpaca", err.Error())
		return
	}
	desc := "Account# " + acct.AccountNumber + "\nID: " + acct.ID + "\nStatus: " + acct.Status + "\n Currency: " + acct.Currency
	blocks := "\nAccount Blocked: " + general.ShowBool(acct.AccountBlocked) + "  Trading Blocked: " + general.ShowBool(acct.TradingBlocked) + "  Transfers Blocked: " + general.ShowBool(acct.TransfersBlocked)
	daytrades := "\nPDT: " + general.ShowBool(acct.PatternDayTrader) + " Day Trade Count: " + strconv.FormatInt(acct.DaytradeCount, 10)
	footer := "Account is in LIVE trading mode"
	if alpaca.PaperTrade == true {
		footer = "Account is in PAPER trading mode"
	}
	e := sapphire.NewEmbed().SetTitle("Alpaca Account Information").SetDescription(desc).SetColor(general.Blue).AddField("Blocks", blocks).AddField("Day Trades", daytrades).SetFooter(footer)
	ctx.ReplyEmbedNoEdit(e.MessageEmbed)
}

//Mode - Check if the account is in paper trade mode of live trading mode
func Mode(ctx *sapphire.CommandContext) {
	//TODO: change to URL check
	if alpaca.PaperTrade == true {
		ctx.ReplyEmbedNoEdit(sapphire.NewEmbed().SetTitle("Paper Trade").SetDescription("PAPER trading mode is On").SetColor(general.Yellow).Build())
		return
	}
	ctx.ReplyEmbedNoEdit(sapphire.NewEmbed().SetTitle("Live Trade").SetDescription("LIVE trading mode is On").SetColor(general.Green).Build())
}

//Balance - Display account balances
func Balance(ctx *sapphire.CommandContext) {
	acct, err := alpaca.Client.GetAccount()
	if err != nil {
		general.CommandError(ctx, "Error while fetching account information from Alpaca", err.Error())
		return
	}
	desc := "Account# " + acct.AccountNumber + "\n Currency: " + acct.Currency + "\nBuying Power: " + acct.BuyingPower.String() + +"\nDay Trade Buying Power: " + acct.DaytradingBuyingPower.String() + "\nReg T Buying Power:" + acct.RegTBuyingPower.String()
	portTitle := "Portfolio (MTM): " + acct.Equity.String()
	portDetail := "Cash: " + acct.Cash + "\nLong Market Value: " + acct.LongMarketValue.String() + "\nShort Market Value: " + acct.ShortMarketValue.String()
	lastEquity := "Last Equity Value: " + acct.LastEquity.String()
	footer := "Account is in LIVE trading mode"
	if alpaca.PaperTrade == true {
		footer = "Account is in PAPER trading mode"
	}
	e := sapphire.NewEmbed().SetTitle("Alpaca Account Balance Details").SetDescription(desc).SetColor(general.Blue).AddField(portTitle, portDetail).AddField("Previous Day Equity Value", lastEquity).SetFooter(footer)
	ctx.ReplyEmbedNoEdit(e.MessageEmbed)
}
