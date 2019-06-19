package network

const (
	ETC   = "2204c1ee-0ea2-4add-bb9a-b3719cfff93a"
	XRP   = "23dfb5a5-5d7b-48b6-905f-3970e3176e27"
	XEM   = "27921032-f73e-434e-955f-43d55672ee31"
	ETH   = "43d61dcd-e413-450d-80b8-101d5e903357"
	DASH  = "6472e7e3-75fd-48b6-b1dc-28d294ee1476"
	DOGE  = "6770a1e5-6086-44d5-b60f-545f9d9e8ffd"
	EOS   = "6cfe566e-4aad-470b-8c9a-2fd35b49c68d"
	LTC   = "76c802a2-7c88-447f-a93e-c29c9e5dd9c8"
	SC    = "990c4c29-57e9-48f6-9819-7d986ea44985"
	ZEN   = "a2c5d22b-62a2-4c13-b3f0-013290dbac60"
	BTC   = "c6d0c728-2624-429b-8e0d-d9d19b6592fa"
	ZEC   = "c996abc9-d94e-4494-b1cf-2a3fd3ac5714"
	BCH   = "fd11b6e3-0b87-41f1-a41f-f0e9b49e5bf0"
	CNB   = "965e5c6e-434c-3fa9-b780-c50f43cd955c"
	XIN   = "c94ac88f-4671-3976-b60a-09064f1811e8"
	CANDY = "43b645fc-a52c-38a3-8d3b-705e7aaefa15"
	USDT  = "815b0b1a-2764-3736-8faa-42d694fa620a"
)

var cast = map[string]string{
	"ETC":   ETC,
	"XRP":   XRP,
	"XEM":   XEM,
	"ETH":   ETH,
	"DASH":  DASH,
	"DOGE":  DOGE,
	"EOS":   EOS,
	"LTC":   LTC,
	"SC":    SC,
	"ZEN":   ZEN,
	"BTC":   BTC,
	"ZEC":   ZEC,
	"BCH":   BCH,
	"CNB":   CNB,
	"XIN":   XIN,
	"CANDY": CANDY,
	"USDT":  USDT,
}

func GetAssetId(symbol string) string {
	return cast[symbol]
}

func Who(assetId string) string {
	for k, v := range cast {
		if v == assetId {
			return k
		}
	}
	return "Not found"
}
