package network

import "time"

func NetworkSnapshots(asset string, offset time.Time, asc bool, limit int, usedId, sessionId, privateKey string) ([]byte, error) {
	method := "GET"
	uri := "/network/snapshots"

	order := "DESC"
	if asc {
		order = "ASC"
	}
	params := P{
		"limit":  limit,
		"offset": offset.UTC().Format(time.RFC3339Nano),
		"asset":  asset,
		"order":  order,
	}

	return MixinRequest(method, uri, params, usedId, sessionId, privateKey)
}

func NetworkSnapshot(snapshotID string, usedId, sessionId, privateKey string) ([]byte, error) {
	method := "GET"
	uri := "/network/snapshots/" + snapshotID
	return MixinRequest(method, uri, nil, usedId, sessionId, privateKey)
}

func ExternalTransactions(asset, publicKeyOrTag, emptyOrName string, offset time.Time, limit int, usedId, sessionId, privateKey string) ([]byte, error) {
	method := "GET"
	uri := "/external/transactions"

	params := P{
		"asset":  asset,
		"limit":  limit,
		"offset": offset.UTC().Format(time.RFC3339Nano),
	}
	if len(emptyOrName) == 0 {
		params["public_key"] = publicKeyOrTag
	} else {
		params["account_tag"] = publicKeyOrTag
		params["account_name"] = emptyOrName
	}

	return MixinRequest(method, uri, params, usedId, sessionId, privateKey)
}
