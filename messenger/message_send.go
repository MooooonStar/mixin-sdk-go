package messenger

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
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
	Duration     int64  `json:"duration,omitempty"`
	Name         string `json:"name,omitempty"`
}

type AppCard struct {
	IconUrl     string `json:"icon_url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Action      string `json:"action"`
}

type PlainMessage struct {
	ConversationID string `json:"conversation_id"`
	RecipentID     string `json:"recipient_id"`
	MessageID      string `json:"message_id"`
	Category       string `json:"category"`
	Data           string `json:"data"`
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
	if err := writeMessageAndWait(ctx, b.mc, "CREATE_MESSAGE", params); err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

// should have mime_type,width,height in image
func (b *Messenger) SendPlainImage(ctx context.Context, conversationId, recipientId string, image Multimedia) error {
	data, _ := json.Marshal(image)
	params := map[string]interface{}{
		"conversation_id": conversationId,
		"recipient_id":    recipientId,
		"message_id":      UuidNewV4().String(),
		"category":        "PLAIN_IMAGE",
		"data":            base64.StdEncoding.EncodeToString(data),
	}
	if err := writeMessageAndWait(ctx, b.mc, "CREATE_MESSAGE", params); err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

//should have name,size,mime_type in raw
func (b *Messenger) SendPlainData(ctx context.Context, conversationId, recipientId string, raw Multimedia) error {
	data, _ := json.Marshal(raw)
	params := map[string]interface{}{
		"conversation_id": conversationId,
		"recipient_id":    recipientId,
		"message_id":      UuidNewV4().String(),
		"category":        "PLAIN_DATA",
		"data":            base64.StdEncoding.EncodeToString(data),
	}
	if err := writeMessageAndWait(ctx, b.mc, "CREATE_MESSAGE", params); err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

//send sticker to recipientId, a valid sticker_id: b14bc6e3-b1ac-45fd-a5e2-60340c9880ef
func (b *Messenger) SendPlainSticker(ctx context.Context, conversationId, recipientId string, stickerID string) error {
	format := map[string]interface{}{"sticker_id": stickerID}
	data, _ := json.Marshal(format)
	params := map[string]interface{}{
		"conversation_id": conversationId,
		"recipient_id":    recipientId,
		"message_id":      UuidNewV4().String(),
		"category":        "PLAIN_STICKER",
		"data":            base64.StdEncoding.EncodeToString(data),
	}
	if err := writeMessageAndWait(ctx, b.mc, "CREATE_MESSAGE", params); err != nil {
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
	if err := writeMessageAndWait(ctx, b.mc, "CREATE_MESSAGE", params); err != nil {
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
	err := writeMessageAndWait(ctx, b.mc, "CREATE_MESSAGE", params)
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
	err := writeMessageAndWait(ctx, b.mc, "CREATE_MESSAGE", params)
	if err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

//should have mime_type,width,height,size,duration(ms) in multimeida
func (b *Messenger) SendPlainVideo(ctx context.Context, conversationId, recipientId string, video Multimedia) error {
	data, _ := json.Marshal(video)
	params := map[string]interface{}{
		"conversation_id": conversationId,
		"recipient_id":    recipientId,
		"message_id":      UuidNewV4().String(),
		"category":        "PLAIN_VIDEO",
		"data":            base64.StdEncoding.EncodeToString(data),
	}
	err := writeMessageAndWait(ctx, b.mc, "CREATE_MESSAGE", params)
	if err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

// send content to multi-user
func (b *Messenger) SendPlainMessages(ctx context.Context, messages ...PlainMessage) error {
	params := map[string]interface{}{"messages": messages}
	err := writeMessageAndWait(ctx, b.mc, "CREATE_PLAIN_MESSAGES", params)
	if err != nil {
		return BlazeServerError(ctx, err)
	}
	return nil
}

// send content to multi-user
func (b *Messenger) SendGroupMessage(ctx context.Context, content string, recipientId ...string) error {
	messages := make([]PlainMessage, 0)
	for _, recipient := range recipientId {
		message := PlainMessage{
			ConversationID: UniqueConversationId(b.UserId, recipient),
			RecipentID:     recipient,
			MessageID:      UuidNewV4().String(),
			Category:       "PLAIN_TEXT",
			Data:           base64.StdEncoding.EncodeToString([]byte(content)),
		}
		messages = append(messages, message)
	}
	return b.SendPlainMessages(ctx, messages...)
}

// send image in one step, upload to s3 first then to user.
func (b *Messenger) SendImage(ctx context.Context, conversationId, recipientId string, r io.Reader) error {
	bt, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	conf, format, err := image.DecodeConfig(bytes.NewReader(bt))
	if err != nil {
		return err
	}

	id, _, err := b.Upload(ctx, bytes.NewReader(bt))
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

// send video file to recipientId. I do not find grace package to get video info, so  I use command ffprobe
func (b *Messenger) SendVideo(ctx context.Context, conversationId, recipientId string, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	id, _, err := b.Upload(ctx, file)
	if err != nil {
		return err
	}

	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", filename)
	info, err := cmd.Output()
	if err != nil {
		return err
	}
	var Resp struct {
		Streams []struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"streams"`
		Format struct {
			Size     string `json:"size"`
			Duration string `json:"duration"`
		} `json:"format"`
	}

	err = json.Unmarshal(info, &Resp)
	if err != nil {
		return err
	}

	var width, height int
	for _, stream := range Resp.Streams {
		if stream.Height > 0 && stream.Width > 0 {
			width, height = stream.Width, stream.Height
			break
		}
	}

	size, _ := strconv.Atoi(Resp.Format.Size)
	duration, _ := strconv.ParseFloat(Resp.Format.Duration, 64)

	pos := strings.LastIndex(filename, ".")
	video := Multimedia{
		AttachmentID: id,
		MimeType:     "video/" + strings.ToLower(filename[pos:]),
		Width:        width,
		Height:       height,
		Size:         int64(size),
		Duration:     int64(math.Ceil(duration)) * 1000,
	}
	return b.SendPlainVideo(ctx, conversationId, recipientId, video)
}

func (m *Messenger) SendFile(ctx context.Context, conversationId, recipientId string, filename, mimeType string, r io.Reader) error {
	bt, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	id, _, err := m.Upload(ctx, bytes.NewReader(bt))
	if err != nil {
		return err
	}
	raw := Multimedia{
		Name:         filename,
		AttachmentID: id,
		Size:         int64(len(bt)),
		MimeType:     mimeType,
	}
	return m.SendPlainData(ctx, conversationId, recipientId, raw)
}
