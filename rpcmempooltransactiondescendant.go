// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoin

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/apijson"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/requestconfig"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// RpcMempoolTransactionDescendantService contains methods and other services that
// help with interacting with the maestro-bitcoin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRpcMempoolTransactionDescendantService] method instead.
type RpcMempoolTransactionDescendantService struct {
	Options []option.RequestOption
}

// NewRpcMempoolTransactionDescendantService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRpcMempoolTransactionDescendantService(opts ...option.RequestOption) (r *RpcMempoolTransactionDescendantService) {
	r = &RpcMempoolTransactionDescendantService{}
	r.Options = opts
	return
}

// List Mempool Transaction Descendants
func (r *RpcMempoolTransactionDescendantService) List(ctx context.Context, txHash string, opts ...option.RequestOption) (res *TimestampedMempoolTransactionDescendants, err error) {
	opts = append(r.Options[:], opts...)
	if txHash == "" {
		err = errors.New("missing required tx_hash parameter")
		return
	}
	path := fmt.Sprintf("rpc/mempool/transactions/%s/descendants", txHash)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TimestampedMempoolTransactionDescendants struct {
	Data        []string                                            `json:"data,required"`
	LastUpdated TimestampedMempoolTransactionDescendantsLastUpdated `json:"last_updated,required"`
	JSON        timestampedMempoolTransactionDescendantsJSON        `json:"-"`
}

// timestampedMempoolTransactionDescendantsJSON contains the JSON metadata for the
// struct [TimestampedMempoolTransactionDescendants]
type timestampedMempoolTransactionDescendantsJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedMempoolTransactionDescendants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedMempoolTransactionDescendantsJSON) RawJSON() string {
	return r.raw
}

type TimestampedMempoolTransactionDescendantsLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                                                   `json:"block_height,required"`
	JSON        timestampedMempoolTransactionDescendantsLastUpdatedJSON `json:"-"`
}

// timestampedMempoolTransactionDescendantsLastUpdatedJSON contains the JSON
// metadata for the struct [TimestampedMempoolTransactionDescendantsLastUpdated]
type timestampedMempoolTransactionDescendantsLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedMempoolTransactionDescendantsLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedMempoolTransactionDescendantsLastUpdatedJSON) RawJSON() string {
	return r.raw
}
