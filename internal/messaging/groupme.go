package groupme

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	config "github.com/ndeahl1999/StockwatchBot/internal/config"
	message "github.com/ndeahl1999/StockwatchBot/internal/models"
)

var (
	botID string
)

//InitializeBot used to initialize the groupme bot
func InitializeBot() {
	botID = os.Getenv("GROUPME_BOT_ID")

	if botID == "" {
		log.Fatal("ERROR: Groupme bot_id not set")
	}

	groupmeURL := config.GroupmeAPIURL

	if os.Getenv("GO_ENV") == "production" {
		if SendBotMessage("StockwatchBot server is up and running!") {
			log.Println("Sent startup message to " + groupmeURL + " with bot id " + botID)
		} else {
			log.Fatal("ERROR: There was a problem sending the bot message")
		}
	}

}

//SendBotMessage wrapper function for sending message with groupmebot
func SendBotMessage(messageString string) bool {
	botMessage := message.BotMessage{
		BotID:   botID,
		Message: messageString,
	}
	jsonData, _ := json.Marshal(botMessage)

	_, err := http.Post(config.GroupmeAPIURL+"/post", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
