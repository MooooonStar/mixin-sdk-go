package messenger

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type Attachment struct {
	AttachmentId string `json:"attachment_id"`
	UploadUrl    string `json:"upload_url"`
	ViewUrl      string `json:"view_url"`
}

var httpClient *http.Client

func init() {
	httpClient = &http.Client{Timeout: 10 * time.Second}
}

//get the id, upload_url, view_url for your attachment
//the first step to upload something to mixin network or send multimedia (image, video  etc) to others
func (m Messenger) CreateAttachment(ctx context.Context) (*Attachment, error) {
	data, err := m.Request(ctx, "POST", "/attachments", nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Attachment Attachment `json:"data"`
	}
	err = json.Unmarshal(data, &resp)
	return &resp.Attachment, err
}

func (m Messenger) Upload(ctx context.Context, data []byte) (string, string, error) {
	attachment, err := m.CreateAttachment(ctx)
	if err != nil {
		return "", "", err
	}

	req, err := http.NewRequest(attachment.UploadUrl, "PUT", bytes.NewReader(data))
	if err != nil {
		return "", "", err
	}
	//must set those two headers, it's used for s3 signature
	req.Header.Set("x-amz-acl", "public-read")
	req.Header.Add("Content-Type", "application/octet-stream")

	_, err = httpClient.Do(req)
	if err != nil {
		return "", "", err
	}
	return attachment.AttachmentId, attachment.ViewUrl, nil
}
