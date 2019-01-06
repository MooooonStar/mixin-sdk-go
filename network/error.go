package network

import "encoding/json"

type Error struct {
	Status      int    `json:"status"`
	Code        int    `json:"code"`
	Description string `json:"description"`
	trace       string
}

func (e Error) Error() string {
	bt, _ := json.Marshal(e)
	return string(bt)
}
