// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk

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

// AddressRuneService contains methods and other services that help with
// interacting with the Maestro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressRuneService] method instead.
type AddressRuneService struct {
	Options []option.RequestOption
}

// NewAddressRuneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAddressRuneService(opts ...option.RequestOption) (r *AddressRuneService) {
	r = &AddressRuneService{}
	r.Options = opts
	return
}

// Return all UTxOs controlled by the specified address or script pubkey which
// contain some of a specific kind of rune.
func (r *AddressRuneService) Get(ctx context.Context, address string, rune string, query AddressRuneGetParams, opts ...option.RequestOption) (res *PaginatedAddressRuneUtxo, err error) {
	opts = append(r.Options[:], opts...)
	if address == "" {
		err = errors.New("missing required address parameter")
		return
	}
	if rune == "" {
		err = errors.New("missing required rune parameter")
		return
	}
	path := fmt.Sprintf("addresses/%s/runes/%s", address, rune)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Map of all Runes tokens and corresponding amounts in UTxOs controlled by the
// specified address or script pubkey.
func (r *AddressRuneService) List(ctx context.Context, address string, opts ...option.RequestOption) (res *TimestampedRuneQuantities, err error) {
	opts = append(r.Options[:], opts...)
	if address == "" {
		err = errors.New("missing required address parameter")
		return
	}
	path := fmt.Sprintf("addresses/%s/runes", address)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PaginatedAddressRuneUtxo struct {
	Data        []PaginatedAddressRuneUtxoData      `json:"data,required"`
	LastUpdated PaginatedAddressRuneUtxoLastUpdated `json:"last_updated,required"`
	NextCursor  string                              `json:"next_cursor,nullable"`
	JSON        paginatedAddressRuneUtxoJSON        `json:"-"`
}

// paginatedAddressRuneUtxoJSON contains the JSON metadata for the struct
// [PaginatedAddressRuneUtxo]
type paginatedAddressRuneUtxoJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedAddressRuneUtxo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedAddressRuneUtxoJSON) RawJSON() string {
	return r.raw
}

type PaginatedAddressRuneUtxoData struct {
	Confirmations int64                            `json:"confirmations,required"`
	Height        int64                            `json:"height,required"`
	RuneAmount    string                           `json:"rune_amount,required"`
	Satoshis      string                           `json:"satoshis,required"`
	Txid          string                           `json:"txid,required"`
	Vout          int64                            `json:"vout,required"`
	JSON          paginatedAddressRuneUtxoDataJSON `json:"-"`
}

// paginatedAddressRuneUtxoDataJSON contains the JSON metadata for the struct
// [PaginatedAddressRuneUtxoData]
type paginatedAddressRuneUtxoDataJSON struct {
	Confirmations apijson.Field
	Height        apijson.Field
	RuneAmount    apijson.Field
	Satoshis      apijson.Field
	Txid          apijson.Field
	Vout          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PaginatedAddressRuneUtxoData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedAddressRuneUtxoDataJSON) RawJSON() string {
	return r.raw
}

type PaginatedAddressRuneUtxoLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                                   `json:"block_height,required"`
	JSON        paginatedAddressRuneUtxoLastUpdatedJSON `json:"-"`
}

// paginatedAddressRuneUtxoLastUpdatedJSON contains the JSON metadata for the
// struct [PaginatedAddressRuneUtxoLastUpdated]
type paginatedAddressRuneUtxoLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedAddressRuneUtxoLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedAddressRuneUtxoLastUpdatedJSON) RawJSON() string {
	return r.raw
}

type TimestampedRuneQuantities struct {
	Data        map[string]string                    `json:"data,required"`
	LastUpdated TimestampedRuneQuantitiesLastUpdated `json:"last_updated,required"`
	JSON        timestampedRuneQuantitiesJSON        `json:"-"`
}

// timestampedRuneQuantitiesJSON contains the JSON metadata for the struct
// [TimestampedRuneQuantities]
type timestampedRuneQuantitiesJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedRuneQuantities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedRuneQuantitiesJSON) RawJSON() string {
	return r.raw
}

type TimestampedRuneQuantitiesLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                                    `json:"block_height,required"`
	JSON        timestampedRuneQuantitiesLastUpdatedJSON `json:"-"`
}

// timestampedRuneQuantitiesLastUpdatedJSON contains the JSON metadata for the
// struct [TimestampedRuneQuantitiesLastUpdated]
type timestampedRuneQuantitiesLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedRuneQuantitiesLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedRuneQuantitiesLastUpdatedJSON) RawJSON() string {
	return r.raw
}

type AddressRuneGetParams struct {
	// The max number of results per page
	Count param.Field[int64] `query:"count"`
	// Pagination cursor string, use the cursor included in a page of results to fetch
	// the next page
	Cursor param.Field[string] `query:"cursor"`
	// Return only UTxOs created on or after a specific height
	From param.Field[int64] `query:"from"`
	// The order in which the results are sorted (by height at which UTxO was produced)
	Order param.Field[AddressRuneGetParamsOrder] `query:"order"`
	// Return only UTxOs created on or before a specific height
	To param.Field[int64] `query:"to"`
}

// URLQuery serializes [AddressRuneGetParams]'s query parameters as `url.Values`.
func (r AddressRuneGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The order in which the results are sorted (by height at which UTxO was produced)
type AddressRuneGetParamsOrder string

const (
	AddressRuneGetParamsOrderAsc  AddressRuneGetParamsOrder = "asc"
	AddressRuneGetParamsOrderDesc AddressRuneGetParamsOrder = "desc"
)

func (r AddressRuneGetParamsOrder) IsKnown() bool {
	switch r {
	case AddressRuneGetParamsOrderAsc, AddressRuneGetParamsOrderDesc:
		return true
	}
	return false
}
