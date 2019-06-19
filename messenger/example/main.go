package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/MooooonStar/mixin-sdk-go/messenger"
)

type MyListener struct {
	*messenger.Messenger
}

// interface to implement if you want to handle the message
func (l *MyListener) OnMessage(ctx context.Context, msg messenger.MessageView, userId string) error {
	data, err := base64.StdEncoding.DecodeString(msg.Data)
	if err != nil {
		return err
	}
	if msg.Category == "SYSTEM_ACCOUNT_SNAPSHOT" {
		var transfer messenger.TransferView
		if err := json.Unmarshal(data, &transfer); err != nil {
			return err
		}
		log.Println("I got a snapshot: ", transfer)
		return l.SendPlainText(ctx, msg.ConversationId, msg.UserId, string(data))
	} else {
		log.Printf("I got a message, it said: %s", string(data))
		return l.SendPlainText(ctx, msg.ConversationId, msg.UserId, string(data))
	}
}

func main() {
	ctx := context.Background()
	m := messenger.NewMessenger(messenger.UserId, messenger.SessionId, messenger.PrivateKey)
	//replace with your own listener, like &MyListener{m}
	//go m.Run(ctx, messenger.DefaultBlazeListener)
	go m.Run(ctx, &MyListener{m})

	// //your mixin user id, can get from  "Search User"
	snow := "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"

	//must create conversation first. If have created before, skip this step.
	conversation, err := m.CreateConversation(ctx, messenger.CategoryContact,
		messenger.Participant{UserID: snow},
	)
	if err != nil {
		log.Println("create error", err)
	}

	if err := m.SendPlainText(ctx, conversation.ID, snow, "please send me a message and transfer some CNB to me."); err != nil {
		log.Println("send text error:", err)
	}

	select {}
}
