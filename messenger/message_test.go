package messenger

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

var m *Messenger
var conversationId, snow string
var ctx context.Context

func init() {
	m = NewMessenger(UserID, SessionID, PinToken, PrivateKey)
	ctx = context.Background()
	go m.Run(ctx, DefaultBlazeListener{})
	snow = "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"
	conversationId = UniqueConversationId(m.UserId, snow)
}

func TestConversation(t *testing.T) {
	participant := Participant{UserID: snow}
	conversation, err := m.CreateConversation(ctx, CategoryGroup, participant)
	if err != nil {
		panic(err)
	}
	sample, err := m.ReadConversation(ctx, conversation.ID)
	if err != nil {
		panic(err)
	}
	log.Println("conversation:", conversation)
	log.Println("sample:", sample)

	if err := m.SendPlainText(ctx, conversation.ID, snow, "go go go"); err != nil {
		panic(err)
	}
	time.Sleep(20 * time.Second)
}

func TestSendText(t *testing.T) {
	if err := m.SendPlainText(ctx, conversationId, snow, "hello!"); err != nil {
		panic(err)
	}
}

func TestSendAppCard(t *testing.T) {
	card := AppCard{Title: "CNB", Description: "Chui Niu Bi", Action: "http://www.google.cn",
		IconUrl: "https://images.mixin.one/0sQY63dDMkWTURkJVjowWY6Le4ICjAFuu3ANVyZA4uI3UdkbuOT5fjJUT82ArNYmZvVcxDXyNjxoOv0TAYbQTNKS=s128"}
	if err := m.SendAppCard(ctx, conversationId, snow, card); err != nil {
		panic(err)
	}
}
func TestSendAppButton(t *testing.T) {
	google := Button{Label: "google", Color: "#ABABAB", Action: "https://www.google.cn"}
	baidu := Button{Label: "baidu", Color: "#BABABA", Action: "https://www.baidu.com"}
	if err := m.SendAppButtons(ctx, conversationId, snow, google, baidu); err != nil {
		panic(err)
	}
}

func TestSendContact(t *testing.T) {
	if err := m.SendPlainContact(ctx, conversationId, snow, "c7ff704e-1a74-4f12-b05c-7a2be955a782"); err != nil {
		panic(err)
	}
}

func TestSendSticker(t *testing.T) {
	if err := m.SendPlainSticker(ctx, conversationId, snow, "b14bc6e3-b1ac-45fd-a5e2-60340c9880ef"); err != nil {
		panic(err)
	}
}

func TestSendPlainImage(t *testing.T) {
	file, err := os.Open("donate.png")
	if err != nil {
		log.Fatal("open", err)
	}
	bt, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("read", err)
	}
	id, _, err := m.Upload(ctx, bt)
	if err != nil {
		log.Fatal("upload", err)
	}

	image := Multimedia{
		AttachmentID: id,
		MimeType:     "image/png",
		Width:        256,
		Height:       256,
	}
	if err := m.SendPlainImage(ctx, conversationId, snow, image); err != nil {
		log.Fatal("send", err)
	}
}

func TestSendVideo(t *testing.T) {
	if err := m.SendVideo(ctx, conversationId, snow, "123.mp4"); err != nil {
		panic(err)
	}
}

func TestSendImage(t *testing.T) {
	if err := m.SendImage(ctx, conversationId, snow, "donate.png"); err != nil {
		panic(err)
	}
}

func TestSendPlainData(t *testing.T) {
	file, err := os.Open("123.mp4")
	if err != nil {
		panic(err)
	}
	bt, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	id, _, err := m.Upload(context.Background(), bt)
	if err != nil {
		panic(err)
	}
	media := Multimedia{
		Name:         "123.mp4",
		AttachmentID: id,
		Size:         int64(len(bt)),
		MimeType:     "video/mp4",
	}
	if err := m.SendPlainData(ctx, conversationId, snow, media); err != nil {
		panic(err)
	}
}

func TestSendGroupMessage(t *testing.T) {
	soon := "cd345a58-2c40-4519-9533-d50a1e1b8238"
	err := m.SendGroupMessage(ctx, "hello world", snow, soon)
	if err != nil {
		panic(err)
	}
}
