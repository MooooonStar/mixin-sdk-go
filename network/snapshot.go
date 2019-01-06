package network

import (
	"time"
)

func NetworkSnapshots(symbol, offset, limit, order string) ([]byte, error) {
	method := "GET"
	uri := "/network/snapshots"
	query := make(P, 0)

	count := "500"
	if len(limit) > 0 {
		count = limit
	}
	query["limit"] = count

	ts := time.Now().Add(-10 * time.Minute)
	start := ts.Format(time.RFC3339Nano)
	if len(offset) > 0 {
		start = offset
	}
	query["offset"] = start

	asset := ""
	if len(symbol) > 0 {
		asset = symbolAssetId[symbol]
	}
	query["asset"] = asset

	orderBy := "DESC"
	if len(order) > 0 {
		orderBy = order
	}
	query["order"] = orderBy

	return MixinRequest(method, uri, query)
}

func NetworkSnapshot(snapshot_id string) ([]byte, error) {
	method := "GET"
	uri := "/network/snapshots/" + snapshot_id
	return MixinRequest(method, uri)
}

func ExternalTransactions(symbol, public_key, limit, offset string, account_info ...P) ([]byte, error) {
	method := "GET"
	uri := "/external/transactions"
	query := make(P, 0)

	count := "500"
	if len(limit) > 0 {
		count = limit
	}
	query["limit"] = count

	ts := time.Now().Add(-10 * time.Minute)
	start := ts.Format(time.RFC3339Nano)
	if len(offset) > 0 {
		start = offset
	}
	query["offset"] = start

	asset := ""
	if len(symbol) > 0 {
		asset = symbolAssetId[symbol]
	}
	query["asset"] = asset

	if symbol == "EOS" {
		for k, v := range account_info[0] {
			query[k] = v
		}
	}

	return MixinRequest(method, uri, query)
}
