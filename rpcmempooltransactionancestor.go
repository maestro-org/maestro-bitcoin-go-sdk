// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/apijson"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/requestconfig"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// RpcMempoolTransactionAncestorService contains methods and other services that
// help with interacting with the Maestro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRpcMempoolTransactionAncestorService] method instead.
type RpcMempoolTransactionAncestorService struct {
	Options []option.RequestOption
}

// NewRpcMempoolTransactionAncestorService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRpcMempoolTransactionAncestorService(opts ...option.RequestOption) (r *RpcMempoolTransactionAncestorService) {
	r = &RpcMempoolTransactionAncestorService{}
	r.Options = opts
	return
}

// List Mempool Transaction Ancestors
func (r *RpcMempoolTransactionAncestorService) List(ctx context.Context, txHash string, opts ...option.RequestOption) (res *TimestampedMempoolTransactionAncestors, err error) {
	opts = append(r.Options[:], opts...)
	if txHash == "" {
		err = errors.New("missing required tx_hash parameter")
		return
	}
	path := fmt.Sprintf("rpc/mempool/transactions/%s/ancestors", txHash)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TimestampedMempoolTransactionAncestors struct {
	Data        []string                                          `json:"data,required"`
	LastUpdated TimestampedMempoolTransactionAncestorsLastUpdated `json:"last_updated,required"`
	JSON        timestampedMempoolTransactionAncestorsJSON        `json:"-"`
}

// timestampedMempoolTransactionAncestorsJSON contains the JSON metadata for the
// struct [TimestampedMempoolTransactionAncestors]
type timestampedMempoolTransactionAncestorsJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedMempoolTransactionAncestors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedMempoolTransactionAncestorsJSON) RawJSON() string {
	return r.raw
}

type TimestampedMempoolTransactionAncestorsLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                                                 `json:"block_height,required"`
	JSON        timestampedMempoolTransactionAncestorsLastUpdatedJSON `json:"-"`
}

// timestampedMempoolTransactionAncestorsLastUpdatedJSON contains the JSON metadata
// for the struct [TimestampedMempoolTransactionAncestorsLastUpdated]
type timestampedMempoolTransactionAncestorsLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedMempoolTransactionAncestorsLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedMempoolTransactionAncestorsLastUpdatedJSON) RawJSON() string {
	return r.raw
}
