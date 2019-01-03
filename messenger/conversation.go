package messenger

import (
	"context"
	"encoding/json"

	"github.com/fox-one/mixin-sdk/utils"
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

// create a GROUP or CONTACT conversation
func (m Messenger) CreateConversation(ctx context.Context, category string, participants ...Participant) (*Conversation, error) {
	conversationId := uuid.Must(uuid.NewV4()).String()
	if category == CategoryContact && len(participants) == 1 {
		conversationId = utils.UniqueConversationId(m.User.UserID, participants[0].UserID)
	}

	params, err := json.Marshal(map[string]interface{}{
		"category":        category,
		"conversation_id": conversationId,
		"participants":    participants,
	})
	if err != nil {
		return nil, err
	}

	body, err := m.Request(ctx, "POST", "/conversations", params)
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
	body, err := m.Request(ctx, "GET", "/conversations/"+conversationId, nil)
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

	body, err := m.Request(ctx, "POST", "/conversations", params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Data Conversation `json:"data"`
	}
	err = json.Unmarshal(body, &resp)
	return &resp.Data, err
}
