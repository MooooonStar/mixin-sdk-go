package main

import (
	"log"

	mixin "github.com/MooooonStar/mixin-sdk-go/network"
)

func main() {
	user := mixin.NewUser(UserId, SessionId, PrivateKey, PinCode, PinToken)
	//user := mixin.NewUser(UserId, SessionId, PrivateKey)

	profile, err := user.ReadProfile()
	if err != nil {
		log.Fatal("Read profile error", err)
	}
	log.Println("profile", string(profile))

	assets, err := user.ReadAssets()
	if err != nil {
		log.Fatal("Read assets error", assets)
	}
	log.Println("assets", string(assets))
}
