package network

import "time"

func NetworkSnapshots(asset string, offset time.Time, order string, limit int, usedId, sessionId, privateKey string) ([]byte, error) {
	params := P{
		"limit":  limit,
		"offset": offset.UTC().Format(time.RFC3339Nano),
		"asset":  asset,
		"order":  order,
	}
	return MixinRequest("GET", "/network/snapshots", params, usedId, sessionId, privateKey)
}

func NetworkSnapshot(snapshotID string, usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/network/snapshots/"+snapshotID, nil, usedId, sessionId, privateKey)
}

func ExternalTransactions(assetID, publicOrName, emptyOrTag string, offset time.Time, limit int, usedId, sessionId, privateKey string) ([]byte, error) {
	params := P{
		"asset":  assetID,
		"limit":  limit,
		"offset": offset.UTC().Format(time.RFC3339Nano),
	}
	if len(emptyOrTag) == 0 {
		params["public_key"] = publicOrName
	} else {
		params["account_name"] = publicOrName
		params["account_tag"] = emptyOrTag
	}

	return MixinRequest("GET", "/external/transactions", params, usedId, sessionId, privateKey)
}
func MyNetworkSnapshots(asset string, offset time.Time, limit int, usedId, sessionId, privateKey string) ([]byte, error) {
	params := P{
		"limit":  limit,
		"offset": offset.UTC().Format(time.RFC3339Nano),
		"asset":  asset,
	}

	return MixinRequest("GET", "/snapshots", params, usedId, sessionId, privateKey)
}
