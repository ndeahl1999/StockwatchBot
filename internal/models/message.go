package message

//BotMessage json format for sending a groupme bot message
type BotMessage struct {
	BotID   string `json:"bot_id"`
	Message string `json:"text"`
}
