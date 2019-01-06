package main

import (
	"context"

	"github.com/fox-one/mixin-sdk/messenger"
)

func main() {
	ctx := context.Background()
	m := messenger.NewMessenger(ClientID, SessionID, PINToken, SessionKey)
	//replace with your own listener
	m.Run(ctx, messenger.DefaultBlazeListener{})

	// user := "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"
	// //must create conversation first. If have created before, use its id to send a message.
	// participant := messenger.Participant{UserID: user, Action: messenger.ActionAdd}
	// conversation, err := m.CreateConversation(ctx, messenger.CategoryContact, participant)
	// if err != nil {
	// 	panic(err)
	// }
	// conversationID := conversation.ID
	// //conversationID := utils.UniqueConversationId(ClientID, user)
	// if err := m.SendPlainText(ctx, conversationID, user, "please send me a message."); err != nil {
	// 	log.Println(err)
	// }

	// file, err := os.Open("../donate.png")
	// if err != nil {
	// 	log.Println(err)
	// }
	// bt, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Println(err)
	// }
	// if err := m.SendImage(ctx, conversationID, user, bt); err != nil {
	// 	log.Println(err)
	// }
}
