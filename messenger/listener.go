package messenger

import (
	"context"
	"encoding/base64"
	"log"
)

type DefaultBlazeListener struct{}

func (l DefaultBlazeListener) OnMessage(ctx context.Context, msgView MessageView, userId string) error {
	data, err := base64.StdEncoding.DecodeString(msgView.Data)
	if err != nil {
		return err
	}
	log.Printf("I got your message, you said: %s", string(data))
	return nil
}
