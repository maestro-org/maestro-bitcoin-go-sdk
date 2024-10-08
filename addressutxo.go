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

// AddressUtxoService contains methods and other services that help with
// interacting with the maestro-bitcoin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressUtxoService] method instead.
type AddressUtxoService struct {
	Options []option.RequestOption
}

// NewAddressUtxoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAddressUtxoService(opts ...option.RequestOption) (r *AddressUtxoService) {
	r = &AddressUtxoService{}
	r.Options = opts
	return
}

// List of all UTxOs which reside at the specified address or script pubkey.
func (r *AddressUtxoService) List(ctx context.Context, address string, query AddressUtxoListParams, opts ...option.RequestOption) (res *PaginatedUtxo, err error) {
	opts = append(r.Options[:], opts...)
	if address == "" {
		err = errors.New("missing required address parameter")
		return
	}
	path := fmt.Sprintf("addresses/%s/utxos", address)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PaginatedUtxo struct {
	Data        []PaginatedUtxoData      `json:"data,required"`
	LastUpdated PaginatedUtxoLastUpdated `json:"last_updated,required"`
	NextCursor  string                   `json:"next_cursor,nullable"`
	JSON        paginatedUtxoJSON        `json:"-"`
}

// paginatedUtxoJSON contains the JSON metadata for the struct [PaginatedUtxo]
type paginatedUtxoJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedUtxo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedUtxoJSON) RawJSON() string {
	return r.raw
}

type PaginatedUtxoData struct {
	Confirmations int64                          `json:"confirmations,required"`
	Height        int64                          `json:"height,required"`
	Inscriptions  []PaginatedUtxoDataInscription `json:"inscriptions,required"`
	Runes         []PaginatedUtxoDataRune        `json:"runes,required"`
	Satoshis      string                         `json:"satoshis,required"`
	ScriptPubkey  string                         `json:"script_pubkey,required"`
	Txid          string                         `json:"txid,required"`
	Vout          int64                          `json:"vout,required"`
	Address       string                         `json:"address,nullable"`
	JSON          paginatedUtxoDataJSON          `json:"-"`
}

// paginatedUtxoDataJSON contains the JSON metadata for the struct
// [PaginatedUtxoData]
type paginatedUtxoDataJSON struct {
	Confirmations apijson.Field
	Height        apijson.Field
	Inscriptions  apijson.Field
	Runes         apijson.Field
	Satoshis      apijson.Field
	ScriptPubkey  apijson.Field
	Txid          apijson.Field
	Vout          apijson.Field
	Address       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PaginatedUtxoData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedUtxoDataJSON) RawJSON() string {
	return r.raw
}

type PaginatedUtxoDataInscription struct {
	InscriptionID string                           `json:"inscription_id,required"`
	Offset        int64                            `json:"offset,required"`
	JSON          paginatedUtxoDataInscriptionJSON `json:"-"`
}

// paginatedUtxoDataInscriptionJSON contains the JSON metadata for the struct
// [PaginatedUtxoDataInscription]
type paginatedUtxoDataInscriptionJSON struct {
	InscriptionID apijson.Field
	Offset        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PaginatedUtxoDataInscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedUtxoDataInscriptionJSON) RawJSON() string {
	return r.raw
}

type PaginatedUtxoDataRune struct {
	Amount string                    `json:"amount,required"`
	RuneID string                    `json:"rune_id,required"`
	JSON   paginatedUtxoDataRuneJSON `json:"-"`
}

// paginatedUtxoDataRuneJSON contains the JSON metadata for the struct
// [PaginatedUtxoDataRune]
type paginatedUtxoDataRuneJSON struct {
	Amount      apijson.Field
	RuneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedUtxoDataRune) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedUtxoDataRuneJSON) RawJSON() string {
	return r.raw
}

type PaginatedUtxoLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                        `json:"block_height,required"`
	JSON        paginatedUtxoLastUpdatedJSON `json:"-"`
}

// paginatedUtxoLastUpdatedJSON contains the JSON metadata for the struct
// [PaginatedUtxoLastUpdated]
type paginatedUtxoLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedUtxoLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedUtxoLastUpdatedJSON) RawJSON() string {
	return r.raw
}

type AddressUtxoListParams struct {
	// The max number of results per page
	Count param.Field[int64] `query:"count"`
	// Pagination cursor string, use the cursor included in a page of results to fetch
	// the next page
	Cursor param.Field[string] `query:"cursor"`
	// Exclude UTxOs involved in metaprotocols (currently only runes and inscriptions
	// will be discovered, more metaprotocols may be supported in future)
	ExcludeMetaprotocols param.Field[bool] `query:"exclude_metaprotocols"`
	// Ignore UTxOs containing less than 100000 sats
	FilterDust param.Field[bool] `query:"filter_dust"`
	// Ignore UTxOs containing less than specified satoshis
	FilterDustThreshold param.Field[int64] `query:"filter_dust_threshold"`
	// Return only UTxOs created on or after a specific height
	From param.Field[int64] `query:"from"`
	// The order in which the results are sorted (by height at which UTxO was produced)
	Order param.Field[AddressUtxoListParamsOrder] `query:"order"`
	// Return only UTxOs created on or before a specific height
	To param.Field[int64] `query:"to"`
}

// URLQuery serializes [AddressUtxoListParams]'s query parameters as `url.Values`.
func (r AddressUtxoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The order in which the results are sorted (by height at which UTxO was produced)
type AddressUtxoListParamsOrder string

const (
	AddressUtxoListParamsOrderAsc  AddressUtxoListParamsOrder = "asc"
	AddressUtxoListParamsOrderDesc AddressUtxoListParamsOrder = "desc"
)

func (r AddressUtxoListParamsOrder) IsKnown() bool {
	switch r {
	case AddressUtxoListParamsOrderAsc, AddressUtxoListParamsOrderDesc:
		return true
	}
	return false
}
