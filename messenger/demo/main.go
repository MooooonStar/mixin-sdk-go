package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/hokaccha/go-prettyjson"

	"github.com/fox-one/mixin-sdk/messenger"
	"github.com/fox-one/mixin-sdk/mixin"
	"github.com/fox-one/mixin-sdk/utils"
)

type Handler struct {
	*messenger.Messenger
}

func (h Handler) OnMessage(ctx context.Context, msgView messenger.MessageView, userId string) error {
	log.Println("I received a msg", msgView)

	if msgView.Category != messenger.MessageCategoryPlainText {
		return nil
	}

	data, err := base64.StdEncoding.DecodeString(msgView.Data)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("I got your message, you said: %s", string(data))
	log.Println(msg)

	return h.SendPlainText(ctx, msgView, msg)
}

func (h Handler) Run(ctx context.Context) {
	for {
		if err := h.Loop(ctx, h); err != nil {
			log.Println("something is wrong", err)
			time.Sleep(1 * time.Second)
		}
	}
}

func (h Handler) Send(ctx context.Context, userId, content string) error {
	msgView := messenger.MessageView{
		ConversationId: utils.UniqueConversationId(ClientID, userId),
		UserId:         userId,
	}
	return h.SendPlainText(ctx, msgView, content)
}

func main() {
	ctx := context.Background()
	user := mixin.NewUser(ClientID, SessionID, PINToken, SessionKey)
	m := messenger.NewMessenger(user)

	participant := messenger.Participant{
		UserID: "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909",
		//Role:   messenger.RoleAdmin,
		Action: messenger.ActionAdd,
	}
	conversation, err := m.CreateConversation(ctx, messenger.CategoryContact, participant)
	if err != nil {
		log.Fatal("create conversation", err)
	}
	v, _ := prettyjson.Marshal(conversation)
	log.Println("conversation:", string(v))

	sample, err := m.ReadConversation(ctx, conversation.ConversationID)
	if err != nil {
		log.Fatal(err)
	}
	v, _ = prettyjson.Marshal(sample)
	log.Println("sample:", string(v))

	//h := Handler{m}
	//h.Run(ctx)
}
