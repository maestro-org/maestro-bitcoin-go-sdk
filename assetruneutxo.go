// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoin

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/apijson"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/apiquery"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/param"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/requestconfig"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// AssetRuneUtxoService contains methods and other services that help with
// interacting with the maestro-bitcoin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetRuneUtxoService] method instead.
type AssetRuneUtxoService struct {
	Options []option.RequestOption
}

// NewAssetRuneUtxoService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssetRuneUtxoService(opts ...option.RequestOption) (r *AssetRuneUtxoService) {
	r = &AssetRuneUtxoService{}
	r.Options = opts
	return
}

// List of all UTxOs which contain the specified Rune.
func (r *AssetRuneUtxoService) List(ctx context.Context, rune string, query AssetRuneUtxoListParams, opts ...option.RequestOption) (res *PaginatedRuneUtxo, err error) {
	opts = append(r.Options[:], opts...)
	if rune == "" {
		err = errors.New("missing required rune parameter")
		return
	}
	path := fmt.Sprintf("assets/runes/%s/utxos", rune)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PaginatedRuneUtxo struct {
	Data        []PaginatedRuneUtxoData      `json:"data,required"`
	LastUpdated PaginatedRuneUtxoLastUpdated `json:"last_updated,required"`
	NextCursor  string                       `json:"next_cursor,nullable"`
	JSON        paginatedRuneUtxoJSON        `json:"-"`
}

// paginatedRuneUtxoJSON contains the JSON metadata for the struct
// [PaginatedRuneUtxo]
type paginatedRuneUtxoJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedRuneUtxo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedRuneUtxoJSON) RawJSON() string {
	return r.raw
}

type PaginatedRuneUtxoData struct {
	Confirmations int64                     `json:"confirmations,required"`
	Height        int64                     `json:"height,required"`
	RuneAmount    string                    `json:"rune_amount,required"`
	Satoshis      string                    `json:"satoshis,required"`
	ScriptPubkey  string                    `json:"script_pubkey,required"`
	Txid          string                    `json:"txid,required"`
	Vout          int64                     `json:"vout,required"`
	Address       string                    `json:"address,nullable"`
	JSON          paginatedRuneUtxoDataJSON `json:"-"`
}

// paginatedRuneUtxoDataJSON contains the JSON metadata for the struct
// [PaginatedRuneUtxoData]
type paginatedRuneUtxoDataJSON struct {
	Confirmations apijson.Field
	Height        apijson.Field
	RuneAmount    apijson.Field
	Satoshis      apijson.Field
	ScriptPubkey  apijson.Field
	Txid          apijson.Field
	Vout          apijson.Field
	Address       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PaginatedRuneUtxoData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedRuneUtxoDataJSON) RawJSON() string {
	return r.raw
}

type PaginatedRuneUtxoLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                            `json:"block_height,required"`
	JSON        paginatedRuneUtxoLastUpdatedJSON `json:"-"`
}

// paginatedRuneUtxoLastUpdatedJSON contains the JSON metadata for the struct
// [PaginatedRuneUtxoLastUpdated]
type paginatedRuneUtxoLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedRuneUtxoLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedRuneUtxoLastUpdatedJSON) RawJSON() string {
	return r.raw
}

type AssetRuneUtxoListParams struct {
	// The max number of results per page
	Count param.Field[int64] `query:"count"`
	// Pagination cursor string, use the cursor included in a page of results to fetch
	// the next page
	Cursor param.Field[string] `query:"cursor"`
	// Return only UTxOs created on or after a specific height
	From param.Field[int64] `query:"from"`
	// The order in which the results are sorted (by height at which UTxO was produced)
	Order param.Field[AssetRuneUtxoListParamsOrder] `query:"order"`
	// Return only UTxOs created on or before a specific height
	To param.Field[int64] `query:"to"`
}

// URLQuery serializes [AssetRuneUtxoListParams]'s query parameters as
// `url.Values`.
func (r AssetRuneUtxoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The order in which the results are sorted (by height at which UTxO was produced)
type AssetRuneUtxoListParamsOrder string

const (
	AssetRuneUtxoListParamsOrderAsc  AssetRuneUtxoListParamsOrder = "asc"
	AssetRuneUtxoListParamsOrderDesc AssetRuneUtxoListParamsOrder = "desc"
)

func (r AssetRuneUtxoListParamsOrder) IsKnown() bool {
	switch r {
	case AssetRuneUtxoListParamsOrderAsc, AssetRuneUtxoListParamsOrderDesc:
		return true
	}
	return false
}
