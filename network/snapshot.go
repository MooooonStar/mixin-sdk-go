package network

import (
	"time"
)

func NetworkSnapshots(asset string, offset time.Time, order string, limit int, usedId, sessionId, privateKey string) ([]byte, error) {
	params := P{
		"limit":  limit,
		"offset": offset.UTC().Format(time.RFC3339Nano),
		"asset":  asset,
		"order":  order,
	}
	return Request("GET", "/network/snapshots"+BuildQuery(params), nil, usedId, sessionId, privateKey)
}

func NetworkSnapshot(snapshotID string, usedId, sessionId, privateKey string) ([]byte, error) {
	return Request("GET", "/network/snapshots/"+snapshotID, nil, usedId, sessionId, privateKey)
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

	return Request("GET", "/external/transactions"+BuildQuery(params), nil, usedId, sessionId, privateKey)
}
func MyNetworkSnapshots(asset string, offset time.Time, limit int, usedId, sessionId, privateKey string) ([]byte, error) {
	params := P{
		"limit":  limit,
		"offset": offset.UTC().Format(time.RFC3339Nano),
		"asset":  asset,
	}

	return Request("GET", "/snapshots"+BuildQuery(params), nil, usedId, sessionId, privateKey)
}
