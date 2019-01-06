package network

type User struct {
	UserId         string `json:"user_id"`
	SessionId      string `json:"session_id"`
	PinToken       string `json:"pin_token"`
	PinCode        string `json:"pin_code"`
	PrivateKey     string `json:"private_key"`
	IdentityNumber string `json:"identity_number"`

	FullName  string `json:"full_name"`
	AvatarURL string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
}
