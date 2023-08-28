package telego

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func SayBot(text string, ChatID int64) error {
	log.Println("SayBot", text)
	var msg sendMessageReqBody
	msg.ChatID = ChatID
	msg.Text = text
	reqBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// Send a post request with your token
	res, err := http.Post("https://api.telegram.org/bot"+Token+"/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}
