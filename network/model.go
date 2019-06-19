package network

type User struct {
	UserId         string `json:"user_id"`
	SessionId      string `json:"session_id"`
	PrivateKey     string `json:"private_key"`
	PinCode        string `json:"pin_code"`
	PinToken       string `json:"pin_token"`
	IdentityNumber string `json:"identity_number"`
	FullName       string `json:"full_name"`
	AvatarURL      string `json:"avatar_url"`
	CreatedAt      string `json:"created_at"`
}

type Asset struct {
	AssetID  string `json:"asset_id"`
	AssetKey string `json:"asset_key"`
	ChainID  string `json:"chain_id"`
	IconUrl  string `json:"icon_url"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Type     string `json:"type"`
}


