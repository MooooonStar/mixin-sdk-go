package main

import (
	"context"
	"log"

	"github.com/MooooonStar/mixin-sdk-go/messenger"
)

func main() {
	ctx := context.Background()
	m := messenger.NewMessenger(UserId, SessionId, PrivateKey)
	//replace with your own listener
	go m.Run(ctx, messenger.DefaultBlazeListener{})

	snow := "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"

	//must create conversation first. If have created before, skip this step.
	if _, err := m.CreateConversation(ctx, messenger.CategoryContact, messenger.Participant{UserID: snow}); err != nil {
		log.Println("create conversation error", err)
	}
	conversation, err := m.CreateConversation(ctx, messenger.CategoryGroup,
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

	if err := m.SendVideo(ctx, conversation.ID, snow, "../sample.mp4"); err != nil {
		log.Println("send video error", err)
	}

	if err := m.SendFile(ctx, conversation.ID, snow, "../demo.pdf", "application/pdf"); err != nil {
		log.Println("send video error", err)
	}

	for {
	}
}
