package messenger

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
)

type DefaultBlazeListener struct{}

// interface to implement if you want to handle the message
func (l DefaultBlazeListener) OnMessage(ctx context.Context, msg MessageView, userId string) error {
	log.Println("I got a message: ", msg)
	data, err := base64.StdEncoding.DecodeString(msg.Data)
	if err != nil {
		return err
	}
	if msg.Category == "SYSTEM_ACCOUNT_SNAPSHOT" {
		var transfer TransferView
		if err := json.Unmarshal(data, &transfer); err != nil {
			return err
		}
		log.Println("I got a snapshot: ", transfer)
		return nil
	} else {
		log.Printf("I got a message, it said: %s", string(data))
		return nil
	}
}
