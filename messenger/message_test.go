package messenger

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"testing"
)

var (
	m                          *Messenger
	ctx                        context.Context
	conversationId, snow string
)

func init() {
	m = NewMessenger(UserId, SessionId, PrivateKey)
	ctx = context.Background()
	go m.Run(ctx, DefaultBlazeListener)
	//replace with your own mixin messenger wallet id, which can get from network.TestSearchUser
	snow = "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"

	conversation, err := m.CreateConversation(ctx, CategoryContact, Participant{UserID: snow})
	if err != nil {
		panic(err)
	}
	conversationId = conversation.ID
}

func TestSendText(t *testing.T) {
	err := m.SendPlainText(ctx, conversationId, snow, "hello!")
	if err != nil {
		t.Error(err)
	}
}

func TestSendAppCard(t *testing.T) {
	card := AppCard{Title: "CNB", Description: "Chui Niu Bi", Action: "http://www.google.cn",
		IconUrl: "https://images.mixin.one/0sQY63dDMkWTURkJVjowWY6Le4ICjAFuu3ANVyZA4uI3UdkbuOT5fjJUT82ArNYmZvVcxDXyNjxoOv0TAYbQTNKS=s128"}
	if err := m.SendAppCard(ctx, conversationId, snow, card); err != nil {
		t.Error(err)
	}
}

func TestSendAppButton(t *testing.T) {
	google := Button{Label: "google", Color: "#ABABAB", Action: "https://www.google.cn"}
	baidu := Button{Label: "baidu", Color: "#BABABA", Action: "https://www.baidu.com"}
	if err := m.SendAppButtons(ctx, conversationId, snow, google, baidu); err != nil {
		t.Error(err)
	}
}

func TestSendContact(t *testing.T) {
	if err := m.SendPlainContact(ctx, conversationId, snow, "c7ff704e-1a74-4f12-b05c-7a2be955a782"); err != nil {
		t.Error(err)
	}
}

func TestSendSticker(t *testing.T) {
	if err := m.SendPlainSticker(ctx, conversationId, snow, "b14bc6e3-b1ac-45fd-a5e2-60340c9880ef"); err != nil {
		t.Error(err)
	}
}

func TestSendPlainImage(t *testing.T) {
	file, err := os.Open("donate.png")
	if err != nil {
		t.Error(err)
	}
	bt, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
	}
	id, _, err := m.Upload(ctx, bytes.NewReader(bt))
	if err != nil {
		t.Error(err)
	}

	image := Multimedia{
		AttachmentID: id,
		MimeType:     "image/png",
		Width:        256,
		Height:       256,
	}
	if err := m.SendPlainImage(ctx, conversationId, snow, image); err != nil {
		t.Error(err)
	}
}

func TestSendVideo(t *testing.T) {
	err := m.SendVideo(ctx, conversationId, snow, "sample.mp4")
	if err != nil {
		t.Error(err)
	}
}

func TestSendImage(t *testing.T) {
	file, err := os.Open("donate.png")
	if err != nil {
		t.Error(err)
	}
	err = m.SendImage(ctx, conversationId, snow, file)
	if err != nil {
		t.Error(err)
	}
}

func TestSendFile(t *testing.T) {
	filename := "demo.pdf"
	file, err := os.Open(filename)
	if err != nil {
		t.Error(err)
	}
	err = m.SendFile(ctx, conversationId, snow, filename, "application/pdf", file)
	if err != nil {
		t.Error(err)
	}
}

func TestSendPlainData(t *testing.T) {
	file, err := os.Open("123.mp4")
	if err != nil {
		t.Error(err)
	}
	bt, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
	}
	id, _, err := m.Upload(ctx, bytes.NewBuffer(bt))
	if err != nil {
		t.Error(err)
	}
	media := Multimedia{
		Name:         "123.mp4",
		AttachmentID: id,
		Size:         int64(len(bt)),
		MimeType:     "video/mp4",
	}
	if err := m.SendPlainData(ctx, conversationId, snow, media); err != nil {
		t.Error(err)
	}
}

func TestSendGroupMessage(t *testing.T) {
	err := m.SendGroupMessage(ctx, "hello world", snow, snow)
	if err != nil {
		panic(err)
	}
}
