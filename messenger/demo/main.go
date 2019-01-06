package main

import (
	"context"
	"log"

	"github.com/MooooonStar/mixin-sdk/messenger"
)

func main() {
	ctx := context.Background()
	m := messenger.NewMessenger(ClientID, SessionID, PINToken, SessionKey)
	//replace with your own listener
	go m.Run(ctx, messenger.DefaultBlazeListener{})

	snow := "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"
	soon := "cd345a58-2c40-4519-9533-d50a1e1b8238"

	//must create conversation first. If have created before, skip this step.
	if _, err := m.CreateConversation(ctx, messenger.CategoryContact, messenger.Participant{UserID: soon}); err != nil {
		log.Println("create conversation error", err)
	}
	if _, err := m.CreateConversation(ctx, messenger.CategoryContact, messenger.Participant{UserID: snow}); err != nil {
		log.Println("create conversation error", err)
	}
	conversation, err := m.CreateConversation(ctx, messenger.CategoryGroup,
		messenger.Participant{UserID: soon},
		messenger.Participant{UserID: snow},
	)
	if err != nil {
		log.Println("create error", err)
	}

	if err := m.SendPlainText(ctx, conversation.ID, snow, "please send me a message."); err != nil {
		log.Println("send text error:", err)
	}

	if err := m.SendImage(ctx, conversation.ID, snow, "../donate.png"); err != nil {
		log.Println("send image error:", err)
	}

	if err := m.SendVideo(ctx, conversation.ID, snow, "../123.mp4"); err != nil {
		log.Println("send video error", err)
	}

	for {
	}
}
