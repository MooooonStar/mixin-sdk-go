package messenger

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
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

func TestSendSticker(t *testing.T) {
	if err := m.SendPlainSticker(ctx, conversationId, userID, "b14bc6e3-b1ac-45fd-a5e2-60340c9880ef"); err != nil {
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

func TestSendVideo(t *testing.T) {
	filename := "123.mp4"
	file, err := os.Open(filename)
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

	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams")
	cmd.Stdin = bytes.NewReader(bt)
	info, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
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
		panic(err)
	}

	var width, height int
	for _, stream := range Resp.Streams {
		if stream.Height > 0 && stream.Width > 0 {
			width, height = stream.Width, stream.Height
		}
	}

	size, _ := strconv.Atoi(Resp.Format.Size)
	duration, _ := strconv.ParseFloat(Resp.Format.Duration, 64)

	video := Multimedia{
		AttachmentID: id,
		MimeType:     "video/mp4",
		Width:        width,
		Height:       height,
		Size:         int64(size),
		Duration:     int64(duration) * 1000,
	}
	if err := m.SendPlainVideo(ctx, conversationId, userID, video); err != nil {
		panic(err)
	}
}

func TestSendImageOneStep(t *testing.T) {
	if err := m.SendImage(ctx, conversationId, userID, "donate.png"); err != nil {
		panic(err)
	}
}

func TestSendVideoOneStep(t *testing.T) {
	if err := m.SendVideo(ctx, conversationId, userID, "123.mp4"); err != nil {
		panic(err)
	}
}
