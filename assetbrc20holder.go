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

// AssetBrc20HolderService contains methods and other services that help with
// interacting with the Maestro Dapp Platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetBrc20HolderService] method instead.
type AssetBrc20HolderService struct {
	Options []option.RequestOption
}

// NewAssetBrc20HolderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssetBrc20HolderService(opts ...option.RequestOption) (r *AssetBrc20HolderService) {
	r = &AssetBrc20HolderService{}
	r.Options = opts
	return
}

// Script pubkeys or addresses that hold the specified BRC20 asset and
// corresponding total balances.
func (r *AssetBrc20HolderService) List(ctx context.Context, ticker string, query AssetBrc20HolderListParams, opts ...option.RequestOption) (res *PaginatedBrc20Holder, err error) {
	opts = append(r.Options[:], opts...)
	if ticker == "" {
		err = errors.New("missing required ticker parameter")
		return
	}
	path := fmt.Sprintf("assets/brc20/%s/holders", ticker)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PaginatedBrc20Holder struct {
	Data        []PaginatedBrc20HolderData      `json:"data,required"`
	LastUpdated PaginatedBrc20HolderLastUpdated `json:"last_updated,required"`
	NextCursor  string                          `json:"next_cursor,nullable"`
	JSON        paginatedBrc20HolderJSON        `json:"-"`
}

// paginatedBrc20HolderJSON contains the JSON metadata for the struct
// [PaginatedBrc20Holder]
type paginatedBrc20HolderJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedBrc20Holder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedBrc20HolderJSON) RawJSON() string {
	return r.raw
}

type PaginatedBrc20HolderData struct {
	Balance      string                       `json:"balance,required"`
	ScriptPubkey string                       `json:"script_pubkey,required"`
	Address      string                       `json:"address,nullable"`
	JSON         paginatedBrc20HolderDataJSON `json:"-"`
}

// paginatedBrc20HolderDataJSON contains the JSON metadata for the struct
// [PaginatedBrc20HolderData]
type paginatedBrc20HolderDataJSON struct {
	Balance      apijson.Field
	ScriptPubkey apijson.Field
	Address      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PaginatedBrc20HolderData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedBrc20HolderDataJSON) RawJSON() string {
	return r.raw
}

type PaginatedBrc20HolderLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                               `json:"block_height,required"`
	JSON        paginatedBrc20HolderLastUpdatedJSON `json:"-"`
}

// paginatedBrc20HolderLastUpdatedJSON contains the JSON metadata for the struct
// [PaginatedBrc20HolderLastUpdated]
type paginatedBrc20HolderLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedBrc20HolderLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedBrc20HolderLastUpdatedJSON) RawJSON() string {
	return r.raw
}

type AssetBrc20HolderListParams struct {
	// The max number of results per page
	Count param.Field[int64] `query:"count"`
	// Pagination cursor string, use the cursor included in a page of results to fetch
	// the next page
	Cursor param.Field[string] `query:"cursor"`
}

// URLQuery serializes [AssetBrc20HolderListParams]'s query parameters as
// `url.Values`.
func (r AssetBrc20HolderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
