package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/fox-one/mixin-sdk/messenger"
	"github.com/fox-one/mixin-sdk/utils"
)

func main() {
	ctx := context.Background()
	m := messenger.NewMessenger(ClientID, SessionID, PINToken, SessionKey)
	//replace with your own listener
	go m.Run(ctx, messenger.DefaultBlazeListener{})

	user := "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"
	conversationId := utils.UniqueConversationId(ClientID, user)
	if err := m.SendPlainText(ctx, conversationId, user, "please send me a message."); err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("../donate.png")
	if err != nil {
		log.Fatal(err)
	}
	bt, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.SendImage(ctx, conversationId, user, bt); err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)
}
