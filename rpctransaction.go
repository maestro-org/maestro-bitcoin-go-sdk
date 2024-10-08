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

// RpcTransactionService contains methods and other services that help with
// interacting with the maestro-bitcoin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRpcTransactionService] method instead.
type RpcTransactionService struct {
	Options []option.RequestOption
}

// NewRpcTransactionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRpcTransactionService(opts ...option.RequestOption) (r *RpcTransactionService) {
	r = &RpcTransactionService{}
	r.Options = opts
	return
}

// Transaction Details
func (r *RpcTransactionService) Get(ctx context.Context, txHash string, opts ...option.RequestOption) (res *TimestampedTransactionDetails, err error) {
	opts = append(r.Options[:], opts...)
	if txHash == "" {
		err = errors.New("missing required tx_hash parameter")
		return
	}
	path := fmt.Sprintf("rpc/transactions/%s", txHash)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TimestampedTransactionDetails struct {
	// Represents the details of a Bitcoin transaction.
	Data        TimestampedTransactionDetailsData        `json:"data,required"`
	LastUpdated TimestampedTransactionDetailsLastUpdated `json:"last_updated,required"`
	JSON        timestampedTransactionDetailsJSON        `json:"-"`
}

// timestampedTransactionDetailsJSON contains the JSON metadata for the struct
// [TimestampedTransactionDetails]
type timestampedTransactionDetailsJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedTransactionDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedTransactionDetailsJSON) RawJSON() string {
	return r.raw
}

// Represents the details of a Bitcoin transaction.
type TimestampedTransactionDetailsData struct {
	// The hash of the block that contains this transaction.
	BlockHash string `json:"block_hash,required"`
	// The time the block containing this transaction was mined (in UNIX timestamp).
	BlockTime int64 `json:"block_time,required"`
	// The number of confirmations for this transaction.
	Confirmations int64 `json:"confirmations,required"`
	// The transaction hash (TXID).
	Hash string `json:"hash,required"`
	// The raw transaction in hexadecimal form.
	Hex string `json:"hex,required"`
	// The locktime of the transaction, which specifies the earliest time or block when
	// the transaction can be included in a block.
	LockTime int64 `json:"lock_time,required"`
	// The size of the transaction in bytes.
	Size int64 `json:"size,required"`
	// The timestamp when the transaction was included in the block (in UNIX
	// timestamp).
	Time int64 `json:"time,required"`
	// The transaction ID (TXID) of this transaction.
	TxID string `json:"tx_id,required"`
	// The virtual size of the transaction, which accounts for SegWit (witness data)
	// scaling rules.
	VSize int64 `json:"v_size,required"`
	// The version number of the transaction, indicating the format or type of the
	// transaction.
	Version int64 `json:"version,required"`
	// The transaction inputs, which refer to the previous transaction outputs that are
	// being spent.
	Vin []TimestampedTransactionDetailsDataVin `json:"vin,required"`
	// The transaction outputs, which specify how much Bitcoin is being sent to which
	// addresses.
	Vout []TimestampedTransactionDetailsDataVout `json:"vout,required"`
	// The weight of the transaction as defined by SegWit rules (v_size \* 4).
	Weight int64                                 `json:"weight,required"`
	JSON   timestampedTransactionDetailsDataJSON `json:"-"`
}

// timestampedTransactionDetailsDataJSON contains the JSON metadata for the struct
// [TimestampedTransactionDetailsData]
type timestampedTransactionDetailsDataJSON struct {
	BlockHash     apijson.Field
	BlockTime     apijson.Field
	Confirmations apijson.Field
	Hash          apijson.Field
	Hex           apijson.Field
	LockTime      apijson.Field
	Size          apijson.Field
	Time          apijson.Field
	TxID          apijson.Field
	VSize         apijson.Field
	Version       apijson.Field
	Vin           apijson.Field
	Vout          apijson.Field
	Weight        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TimestampedTransactionDetailsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedTransactionDetailsDataJSON) RawJSON() string {
	return r.raw
}

// Represents the input of a Bitcoin transaction.
type TimestampedTransactionDetailsDataVin struct {
	// The sequence number for the input, which is used for BIP-125 replace-by-fee
	// (RBF).
	Sequence int64 `json:"sequence,required"`
	// The coinbase transaction data (for block rewards) if this is a coinbase
	// transaction.
	Coinbase string `json:"coinbase,nullable"`
	// The witness data (SegWit) for the transaction input.
	TxinWitness []string                                 `json:"txin_witness,nullable"`
	JSON        timestampedTransactionDetailsDataVinJSON `json:"-"`
}

// timestampedTransactionDetailsDataVinJSON contains the JSON metadata for the
// struct [TimestampedTransactionDetailsDataVin]
type timestampedTransactionDetailsDataVinJSON struct {
	Sequence    apijson.Field
	Coinbase    apijson.Field
	TxinWitness apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedTransactionDetailsDataVin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedTransactionDetailsDataVinJSON) RawJSON() string {
	return r.raw
}

// Represents the output of a Bitcoin transaction.
type TimestampedTransactionDetailsDataVout struct {
	// The index of this output in the transaction (0-based).
	N int64 `json:"n,required"`
	// Represents the script used to lock/unlock Bitcoin in an output.
	ScriptPubKey TimestampedTransactionDetailsDataVoutScriptPubKey `json:"script_pub_key,required"`
	// The amount of Bitcoin (in BTC) sent to this output.
	Value float64                                   `json:"value,required"`
	JSON  timestampedTransactionDetailsDataVoutJSON `json:"-"`
}

// timestampedTransactionDetailsDataVoutJSON contains the JSON metadata for the
// struct [TimestampedTransactionDetailsDataVout]
type timestampedTransactionDetailsDataVoutJSON struct {
	N            apijson.Field
	ScriptPubKey apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TimestampedTransactionDetailsDataVout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedTransactionDetailsDataVoutJSON) RawJSON() string {
	return r.raw
}

// Represents the script used to lock/unlock Bitcoin in an output.
type TimestampedTransactionDetailsDataVoutScriptPubKey struct {
	// The assembly (ASM) representation of the script, showing the operation codes and
	// operands.
	Asm string `json:"asm,required"`
	// The detailed descriptor of the scriptPubKey, providing more information about
	// the address and script.
	Desc string `json:"desc,required"`
	// The raw hexadecimal representation of the script.
	Hex string `json:"hex,required"`
	// The type of the script (e.g., "scripthash", "pubkeyhash").
	Type string `json:"type,required"`
	// The Bitcoin address to which this output belongs.
	Address string                                                `json:"address,nullable"`
	JSON    timestampedTransactionDetailsDataVoutScriptPubKeyJSON `json:"-"`
}

// timestampedTransactionDetailsDataVoutScriptPubKeyJSON contains the JSON metadata
// for the struct [TimestampedTransactionDetailsDataVoutScriptPubKey]
type timestampedTransactionDetailsDataVoutScriptPubKeyJSON struct {
	Asm         apijson.Field
	Desc        apijson.Field
	Hex         apijson.Field
	Type        apijson.Field
	Address     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedTransactionDetailsDataVoutScriptPubKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedTransactionDetailsDataVoutScriptPubKeyJSON) RawJSON() string {
	return r.raw
}

type TimestampedTransactionDetailsLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                                        `json:"block_height,required"`
	JSON        timestampedTransactionDetailsLastUpdatedJSON `json:"-"`
}

// timestampedTransactionDetailsLastUpdatedJSON contains the JSON metadata for the
// struct [TimestampedTransactionDetailsLastUpdated]
type timestampedTransactionDetailsLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedTransactionDetailsLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedTransactionDetailsLastUpdatedJSON) RawJSON() string {
	return r.raw
}
