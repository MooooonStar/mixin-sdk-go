package messenger

import (
	"context"
	"encoding/json"
)

func (m Messenger) ReadProfile(ctx context.Context) ([]byte, error) {
	return m.Request("GET", "/me", nil)
}

func (m Messenger) UpdateProfile(ctx context.Context, fullname, avatarBase64 string) ([]byte, error) {
	paras := make(map[string]interface{})
	if len(fullname) > 0 {
		paras["full_name"] = fullname
	}
	if len(avatarBase64) > 0 {
		paras["avatar_base64"] = avatarBase64
	}
	bt, _ := json.Marshal(paras)
	return m.Request("POST", "/me", bt)
}

func (m Messenger) UpdatePreference(ctx context.Context, receiveMessageSource, acceptConversationSource string) ([]byte, error) {
	paras := map[string]interface{}{}
	if len(receiveMessageSource) > 0 {
		paras["receive_message_source"] = receiveMessageSource
	}
	if len(acceptConversationSource) > 0 {
		paras["accept_conversation_source"] = acceptConversationSource
	}
	bt, _ := json.Marshal(paras)
	return m.Request("POST", "/me/preferences", bt)
}

func (m Messenger) FetchUsers(ctx context.Context, userIDs ...string) ([]byte, error) {
	bt, _ := json.Marshal(userIDs)
	return m.Request("POST", "/users/fetch", bt)
}

func (m Messenger) FetchUser(ctx context.Context, userID string) ([]byte, error) {
	return m.Request("GET", "/users/"+userID, nil)
}

// SearchUser search user; q is String: Mixin Id or Phone Numbe
func (m Messenger) SearchUser(ctx context.Context, q string) ([]byte, error) {
	return m.Request("GET", "/search/"+q, nil)
}

// FetchFriends fetch friends
func (m Messenger) FetchFriends(ctx context.Context) ([]byte, error) {
	return m.Request("GET", "/friends", nil)
}
