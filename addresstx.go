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

// AddressTxService contains methods and other services that help with interacting
// with the Maestro Dapp Platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressTxService] method instead.
type AddressTxService struct {
	Options []option.RequestOption
}

// NewAddressTxService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAddressTxService(opts ...option.RequestOption) (r *AddressTxService) {
	r = &AddressTxService{}
	r.Options = opts
	return
}

// List of all transactions which consumed or produced a UTxO controlled by the
// specified address or script pubkey.
func (r *AddressTxService) List(ctx context.Context, address string, query AddressTxListParams, opts ...option.RequestOption) (res *PaginatedInvolvedTransaction, err error) {
	opts = append(r.Options[:], opts...)
	if address == "" {
		err = errors.New("missing required address parameter")
		return
	}
	path := fmt.Sprintf("addresses/%s/txs", address)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PaginatedInvolvedTransaction struct {
	Data        []PaginatedInvolvedTransactionData      `json:"data,required"`
	LastUpdated PaginatedInvolvedTransactionLastUpdated `json:"last_updated,required"`
	NextCursor  string                                  `json:"next_cursor,nullable"`
	JSON        paginatedInvolvedTransactionJSON        `json:"-"`
}

// paginatedInvolvedTransactionJSON contains the JSON metadata for the struct
// [PaginatedInvolvedTransaction]
type paginatedInvolvedTransactionJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedInvolvedTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedInvolvedTransactionJSON) RawJSON() string {
	return r.raw
}

type PaginatedInvolvedTransactionData struct {
	// Height of the block which included the transaction
	Height int64 `json:"height,required"`
	// Address/pubkey controlled an input UTxO
	Input bool `json:"input,required"`
	// Address/pubkey controlled an output UTxO
	Output bool `json:"output,required"`
	// The transaction's txid
	TxHash string                               `json:"tx_hash,required"`
	JSON   paginatedInvolvedTransactionDataJSON `json:"-"`
}

// paginatedInvolvedTransactionDataJSON contains the JSON metadata for the struct
// [PaginatedInvolvedTransactionData]
type paginatedInvolvedTransactionDataJSON struct {
	Height      apijson.Field
	Input       apijson.Field
	Output      apijson.Field
	TxHash      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedInvolvedTransactionData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedInvolvedTransactionDataJSON) RawJSON() string {
	return r.raw
}

type PaginatedInvolvedTransactionLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                                       `json:"block_height,required"`
	JSON        paginatedInvolvedTransactionLastUpdatedJSON `json:"-"`
}

// paginatedInvolvedTransactionLastUpdatedJSON contains the JSON metadata for the
// struct [PaginatedInvolvedTransactionLastUpdated]
type paginatedInvolvedTransactionLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedInvolvedTransactionLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedInvolvedTransactionLastUpdatedJSON) RawJSON() string {
	return r.raw
}

type AddressTxListParams struct {
	// The max number of results per page
	Count param.Field[int64] `query:"count"`
	// Pagination cursor string, use the cursor included in a page of results to fetch
	// the next page
	Cursor param.Field[string] `query:"cursor"`
	// Return only transactions included on or after a specific height
	From param.Field[int64] `query:"from"`
	// The order in which the results are sorted (by height at which transaction was
	// included in a block)
	Order param.Field[AddressTxListParamsOrder] `query:"order"`
	// Return only transactions included on or before a specific height
	To param.Field[int64] `query:"to"`
}

// URLQuery serializes [AddressTxListParams]'s query parameters as `url.Values`.
func (r AddressTxListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The order in which the results are sorted (by height at which transaction was
// included in a block)
type AddressTxListParamsOrder string

const (
	AddressTxListParamsOrderAsc  AddressTxListParamsOrder = "asc"
	AddressTxListParamsOrderDesc AddressTxListParamsOrder = "desc"
)

func (r AddressTxListParamsOrder) IsKnown() bool {
	switch r {
	case AddressTxListParamsOrderAsc, AddressTxListParamsOrderDesc:
		return true
	}
	return false
}
