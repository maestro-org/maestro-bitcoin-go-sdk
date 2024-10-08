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

// AssetRuneService contains methods and other services that help with interacting
// with the Maestro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetRuneService] method instead.
type AssetRuneService struct {
	Options []option.RequestOption
	Holders *AssetRuneHolderService
	Utxos   *AssetRuneUtxoService
}

// NewAssetRuneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssetRuneService(opts ...option.RequestOption) (r *AssetRuneService) {
	r = &AssetRuneService{}
	r.Options = opts
	r.Holders = NewAssetRuneHolderService(opts...)
	r.Utxos = NewAssetRuneUtxoService(opts...)
	return
}

// Information about the specified Rune, including etching, current supply and
// number of holders.
func (r *AssetRuneService) Get(ctx context.Context, rune string, opts ...option.RequestOption) (res *TimestampedRuneInfo, err error) {
	opts = append(r.Options[:], opts...)
	if rune == "" {
		err = errors.New("missing required rune parameter")
		return
	}
	path := fmt.Sprintf("assets/runes/%s", rune)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List of ID and names of all deployed Rune assets.
func (r *AssetRuneService) List(ctx context.Context, query AssetRuneListParams, opts ...option.RequestOption) (res *PaginatedRuneIDAndName, err error) {
	opts = append(r.Options[:], opts...)
	path := "assets/runes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PaginatedRuneIDAndName struct {
	Data        []PaginatedRuneIDAndNameData      `json:"data,required"`
	LastUpdated PaginatedRuneIDAndNameLastUpdated `json:"last_updated,required"`
	NextCursor  string                            `json:"next_cursor,nullable"`
	JSON        paginatedRuneIDAndNameJSON        `json:"-"`
}

// paginatedRuneIDAndNameJSON contains the JSON metadata for the struct
// [PaginatedRuneIDAndName]
type paginatedRuneIDAndNameJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedRuneIDAndName) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedRuneIDAndNameJSON) RawJSON() string {
	return r.raw
}

type PaginatedRuneIDAndNameData struct {
	ID         string                         `json:"id,required"`
	SpacedName string                         `json:"spaced_name,required"`
	JSON       paginatedRuneIDAndNameDataJSON `json:"-"`
}

// paginatedRuneIDAndNameDataJSON contains the JSON metadata for the struct
// [PaginatedRuneIDAndNameData]
type paginatedRuneIDAndNameDataJSON struct {
	ID          apijson.Field
	SpacedName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedRuneIDAndNameData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedRuneIDAndNameDataJSON) RawJSON() string {
	return r.raw
}

type PaginatedRuneIDAndNameLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                                 `json:"block_height,required"`
	JSON        paginatedRuneIDAndNameLastUpdatedJSON `json:"-"`
}

// paginatedRuneIDAndNameLastUpdatedJSON contains the JSON metadata for the struct
// [PaginatedRuneIDAndNameLastUpdated]
type paginatedRuneIDAndNameLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedRuneIDAndNameLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedRuneIDAndNameLastUpdatedJSON) RawJSON() string {
	return r.raw
}

type TimestampedRuneInfo struct {
	Data        TimestampedRuneInfoData        `json:"data,required"`
	LastUpdated TimestampedRuneInfoLastUpdated `json:"last_updated,required"`
	JSON        timestampedRuneInfoJSON        `json:"-"`
}

// timestampedRuneInfoJSON contains the JSON metadata for the struct
// [TimestampedRuneInfo]
type timestampedRuneInfoJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedRuneInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedRuneInfoJSON) RawJSON() string {
	return r.raw
}

type TimestampedRuneInfoData struct {
	ID                string                       `json:"id,required"`
	CirculatingSupply string                       `json:"circulating_supply,required"`
	EtchingCenotaph   bool                         `json:"etching_cenotaph,required"`
	EtchingHeight     int64                        `json:"etching_height,required"`
	EtchingTx         string                       `json:"etching_tx,required"`
	MaxSupply         string                       `json:"max_supply,required"`
	Mints             int64                        `json:"mints,required"`
	Name              string                       `json:"name,required"`
	SpacedName        string                       `json:"spaced_name,required"`
	Terms             TimestampedRuneInfoDataTerms `json:"terms,required"`
	TotalUtxos        int64                        `json:"total_utxos,required"`
	UniqueHolders     int64                        `json:"unique_holders,required"`
	Divisibility      int64                        `json:"divisibility,nullable"`
	Premine           string                       `json:"premine,nullable"`
	Symbol            string                       `json:"symbol,nullable"`
	JSON              timestampedRuneInfoDataJSON  `json:"-"`
}

// timestampedRuneInfoDataJSON contains the JSON metadata for the struct
// [TimestampedRuneInfoData]
type timestampedRuneInfoDataJSON struct {
	ID                apijson.Field
	CirculatingSupply apijson.Field
	EtchingCenotaph   apijson.Field
	EtchingHeight     apijson.Field
	EtchingTx         apijson.Field
	MaxSupply         apijson.Field
	Mints             apijson.Field
	Name              apijson.Field
	SpacedName        apijson.Field
	Terms             apijson.Field
	TotalUtxos        apijson.Field
	UniqueHolders     apijson.Field
	Divisibility      apijson.Field
	Premine           apijson.Field
	Symbol            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TimestampedRuneInfoData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedRuneInfoDataJSON) RawJSON() string {
	return r.raw
}

type TimestampedRuneInfoDataTerms struct {
	AmountPerMint string                           `json:"amount_per_mint,nullable"`
	EndHeight     string                           `json:"end_height,nullable"`
	EndOffset     string                           `json:"end_offset,nullable"`
	MintTxsCap    string                           `json:"mint_txs_cap,nullable"`
	StartHeight   string                           `json:"start_height,nullable"`
	StartOffset   string                           `json:"start_offset,nullable"`
	JSON          timestampedRuneInfoDataTermsJSON `json:"-"`
}

// timestampedRuneInfoDataTermsJSON contains the JSON metadata for the struct
// [TimestampedRuneInfoDataTerms]
type timestampedRuneInfoDataTermsJSON struct {
	AmountPerMint apijson.Field
	EndHeight     apijson.Field
	EndOffset     apijson.Field
	MintTxsCap    apijson.Field
	StartHeight   apijson.Field
	StartOffset   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TimestampedRuneInfoDataTerms) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedRuneInfoDataTermsJSON) RawJSON() string {
	return r.raw
}

type TimestampedRuneInfoLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                              `json:"block_height,required"`
	JSON        timestampedRuneInfoLastUpdatedJSON `json:"-"`
}

// timestampedRuneInfoLastUpdatedJSON contains the JSON metadata for the struct
// [TimestampedRuneInfoLastUpdated]
type timestampedRuneInfoLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedRuneInfoLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedRuneInfoLastUpdatedJSON) RawJSON() string {
	return r.raw
}

type AssetRuneListParams struct {
	// The max number of results per page
	Count param.Field[int64] `query:"count"`
	// Pagination cursor string, use the cursor included in a page of results to fetch
	// the next page
	Cursor param.Field[string] `query:"cursor"`
}

// URLQuery serializes [AssetRuneListParams]'s query parameters as `url.Values`.
func (r AssetRuneListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
