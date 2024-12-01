package entry // import "github.com/stateless-minds/go-ipfs-log/entry"

import (
	"context"

	"github.com/ipfs/go-cid"
	coreiface "github.com/ipfs/kubo/core/coreiface"

	"github.com/stateless-minds/go-ipfs-log/iface"
)

type FetchOptions = iface.FetchOptions

// FetchParallel has the same comportement than FetchAll, we keep it for retrop
// compatibility purpose
func FetchParallel(ctx context.Context, ipfs coreiface.CoreAPI, hashes []cid.Cid, options *FetchOptions) []iface.IPFSLogEntry {
	fetcher := NewFetcher(ipfs, options)
	return fetcher.Fetch(ctx, hashes)
}

// FetchAll gets entries from their CIDs.
func FetchAll(ctx context.Context, ipfs coreiface.CoreAPI, hashes []cid.Cid, options *FetchOptions) []iface.IPFSLogEntry {
	fetcher := NewFetcher(ipfs, options)
	return fetcher.Fetch(ctx, hashes)
}
