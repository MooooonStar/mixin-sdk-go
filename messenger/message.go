package messenger

import (
	"context"
	"encoding/base64"
	"encoding/json"
)

type Button struct {
	Label  string `json:"label"`
	Action string `json:"action"`
	Color  string `json:"color"`
}

type Multimedia struct {
	AttachmentID string `json:"attachment_id"`
	MimeType     string `json:"mime_type"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Size         int    `json:"size"`
	Thumbnail    string `json:"thumbnail"`
	Duration     int    `json:"duration"`
	Name         string `json:"name"`
}

type AppCard struct {
	IconUrl     string `json:"icon_url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Action      string `json:"action"`
}

func (b *BlazeClient) SendPlainText(ctx context.Context, conversationId, recipientId string, content string) error {
	params := map[string]interface{}{
		"conversation_id": conversationId,
		"recipient_id":    recipientId,
		"message_id":      UuidNewV4().String(),
		"category":        "PLAIN_TEXT",
		"data":            base64.StdEncoding.EncodeToString([]byte(content)),
	}
	if err := writeMessageAndWait(ctx, b.mc, createMessageAction, params); err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

func (b *BlazeClient) SendPlainImage(ctx context.Context, conversationId, recipientId string, image Multimedia) error {
	data, _ := json.Marshal(image)
	params := map[string]interface{}{
		"conversation_id": conversationId,
		"recipient_id":    recipientId,
		"message_id":      UuidNewV4().String(),
		"category":        "PLAIN_IMAGE",
		"data":            base64.StdEncoding.EncodeToString(data),
	}
	if err := writeMessageAndWait(ctx, b.mc, createMessageAction, params); err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

func (b *BlazeClient) SendPlainData(ctx context.Context, msg MessageView, raw Multimedia) error {
	data, _ := json.Marshal(raw)
	params := map[string]interface{}{
		"conversation_id": msg.ConversationId,
		"recipient_id":    msg.UserId,
		"message_id":      UuidNewV4().String(),
		"category":        "PLAIN_DATA",
		"data":            base64.StdEncoding.EncodeToString(data),
	}
	if err := writeMessageAndWait(ctx, b.mc, createMessageAction, params); err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

func (b *BlazeClient) SendPlainSticker(ctx context.Context, msg MessageView, name, ablumID string) error {
	format := map[string]interface{}{
		"name":     name,
		"album_id": ablumID,
	}
	data, _ := json.Marshal(format)
	params := map[string]interface{}{
		"conversation_id": msg.ConversationId,
		"recipient_id":    msg.UserId,
		"message_id":      UuidNewV4().String(),
		"category":        "PLAIN_STICKER",
		"data":            base64.StdEncoding.EncodeToString(data),
	}
	if err := writeMessageAndWait(ctx, b.mc, createMessageAction, params); err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

func (b *BlazeClient) SendPlainContact(ctx context.Context, conversationId, recipientId, contactId string) error {
	format := map[string]string{"user_id": contactId}
	data, _ := json.Marshal(format)
	params := map[string]interface{}{
		"conversation_id": conversationId,
		"recipient_id":    recipientId,
		"message_id":      UuidNewV4().String(),
		"category":        "PLAIN_CONTACT",
		"data":            base64.StdEncoding.EncodeToString(data),
	}
	if err := writeMessageAndWait(ctx, b.mc, createMessageAction, params); err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

func (b *BlazeClient) SendAppButtons(ctx context.Context, conversationId, recipientId string, buttons ...Button) error {
	data, _ := json.Marshal(buttons)
	params := map[string]interface{}{
		"conversation_id": conversationId,
		"recipient_id":    recipientId,
		"message_id":      UuidNewV4().String(),
		"category":        "APP_BUTTON_GROUP",
		"data":            base64.StdEncoding.EncodeToString(data),
	}
	err := writeMessageAndWait(ctx, b.mc, createMessageAction, params)
	if err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

func (b *BlazeClient) SendAppCard(ctx context.Context, conversationId, recipientId string, card AppCard) error {
	data, _ := json.Marshal(card)
	params := map[string]interface{}{
		"conversation_id": conversationId,
		"recipient_id":    recipientId,
		"message_id":      UuidNewV4().String(),
		"category":        "APP_CARD",
		"data":            base64.StdEncoding.EncodeToString(data),
	}
	err := writeMessageAndWait(ctx, b.mc, createMessageAction, params)
	if err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

func (b *BlazeClient) SendPlainVideo(ctx context.Context, conversationId, recipientId string, video Multimedia) error {
	data, _ := json.Marshal(video)
	params := map[string]interface{}{
		"conversation_id": conversationId,
		"recipient_id":    recipientId,
		"message_id":      UuidNewV4().String(),
		"category":        "PLAIN_VIDEO",
		"data":            base64.StdEncoding.EncodeToString(data),
	}
	err := writeMessageAndWait(ctx, b.mc, createMessageAction, params)
	if err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

func (b *BlazeClient) SendPlainMessages(ctx context.Context, conversationId, recipientId string, video Multimedia) error {
	data, _ := json.Marshal(video)
	params := map[string]interface{}{
		"conversation_id": conversationId,
		"recipient_id":    recipientId,
		"message_id":      UuidNewV4().String(),
		"category":        "PLAIN_TEXT",
		"data":            base64.StdEncoding.EncodeToString(data),
	}
	err := writeMessageAndWait(ctx, b.mc, createMessageAction, params)
	if err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}
