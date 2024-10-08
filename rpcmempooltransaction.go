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

// RpcMempoolTransactionService contains methods and other services that help with
// interacting with the Maestro Dapp Platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRpcMempoolTransactionService] method instead.
type RpcMempoolTransactionService struct {
	Options     []option.RequestOption
	Ancestors   *RpcMempoolTransactionAncestorService
	Descendants *RpcMempoolTransactionDescendantService
}

// NewRpcMempoolTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRpcMempoolTransactionService(opts ...option.RequestOption) (r *RpcMempoolTransactionService) {
	r = &RpcMempoolTransactionService{}
	r.Options = opts
	r.Ancestors = NewRpcMempoolTransactionAncestorService(opts...)
	r.Descendants = NewRpcMempoolTransactionDescendantService(opts...)
	return
}

// Mempool Transaction Details
func (r *RpcMempoolTransactionService) Get(ctx context.Context, txHash string, opts ...option.RequestOption) (res *TimestampedMempoolTransactionDetails, err error) {
	opts = append(r.Options[:], opts...)
	if txHash == "" {
		err = errors.New("missing required tx_hash parameter")
		return
	}
	path := fmt.Sprintf("rpc/mempool/transactions/%s", txHash)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Mempool Transactions
func (r *RpcMempoolTransactionService) List(ctx context.Context, opts ...option.RequestOption) (res *TimestampedMempoolTransactions, err error) {
	opts = append(r.Options[:], opts...)
	path := "rpc/mempool/transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TimestampedMempoolTransactionDetails struct {
	Data        TimestampedMempoolTransactionDetailsData        `json:"data,required"`
	LastUpdated TimestampedMempoolTransactionDetailsLastUpdated `json:"last_updated,required"`
	JSON        timestampedMempoolTransactionDetailsJSON        `json:"-"`
}

// timestampedMempoolTransactionDetailsJSON contains the JSON metadata for the
// struct [TimestampedMempoolTransactionDetails]
type timestampedMempoolTransactionDetailsJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedMempoolTransactionDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedMempoolTransactionDetailsJSON) RawJSON() string {
	return r.raw
}

type TimestampedMempoolTransactionDetailsData struct {
	// Number of ancestors
	AncestorCount int64 `json:"ancestor_count,required"`
	// Size of ancestors
	AncestorSize int64 `json:"ancestor_size,required"`
	// Whether the transaction is BIP 125 replaceable
	Bip125Replaceable bool `json:"bip125_replaceable,required"`
	// Dependencies of the transaction
	Depends []string `json:"depends,required"`
	// Number of descendants
	DescendantCount int64 `json:"descendant_count,required"`
	// Size of descendants
	DescendantSize int64                                        `json:"descendant_size,required"`
	Fees           TimestampedMempoolTransactionDetailsDataFees `json:"fees,required"`
	// Block height
	Height int64 `json:"height,required"`
	// Transactions that spend this one
	SpentBy []string `json:"spent_by,required"`
	// Time of the transaction
	Time int64 `json:"time,required"`
	// Whether the transaction is unbroadcast
	Unbroadcast bool `json:"unbroadcast,required"`
	// Virtual size of the transaction
	Vsize int64 `json:"vsize,required"`
	// Weight of the transaction
	Weight int64 `json:"weight,required"`
	// Witness transaction ID
	WtxID string                                       `json:"wtx_id,required"`
	JSON  timestampedMempoolTransactionDetailsDataJSON `json:"-"`
}

// timestampedMempoolTransactionDetailsDataJSON contains the JSON metadata for the
// struct [TimestampedMempoolTransactionDetailsData]
type timestampedMempoolTransactionDetailsDataJSON struct {
	AncestorCount     apijson.Field
	AncestorSize      apijson.Field
	Bip125Replaceable apijson.Field
	Depends           apijson.Field
	DescendantCount   apijson.Field
	DescendantSize    apijson.Field
	Fees              apijson.Field
	Height            apijson.Field
	SpentBy           apijson.Field
	Time              apijson.Field
	Unbroadcast       apijson.Field
	Vsize             apijson.Field
	Weight            apijson.Field
	WtxID             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TimestampedMempoolTransactionDetailsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedMempoolTransactionDetailsDataJSON) RawJSON() string {
	return r.raw
}

type TimestampedMempoolTransactionDetailsDataFees struct {
	// Ancestor fee
	Ancestor string `json:"ancestor,required"`
	// Base fee
	Base string `json:"base,required"`
	// Descendant fee
	Descendant string `json:"descendant,required"`
	// Modified fee
	Modified string                                           `json:"modified,required"`
	JSON     timestampedMempoolTransactionDetailsDataFeesJSON `json:"-"`
}

// timestampedMempoolTransactionDetailsDataFeesJSON contains the JSON metadata for
// the struct [TimestampedMempoolTransactionDetailsDataFees]
type timestampedMempoolTransactionDetailsDataFeesJSON struct {
	Ancestor    apijson.Field
	Base        apijson.Field
	Descendant  apijson.Field
	Modified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedMempoolTransactionDetailsDataFees) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedMempoolTransactionDetailsDataFeesJSON) RawJSON() string {
	return r.raw
}

type TimestampedMempoolTransactionDetailsLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                                               `json:"block_height,required"`
	JSON        timestampedMempoolTransactionDetailsLastUpdatedJSON `json:"-"`
}

// timestampedMempoolTransactionDetailsLastUpdatedJSON contains the JSON metadata
// for the struct [TimestampedMempoolTransactionDetailsLastUpdated]
type timestampedMempoolTransactionDetailsLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedMempoolTransactionDetailsLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedMempoolTransactionDetailsLastUpdatedJSON) RawJSON() string {
	return r.raw
}

type TimestampedMempoolTransactions struct {
	Data        TimestampedMempoolTransactionsData        `json:"data,required"`
	LastUpdated TimestampedMempoolTransactionsLastUpdated `json:"last_updated,required"`
	JSON        timestampedMempoolTransactionsJSON        `json:"-"`
}

// timestampedMempoolTransactionsJSON contains the JSON metadata for the struct
// [TimestampedMempoolTransactions]
type timestampedMempoolTransactionsJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedMempoolTransactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedMempoolTransactionsJSON) RawJSON() string {
	return r.raw
}

type TimestampedMempoolTransactionsData struct {
	// Sequence number of the mempool
	MempoolSequence int64 `json:"mempool_sequence,required"`
	// List of transaction IDs
	TxIDs []string                               `json:"tx_ids,required"`
	JSON  timestampedMempoolTransactionsDataJSON `json:"-"`
}

// timestampedMempoolTransactionsDataJSON contains the JSON metadata for the struct
// [TimestampedMempoolTransactionsData]
type timestampedMempoolTransactionsDataJSON struct {
	MempoolSequence apijson.Field
	TxIDs           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TimestampedMempoolTransactionsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedMempoolTransactionsDataJSON) RawJSON() string {
	return r.raw
}

type TimestampedMempoolTransactionsLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                                         `json:"block_height,required"`
	JSON        timestampedMempoolTransactionsLastUpdatedJSON `json:"-"`
}

// timestampedMempoolTransactionsLastUpdatedJSON contains the JSON metadata for the
// struct [TimestampedMempoolTransactionsLastUpdated]
type timestampedMempoolTransactionsLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedMempoolTransactionsLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedMempoolTransactionsLastUpdatedJSON) RawJSON() string {
	return r.raw
}
