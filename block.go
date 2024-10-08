// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/apijson"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/requestconfig"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// BlockService contains methods and other services that help with interacting with
// the maestro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockService] method instead.
type BlockService struct {
	Options []option.RequestOption
	Latest  *BlockLatestService
}

// NewBlockService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBlockService(opts ...option.RequestOption) (r *BlockService) {
	r = &BlockService{}
	r.Options = opts
	r.Latest = NewBlockLatestService(opts...)
	return
}

// Provides detailed information about a specific block by hash.
func (r *BlockService) Get(ctx context.Context, blockHash string, opts ...option.RequestOption) (res *TimestampedBlock, err error) {
	opts = append(r.Options[:], opts...)
	if blockHash == "" {
		err = errors.New("missing required block_hash parameter")
		return
	}
	path := fmt.Sprintf("blocks/%s", blockHash)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TimestampedBlock struct {
	Data        TimestampedBlockData        `json:"data,required"`
	LastUpdated TimestampedBlockLastUpdated `json:"last_updated,required"`
	JSON        timestampedBlockJSON        `json:"-"`
}

// timestampedBlockJSON contains the JSON metadata for the struct
// [TimestampedBlock]
type timestampedBlockJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBlockJSON) RawJSON() string {
	return r.raw
}

type TimestampedBlockData struct {
	Header TimestampedBlockDataHeader `json:"header,required"`
	// List of transactions contained in the block
	Txdata []TimestampedBlockDataTxdata `json:"txdata,required"`
	JSON   timestampedBlockDataJSON     `json:"-"`
}

// timestampedBlockDataJSON contains the JSON metadata for the struct
// [TimestampedBlockData]
type timestampedBlockDataJSON struct {
	Header      apijson.Field
	Txdata      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBlockData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBlockDataJSON) RawJSON() string {
	return r.raw
}

type TimestampedBlockDataHeader struct {
	// The target value below which the blockhash must lie, encoded as a a float (with
	// well-defined rounding, of course)
	Bits int64 `json:"bits,required"`
	// The root hash of the merkle tree of transactions in the block
	MerkleRoot io.Reader `json:"merkle_root,required" format:"binary"`
	// The nonce, selected to obtain a low enough blockhash
	Nonce int64 `json:"nonce,required"`
	// Reference to the previous block in the chain
	PrevBlockhash io.Reader `json:"prev_blockhash,required" format:"binary"`
	// The timestamp of the block, as claimed by the miner
	Time int64 `json:"time,required"`
	// The protocol version. Should always be 1.
	Version int64                          `json:"version,required"`
	JSON    timestampedBlockDataHeaderJSON `json:"-"`
}

// timestampedBlockDataHeaderJSON contains the JSON metadata for the struct
// [TimestampedBlockDataHeader]
type timestampedBlockDataHeaderJSON struct {
	Bits          apijson.Field
	MerkleRoot    apijson.Field
	Nonce         apijson.Field
	PrevBlockhash apijson.Field
	Time          apijson.Field
	Version       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TimestampedBlockDataHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBlockDataHeaderJSON) RawJSON() string {
	return r.raw
}

type TimestampedBlockDataTxdata struct {
	// List of inputs
	Input []TimestampedBlockDataTxdataInput `json:"input,required"`
	// Block number before which this transaction is valid, or 0 for valid immediately.
	LockTime int64 `json:"lock_time,required"`
	// List of outputs
	Output []TimestampedBlockDataTxdataOutput `json:"output,required"`
	// The protocol version, is currently expected to be 1 or 2 (BIP 68).
	Version int64                          `json:"version,required"`
	JSON    timestampedBlockDataTxdataJSON `json:"-"`
}

// timestampedBlockDataTxdataJSON contains the JSON metadata for the struct
// [TimestampedBlockDataTxdata]
type timestampedBlockDataTxdataJSON struct {
	Input       apijson.Field
	LockTime    apijson.Field
	Output      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBlockDataTxdata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBlockDataTxdataJSON) RawJSON() string {
	return r.raw
}

type TimestampedBlockDataTxdataInput struct {
	PreviousOutput TimestampedBlockDataTxdataInputPreviousOutput `json:"previous_output,required"`
	ScriptSig      io.Reader                                     `json:"script_sig,required" format:"binary"`
	// The sequence number, which suggests to miners which of two conflicting
	// transactions should be preferred, or 0xFFFFFFFF to ignore this feature. This is
	// generally never used since the miner behaviour cannot be enforced.
	Sequence int64 `json:"sequence,required"`
	// Witness data: an array of byte-arrays. Note that this field is _not_
	// (de)serialized with the rest of the TxIn in Encodable/Decodable, as it is
	// (de)serialized at the end of the full Transaction. It _is_ (de)serialized with
	// the rest of the TxIn in other (de)serialization routines.
	Witness []io.Reader                         `json:"witness,required" format:"binary"`
	JSON    timestampedBlockDataTxdataInputJSON `json:"-"`
}

// timestampedBlockDataTxdataInputJSON contains the JSON metadata for the struct
// [TimestampedBlockDataTxdataInput]
type timestampedBlockDataTxdataInputJSON struct {
	PreviousOutput apijson.Field
	ScriptSig      apijson.Field
	Sequence       apijson.Field
	Witness        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TimestampedBlockDataTxdataInput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBlockDataTxdataInputJSON) RawJSON() string {
	return r.raw
}

type TimestampedBlockDataTxdataInputPreviousOutput struct {
	Txid string `json:"txid,required"`
	// The index of the referenced output in its transaction's vout
	Vout int64                                             `json:"vout,required"`
	JSON timestampedBlockDataTxdataInputPreviousOutputJSON `json:"-"`
}

// timestampedBlockDataTxdataInputPreviousOutputJSON contains the JSON metadata for
// the struct [TimestampedBlockDataTxdataInputPreviousOutput]
type timestampedBlockDataTxdataInputPreviousOutputJSON struct {
	Txid        apijson.Field
	Vout        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBlockDataTxdataInputPreviousOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBlockDataTxdataInputPreviousOutputJSON) RawJSON() string {
	return r.raw
}

type TimestampedBlockDataTxdataOutput struct {
	ScriptPubkey io.Reader `json:"script_pubkey,required" format:"binary"`
	// The value of the output, in satoshis
	Value int64                                `json:"value,required"`
	JSON  timestampedBlockDataTxdataOutputJSON `json:"-"`
}

// timestampedBlockDataTxdataOutputJSON contains the JSON metadata for the struct
// [TimestampedBlockDataTxdataOutput]
type timestampedBlockDataTxdataOutputJSON struct {
	ScriptPubkey apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TimestampedBlockDataTxdataOutput) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBlockDataTxdataOutputJSON) RawJSON() string {
	return r.raw
}

type TimestampedBlockLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                           `json:"block_height,required"`
	JSON        timestampedBlockLastUpdatedJSON `json:"-"`
}

// timestampedBlockLastUpdatedJSON contains the JSON metadata for the struct
// [TimestampedBlockLastUpdated]
type timestampedBlockLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBlockLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBlockLastUpdatedJSON) RawJSON() string {
	return r.raw
}
