// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk

import (
	"context"
	"net/http"

	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/apijson"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/requestconfig"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// RpcMempoolInfoService contains methods and other services that help with
// interacting with the Maestro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRpcMempoolInfoService] method instead.
type RpcMempoolInfoService struct {
	Options []option.RequestOption
}

// NewRpcMempoolInfoService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRpcMempoolInfoService(opts ...option.RequestOption) (r *RpcMempoolInfoService) {
	r = &RpcMempoolInfoService{}
	r.Options = opts
	return
}

// Mempool Info
func (r *RpcMempoolInfoService) Get(ctx context.Context, opts ...option.RequestOption) (res *TimestampedMempoolInfo, err error) {
	opts = append(r.Options[:], opts...)
	path := "rpc/mempool/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TimestampedMempoolInfo struct {
	Data        TimestampedMempoolInfoData        `json:"data,required"`
	LastUpdated TimestampedMempoolInfoLastUpdated `json:"last_updated,required"`
	JSON        timestampedMempoolInfoJSON        `json:"-"`
}

// timestampedMempoolInfoJSON contains the JSON metadata for the struct
// [TimestampedMempoolInfo]
type timestampedMempoolInfoJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedMempoolInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedMempoolInfoJSON) RawJSON() string {
	return r.raw
}

type TimestampedMempoolInfoData struct {
	// Total memory usage for the mempool (in bytes)
	Bytes int64 `json:"bytes,required"`
	// Whether full replace-by-fee (RBF) is enabled
	FullRbf bool `json:"full_rbf,required"`
	// The incremental relay fee setting (in BTC)
	IncrementalRelayFee string `json:"incremental_relay_fee,required"`
	// Whether the mempool is fully loaded
	Loaded bool `json:"loaded,required"`
	// Maximum memory usage for the mempool (in bytes)
	MaxMempool int64 `json:"max_mempool,required"`
	// The minimum fee rate (in BTC/kB) for mempool transactions
	MempoolMinFee string `json:"mempool_min_fee,required"`
	// The minimum fee rate (in BTC/kB) for relaying transactions
	MinRelayTxFee string `json:"min_relay_tx_fee,required"`
	// Number of transactions in the mempool
	Size int64 `json:"size,required"`
	// The total fees (in BTC) in the mempool
	TotalFee string `json:"total_fee,required"`
	// Number of transactions that have not been broadcast
	UnbroadcastCount int64 `json:"unbroadcast_count,required"`
	// Total usage of the mempool (in bytes)
	Usage int64                          `json:"usage,required"`
	JSON  timestampedMempoolInfoDataJSON `json:"-"`
}

// timestampedMempoolInfoDataJSON contains the JSON metadata for the struct
// [TimestampedMempoolInfoData]
type timestampedMempoolInfoDataJSON struct {
	Bytes               apijson.Field
	FullRbf             apijson.Field
	IncrementalRelayFee apijson.Field
	Loaded              apijson.Field
	MaxMempool          apijson.Field
	MempoolMinFee       apijson.Field
	MinRelayTxFee       apijson.Field
	Size                apijson.Field
	TotalFee            apijson.Field
	UnbroadcastCount    apijson.Field
	Usage               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TimestampedMempoolInfoData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedMempoolInfoDataJSON) RawJSON() string {
	return r.raw
}

type TimestampedMempoolInfoLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                                 `json:"block_height,required"`
	JSON        timestampedMempoolInfoLastUpdatedJSON `json:"-"`
}

// timestampedMempoolInfoLastUpdatedJSON contains the JSON metadata for the struct
// [TimestampedMempoolInfoLastUpdated]
type timestampedMempoolInfoLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedMempoolInfoLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedMempoolInfoLastUpdatedJSON) RawJSON() string {
	return r.raw
}
