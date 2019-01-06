package main

import (
	"log"

	mixin "github.com/MooooonStar/mixin-sdk/network"
	"github.com/hokaccha/go-prettyjson"
)

const (
	UserId       = "c7ff704e-1a74-4f12-b05c-7a2be955a782"
	ClientSecret = "4d0b97de5edb6034f2ab9da31356aaab68f279fd39b4dc1739c9a30406a9775c"
	PinCode      = "944855"
	SessionId    = "001419ba-7316-421d-ba62-8a1113d19672"
	PinToken     = "D5VJmO+K6D5kOmlEK9I9M7/2/7dZvgIxq0HcDqHxfl+8IWvvPrcvh8D+XQ0XfloUy/rsPlnYNuJyzT/cduPKAsXnz1DkIqnpDu5NrHA9jZyIr/9iWgkt+z2kDipDA8PSYmOoWKW5QV6fpor96bsl44T2fNEbSoFJG5rxRbHQtJo="
	PrivateKey   = `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCMxJ3yPNOaRsvHgz0x9FtyqZ0hV6Vmhrn+sWYypK13b0yeV6W7
WEDVw8NuSm87dHN3LMZZxIDR84t6D2LOXNdm3EhYmP890PpVWvFyg8X7D1lMefrs
YEubhxoe4WbjXHUU1fmvlUfvdquS59/zLuGQvZqTkrwa91TOQYbe445hhwIDAQAB
AoGAGiuTka1tSYlP6U+k2NytA6w04jYBMgZqHcetUEz9Uu8GN4nj7eiCZTt34dFE
zLDhpo5UcevuZxn4HEEwBV2NTfNKrM0qmLp51rwcBJtzF2lLNvANQCWrr8HejWe1
Q7ArIWh7gxaIOmFIP10Vkn7QWELSCSaK2qh5tdqfre8i0IECQQDRPrFJYGxfti5k
2BnA9IqykeGPF30/6Smx7HMktT7Aas6vB4UUN5FXAx+MrWMybJ5hnBc6pJC+1MIU
vfRvHjfpAkEArDjie1YUZ0uFzopJpOeRDDWC4rtstj1Z2OlxHEfX7kK4Cm98Uhjr
yoyzyl5xTT+DhaHI+w3vARtlrOuy15hX7wJBALpRTKO9zEJdgmohUq1SEr52z5YO
oGRsRcg8dzrUeI/1ixynYYRjBnOoQEuPiKi5tz3LM5PwPULvR/IYQrM/ASECQA1W
ypPq8uGdQ9vfchzHosBjVKPjCGSFE/RtAEnEdsEJgd+tCuAA9iJWC4bdEcF97d3n
zf1D8wMO8C0YhF2WexkCQCF04+yraHTuEQJOqYkPCCUyesb9tgXuLtKuVLeu0t2t
l2vF9qo3e4dVlWaKoGdzF/CVtOY0drWeMHyY0FKMgoo=
-----END RSA PRIVATE KEY-----`
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
