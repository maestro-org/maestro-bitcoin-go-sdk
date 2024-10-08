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

// AssetRuneHolderService contains methods and other services that help with
// interacting with the maestro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetRuneHolderService] method instead.
type AssetRuneHolderService struct {
	Options []option.RequestOption
}

// NewAssetRuneHolderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssetRuneHolderService(opts ...option.RequestOption) (r *AssetRuneHolderService) {
	r = &AssetRuneHolderService{}
	r.Options = opts
	return
}

// List of all addresses that hold the specified Rune, together with the respective
// amount of that rune.
func (r *AssetRuneHolderService) List(ctx context.Context, rune string, query AssetRuneHolderListParams, opts ...option.RequestOption) (res *PaginatedRuneHolder, err error) {
	opts = append(r.Options[:], opts...)
	if rune == "" {
		err = errors.New("missing required rune parameter")
		return
	}
	path := fmt.Sprintf("assets/runes/%s/holders", rune)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PaginatedRuneHolder struct {
	Data        []PaginatedRuneHolderData      `json:"data,required"`
	LastUpdated PaginatedRuneHolderLastUpdated `json:"last_updated,required"`
	NextCursor  string                         `json:"next_cursor,nullable"`
	JSON        paginatedRuneHolderJSON        `json:"-"`
}

// paginatedRuneHolderJSON contains the JSON metadata for the struct
// [PaginatedRuneHolder]
type paginatedRuneHolderJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedRuneHolder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedRuneHolderJSON) RawJSON() string {
	return r.raw
}

type PaginatedRuneHolderData struct {
	Balance      string                      `json:"balance,required"`
	ScriptPubkey string                      `json:"script_pubkey,required"`
	Address      string                      `json:"address,nullable"`
	JSON         paginatedRuneHolderDataJSON `json:"-"`
}

// paginatedRuneHolderDataJSON contains the JSON metadata for the struct
// [PaginatedRuneHolderData]
type paginatedRuneHolderDataJSON struct {
	Balance      apijson.Field
	ScriptPubkey apijson.Field
	Address      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PaginatedRuneHolderData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedRuneHolderDataJSON) RawJSON() string {
	return r.raw
}

type PaginatedRuneHolderLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                              `json:"block_height,required"`
	JSON        paginatedRuneHolderLastUpdatedJSON `json:"-"`
}

// paginatedRuneHolderLastUpdatedJSON contains the JSON metadata for the struct
// [PaginatedRuneHolderLastUpdated]
type paginatedRuneHolderLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedRuneHolderLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedRuneHolderLastUpdatedJSON) RawJSON() string {
	return r.raw
}

type AssetRuneHolderListParams struct {
	// The max number of results per page
	Count param.Field[int64] `query:"count"`
	// Pagination cursor string, use the cursor included in a page of results to fetch
	// the next page
	Cursor param.Field[string] `query:"cursor"`
}

// URLQuery serializes [AssetRuneHolderListParams]'s query parameters as
// `url.Values`.
func (r AssetRuneHolderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
