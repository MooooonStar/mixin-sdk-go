package messenger

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/fox-one/mixin-sdk/utils"
)

var m *Messenger
var conversationId, userID string
var ctx context.Context

func init() {
	m = NewMessenger(ClientID, SessionID, PINToken, SessionKey)
	ctx = context.Background()
	go m.Run(ctx, DefaultBlazeListener{})
	userID = "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"
	conversationId = utils.UniqueConversationId(ClientID, userID)
}

func TestConversation(t *testing.T) {
	participant := Participant{UserID: userID}
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

	if err := m.SendPlainText(ctx, conversation.ID, userID, "go go go"); err != nil {
		panic(err)
	}
	time.Sleep(20 * time.Second)
}

func TestSendText(t *testing.T) {
	if err := m.SendPlainText(ctx, conversationId, userID, "hello!"); err != nil {
		panic(err)
	}
}

func TestSendAppCard(t *testing.T) {
	card := AppCard{Title: "CNB", Description: "Chui Niu Bi", Action: "http://www.google.cn",
		IconUrl: "https://images.mixin.one/0sQY63dDMkWTURkJVjowWY6Le4ICjAFuu3ANVyZA4uI3UdkbuOT5fjJUT82ArNYmZvVcxDXyNjxoOv0TAYbQTNKS=s128"}
	if err := m.SendAppCard(ctx, conversationId, userID, card); err != nil {
		panic(err)
	}
}
func TestSendAppButton(t *testing.T) {
	google := Button{Label: "google", Color: "#ABABAB", Action: "https://www.google.cn"}
	baidu := Button{Label: "baidu", Color: "#BABABA", Action: "https://www.baidu.com"}
	if err := m.SendAppButtons(ctx, conversationId, userID, google, baidu); err != nil {
		panic(err)
	}
}

func TestSendContact(t *testing.T) {
	if err := m.SendPlainContact(ctx, conversationId, userID, "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"); err != nil {
		panic(err)
	}
}

func TestSendImage(t *testing.T) {
	file, err := os.Open("donate.png")
	if err != nil {
		panic(err)
	}
	bt, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	id, _, err := m.Upload(ctx, bt)
	if err != nil {
		panic(err)
	}

	image := Multimedia{
		AttachmentID: id,
		MimeType:     "image/png",
		Width:        256,
		Height:       256,
	}
	if err := m.SendPlainImage(ctx, conversationId, userID, image); err != nil {
		panic(err)
	}
}

func TestSendImageOneStep(t *testing.T) {
	file, err := os.Open("donate.png")
	if err != nil {
		panic(err)
	}
	bt, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	if err := m.SendImage(ctx, conversationId, userID, bt); err != nil {
		panic(err)
	}
}
