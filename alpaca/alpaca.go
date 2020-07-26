package alpaca

import (
	"fmt"
	"os"
	"strings"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/sachinbhutani/dispaca/models"
)

//Client - Alpaca API client
var Client *alpaca.Client

//PaperTrade - set to true if the account is in paper trading mode
var PaperTrade bool

//Init - Initialize alpaca client
func Init() {

	fmt.Println(models.BotName + ": Connecting to Alpaca ")
	//TODO: determine paper trade mode from URL instead of the env variable.
	if strings.ToLower(os.Getenv("PAPER_TRADE")) == "true" {
		PaperTrade = true
		alpaca.SetBaseUrl(os.Getenv("APCA_API_BASE_URL"))
	}
	//.env file sets the required environment variabled for credentials
	Client = alpaca.NewClient(common.Credentials())
}
