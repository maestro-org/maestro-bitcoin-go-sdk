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

// AssetBrc20Service contains methods and other services that help with interacting
// with the maestro-bitcoin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetBrc20Service] method instead.
type AssetBrc20Service struct {
	Options []option.RequestOption
	Holders *AssetBrc20HolderService
}

// NewAssetBrc20Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssetBrc20Service(opts ...option.RequestOption) (r *AssetBrc20Service) {
	r = &AssetBrc20Service{}
	r.Options = opts
	r.Holders = NewAssetBrc20HolderService(opts...)
	return
}

// Information about the specified BRC20 asset.
func (r *AssetBrc20Service) Get(ctx context.Context, ticker string, opts ...option.RequestOption) (res *TimestampedBrc20Info, err error) {
	opts = append(r.Options[:], opts...)
	if ticker == "" {
		err = errors.New("missing required ticker parameter")
		return
	}
	path := fmt.Sprintf("assets/brc20/%s", ticker)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List of tickers of all deployed BRC20 assets.
func (r *AssetBrc20Service) List(ctx context.Context, query AssetBrc20ListParams, opts ...option.RequestOption) (res *PaginatedBrc20Ticker, err error) {
	opts = append(r.Options[:], opts...)
	path := "assets/brc20"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PaginatedBrc20Ticker struct {
	Data        []string                        `json:"data,required"`
	LastUpdated PaginatedBrc20TickerLastUpdated `json:"last_updated,required"`
	NextCursor  string                          `json:"next_cursor,nullable"`
	JSON        paginatedBrc20TickerJSON        `json:"-"`
}

// paginatedBrc20TickerJSON contains the JSON metadata for the struct
// [PaginatedBrc20Ticker]
type paginatedBrc20TickerJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedBrc20Ticker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedBrc20TickerJSON) RawJSON() string {
	return r.raw
}

type PaginatedBrc20TickerLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                               `json:"block_height,required"`
	JSON        paginatedBrc20TickerLastUpdatedJSON `json:"-"`
}

// paginatedBrc20TickerLastUpdatedJSON contains the JSON metadata for the struct
// [PaginatedBrc20TickerLastUpdated]
type paginatedBrc20TickerLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedBrc20TickerLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedBrc20TickerLastUpdatedJSON) RawJSON() string {
	return r.raw
}

type TimestampedBrc20Info struct {
	Data        TimestampedBrc20InfoData        `json:"data,required"`
	LastUpdated TimestampedBrc20InfoLastUpdated `json:"last_updated,required"`
	JSON        timestampedBrc20InfoJSON        `json:"-"`
}

// timestampedBrc20InfoJSON contains the JSON metadata for the struct
// [TimestampedBrc20Info]
type timestampedBrc20InfoJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBrc20Info) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBrc20InfoJSON) RawJSON() string {
	return r.raw
}

type TimestampedBrc20InfoData struct {
	DeployInscription string                        `json:"deploy_inscription,required"`
	Holders           int64                         `json:"holders,required"`
	MintedSupply      string                        `json:"minted_supply,required"`
	Terms             TimestampedBrc20InfoDataTerms `json:"terms,required"`
	Ticker            string                        `json:"ticker,required"`
	TickerHex         string                        `json:"ticker_hex,required"`
	JSON              timestampedBrc20InfoDataJSON  `json:"-"`
}

// timestampedBrc20InfoDataJSON contains the JSON metadata for the struct
// [TimestampedBrc20InfoData]
type timestampedBrc20InfoDataJSON struct {
	DeployInscription apijson.Field
	Holders           apijson.Field
	MintedSupply      apijson.Field
	Terms             apijson.Field
	Ticker            apijson.Field
	TickerHex         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TimestampedBrc20InfoData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBrc20InfoDataJSON) RawJSON() string {
	return r.raw
}

type TimestampedBrc20InfoDataTerms struct {
	Dec      int64                             `json:"dec,required"`
	Limit    string                            `json:"limit,required"`
	Max      string                            `json:"max,required"`
	SelfMint bool                              `json:"self_mint,required"`
	JSON     timestampedBrc20InfoDataTermsJSON `json:"-"`
}

// timestampedBrc20InfoDataTermsJSON contains the JSON metadata for the struct
// [TimestampedBrc20InfoDataTerms]
type timestampedBrc20InfoDataTermsJSON struct {
	Dec         apijson.Field
	Limit       apijson.Field
	Max         apijson.Field
	SelfMint    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBrc20InfoDataTerms) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBrc20InfoDataTermsJSON) RawJSON() string {
	return r.raw
}

type TimestampedBrc20InfoLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                               `json:"block_height,required"`
	JSON        timestampedBrc20InfoLastUpdatedJSON `json:"-"`
}

// timestampedBrc20InfoLastUpdatedJSON contains the JSON metadata for the struct
// [TimestampedBrc20InfoLastUpdated]
type timestampedBrc20InfoLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBrc20InfoLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBrc20InfoLastUpdatedJSON) RawJSON() string {
	return r.raw
}

type AssetBrc20ListParams struct {
	// The max number of results per page
	Count param.Field[int64] `query:"count"`
	// Pagination cursor string, use the cursor included in a page of results to fetch
	// the next page
	Cursor param.Field[string] `query:"cursor"`
}

// URLQuery serializes [AssetBrc20ListParams]'s query parameters as `url.Values`.
func (r AssetBrc20ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
