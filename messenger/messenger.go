package messenger

import (
	"context"
	"log"
	"time"

	mixin "github.com/MooooonStar/mixin-sdk/network"
)

// Messenger mixin messenger
type Messenger struct {
	*mixin.User
	*BlazeClient
}

// NewMessenger create messenger
func NewMessenger(clientID, sessionID, pinToken, privateKey string) *Messenger {
	user := mixin.NewUser(clientID, sessionID, pinToken, privateKey)
	client := NewBlazeClient(clientID, sessionID, privateKey)
	return &Messenger{user, client}
}

func (m *Messenger) Run(ctx context.Context, listener BlazeListener) {
	for {
		if err := m.Loop(ctx, listener); err != nil {
			log.Println("Blaze server error", err)
			time.Sleep(1 * time.Second)
		}
		m.BlazeClient = NewBlazeClient(ClientID, SessionID, SessionKey)
	}
}
