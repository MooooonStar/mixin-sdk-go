package messenger

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"io"
	"strings"

	uuid "github.com/satori/go.uuid"
)

const (
	CategoryGroup   = "GROUP"
	CategoryContact = "CONTACT"
	ActionAdd       = "ADD"
	ActionRemove    = "REMOVE"
	ActionJoin      = "JOIN"
	ActionExit      = "EXIT"
	ActionRole      = "ROLE"
	RoleAdmin       = "ADMIN"
)

// Participant conversation participant
type Participant struct {
	Type      string `json:"type"`
	UserID    string `json:"user_id"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	Action    string `json:"action,omitempty"`
}

// Conversation conversation
type Conversation struct {
	ID        string `json:"conversation_id"`
	CreatorID string `json:"creator_id"`
	CreatedAt string `json:"created_at"`
	Category  string `json:"category"`

	Name         string `json:"name,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	Announcement string `json:"announcement,omitempty"`

	CodeID  string `json:"code_id"`  //QR code id
	CodeURL string `json:"code_url"` //QR code url

	Participants []Participant `json:"participants"`
}

func UniqueConversationId(userId, recipientId string) string {
	minId, maxId := userId, recipientId
	if strings.Compare(userId, recipientId) > 0 {
		maxId, minId = userId, recipientId
	}
	h := md5.New()
	io.WriteString(h, minId)
	io.WriteString(h, maxId)
	sum := h.Sum(nil)
	sum[6] = (sum[6] & 0x0f) | 0x30
	sum[8] = (sum[8] & 0x3f) | 0x80
	return uuid.FromBytesOrNil(sum).String()
}

// create a GROUP or CONTACT conversation
func (m Messenger) CreateConversation(ctx context.Context, category string, participants ...Participant) (*Conversation, error) {
	conversationId := uuid.Must(uuid.NewV4()).String()
	if category == CategoryContact && len(participants) == 1 {
		conversationId = UniqueConversationId(m.UserId, participants[0].UserID)
	}

	params, err := json.Marshal(map[string]interface{}{
		"category":        category,
		"conversation_id": conversationId,
		"participants":    participants,
	})
	if err != nil {
		return nil, err
	}

	body, err := m.Request("POST", "/conversations", params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Data Conversation `json:"data"`
	}
	err = json.Unmarshal(body, &resp)
	return &resp.Data, err
}

//read the info of the conversation by id
func (m Messenger) ReadConversation(ctx context.Context, conversationId string) (*Conversation, error) {
	body, err := m.Request("GET", "/conversations/"+conversationId, nil)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Data Conversation `json:"data"`
	}
	err = json.Unmarshal(body, &resp)
	return &resp.Data, err
}

// do not work yet
func (m Messenger) ModifyConversation(ctx context.Context, conversationId string, participants ...Participant) (*Conversation, error) {
	params, err := json.Marshal(map[string]interface{}{
		"category":        CategoryGroup,
		"conversation_id": conversationId,
		"participants":    participants,
	})
	if err != nil {
		return nil, err
	}

	body, err := m.Request("POST", "/conversations", params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Data Conversation `json:"data"`
	}
	err = json.Unmarshal(body, &resp)
	return &resp.Data, err
}
