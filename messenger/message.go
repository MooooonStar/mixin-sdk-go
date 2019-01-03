package messenger

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/fox-one/mixin-sdk/utils"
	"github.com/hokaccha/go-prettyjson"
)

type Button struct {
	Label  string `json:"label"`
	Action string `json:"action"`
	Color  string `json:"color"`
}

type Multimedia struct {
	AttachmentID string `json:"attachment_id"`
	MimeType     string `json:"mime_type,omitempty"`
	Width        int    `json:"width,omitempty"`
	Height       int    `json:"height,omitempty"`
	Size         int64  `json:"size,omitempty"`
	Thumbnail    string `json:"thumbnail,omitempty"`
	Duration     int    `json:"duration,omitempty"`
	Name         string `json:"name,omitempty"`
}

type AppCard struct {
	IconUrl     string `json:"icon_url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Action      string `json:"action"`
}

// send a text messeage to recipientId
func (b *Messenger) SendPlainText(ctx context.Context, conversationId, recipientId string, content string) error {
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

// send a image to recipientId
func (b *Messenger) SendPlainImage(ctx context.Context, conversationId, recipientId string, image Multimedia) error {
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

// send image in one step, upload to s3 first then to user.
func (b *Messenger) SendImage(ctx context.Context, conversationId, recipientId string, data []byte) error {
	conf, format, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return err
	}

	id, _, err := b.Upload(ctx, data)
	if err != nil {
		return err
	}

	image := Multimedia{
		AttachmentID: id,
		MimeType:     "image/" + format,
		Width:        conf.Width,
		Height:       conf.Height,
	}
	return b.SendPlainImage(ctx, conversationId, recipientId, image)
}

//do not work yet
func (b *Messenger) SendPlainData(ctx context.Context, conversationId, recipientId string, raw Multimedia) error {
	data, _ := json.Marshal(raw)
	v, _ := prettyjson.Format(data)
	log.Println("data", string(v))
	params := map[string]interface{}{
		"conversation_id": conversationId,
		"recipient_id":    recipientId,
		"message_id":      UuidNewV4().String(),
		"category":        "PLAIN_DATA",
		"data":            base64.StdEncoding.EncodeToString(data),
	}
	if err := writeMessageAndWait(ctx, b.mc, createMessageAction, params); err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

//do not work yet
func (b *Messenger) SendPlainSticker(ctx context.Context, conversationId, recipientId string, name, ablumID string) error {
	format := map[string]interface{}{
		"name":     name,
		"album_id": ablumID,
	}
	data, _ := json.Marshal(format)
	params := map[string]interface{}{
		"conversation_id": conversationId,
		"recipient_id":    recipientId,
		"message_id":      UuidNewV4().String(),
		"category":        "PLAIN_STICKER",
		"data":            base64.StdEncoding.EncodeToString(data),
	}
	if err := writeMessageAndWait(ctx, b.mc, createMessageAction, params); err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

// share a contact to recipientId
func (b *Messenger) SendPlainContact(ctx context.Context, conversationId, recipientId, contactId string) error {
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

// send buttons to recipientId which can jump to a website when click
func (b *Messenger) SendAppButtons(ctx context.Context, conversationId, recipientId string, buttons ...Button) error {
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

// share app to  recipientId
func (b *Messenger) SendAppCard(ctx context.Context, conversationId, recipientId string, card AppCard) error {
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

//do not work yet
func (b *Messenger) SendPlainVideo(ctx context.Context, conversationId, recipientId string, video Multimedia) error {
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

// send content to multi-user
func (b *Messenger) SendPlainMessages(ctx context.Context, content string, recipientID ...string) error {
	messages := make([]interface{}, 0)
	for _, recipient := range recipientID {
		message := map[string]interface{}{
			"conversation_id": utils.UniqueConversationId(ClientID, recipient),
			"recipient_id":    recipient,
			"message_id":      UuidNewV4().String(),
			"category":        "PLAIN_TEXT",
			"data":            base64.StdEncoding.EncodeToString([]byte(content)),
		}
		messages = append(messages, message)
	}
	params := map[string]interface{}{"messages": messages}
	err := writeMessageAndWait(ctx, b.mc, "CREATE_PLAIN_MESSAGES", params)
	if err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}
