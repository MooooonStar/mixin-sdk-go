package messenger

import (
	"context"
	"encoding/json"

	"github.com/fox-one/mixin-sdk/utils"
)

type Attachment struct {
	AttachmentId string `json:"attachment"`
	UploadUrl    string `json:"upload_url"`
	ViewUrl      string `json:"view_url"`
}

func (m Messenger) CreateAttachment(ctx context.Context) (*Attachment, error) {
	data, err := m.Request(ctx, "POST", "/attachment", nil)
	if err != nil {
		return nil, requestError(err)
	}

	var resp struct {
		Attachment Attachment `json:"data"`
	}
	err = json.Unmarshal(data, &resp)
	return &resp.Attachment, err
}

func (m Messenger) Upload(ctx context.Context, file []byte) (string, string, error) {
	attachment, err := m.CreateAttachment(ctx)
	if err != nil {
		return "", "", err
	}

	req, err := utils.NewRequest(attachment.UploadUrl, "PUT", string(file))
	if err != nil {
		return "", "", err
	}
	req.Header.Set("x-amz-acl", "public-read")

	_, err = utils.DoRequest(req)
	return attachment.AttachmentId, attachment.ViewUrl, nil
}
