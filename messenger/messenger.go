package messenger

import "github.com/fox-one/mixin-sdk/mixin"

// Messenger mixin messenger
type Messenger struct {
	*mixin.User
	*BlazeClient
}

// NewMessenger create messenger
func NewMessenger(user *mixin.User) *Messenger {
	return &Messenger{
		user,
		NewBlazeClient(user.UserID, user.SessionID, user.SessionKey),
	}
}

// func NewMessenger(clientID, sessionID, pinToken, sessionKey) *Messenger {
// 	user := mixin.NewUser(clientID, sessionID, pinToken, sessionKey)
// 	return NewMessenger(user)
// }
