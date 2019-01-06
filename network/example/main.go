package main

import (
	"log"

	mixin "github.com/MooooonStar/mixin-sdk/network"
	"github.com/hokaccha/go-prettyjson"
)

func main() {
	user := mixin.NewUser(UserId, SessionId, PinToken, PrivateKey)
	bt, err := user.ReadProfile()
	if err != nil {
		log.Fatal(err)
	}
	v, _ := prettyjson.Format(bt)
	log.Println("profile", string(v))
}
