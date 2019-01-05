package messenger

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/fox-one/mixin-sdk/utils"
)

type DefaultBlazeListener struct{}

// interface to implement if you want to handle the message
func (l DefaultBlazeListener) OnMessage(ctx context.Context, msg MessageView, userId string) error {
	data, err := base64.StdEncoding.DecodeString(msg.Data)
	if err != nil {
		return err
	}
	if msg.Category == "SYSTEM_ACCOUNT_SNAPSHOT" && msg.UserId != ClientID {
		var transfer TransferView
		if err := json.Unmarshal(data, &transfer); err != nil {
			return err
		}
		log.Println("I got a snapshot: ", transfer)
		return nil
	} else if msg.ConversationId == utils.UniqueConversationId(ClientID, msg.UserId) {
		log.Printf("I got a message, it said: %s", string(data))
		return nil
	}
	return nil
}
