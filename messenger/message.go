package messenger

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"image"
	_ "image/jpeg"
	_ "image/png"
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
	Size         int64  `json:"size"`
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

func (b *Messenger) SendPlainData(ctx context.Context, conversationId, recipientId string, raw Multimedia) error {
	data, _ := json.Marshal(raw)
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

//do not work yet
func (b *Messenger) SendPlainMessages(ctx context.Context, conversationId, recipientId string, video Multimedia) error {
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
