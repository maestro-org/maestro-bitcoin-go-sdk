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

// AddressBrc20Service contains methods and other services that help with
// interacting with the maestro-bitcoin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressBrc20Service] method instead.
type AddressBrc20Service struct {
	Options []option.RequestOption
}

// NewAddressBrc20Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAddressBrc20Service(opts ...option.RequestOption) (r *AddressBrc20Service) {
	r = &AddressBrc20Service{}
	r.Options = opts
	return
}

// Map of all BRC20 assets and corresponding total and available balances
// controlled by the specified address or script pubkey.
func (r *AddressBrc20Service) List(ctx context.Context, address string, opts ...option.RequestOption) (res *TimestampedBrc20Quantities, err error) {
	opts = append(r.Options[:], opts...)
	if address == "" {
		err = errors.New("missing required address parameter")
		return
	}
	path := fmt.Sprintf("addresses/%s/brc20", address)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TimestampedBrc20Quantities struct {
	Data        map[string]string                     `json:"data,required"`
	LastUpdated TimestampedBrc20QuantitiesLastUpdated `json:"last_updated,required"`
	JSON        timestampedBrc20QuantitiesJSON        `json:"-"`
}

// timestampedBrc20QuantitiesJSON contains the JSON metadata for the struct
// [TimestampedBrc20Quantities]
type timestampedBrc20QuantitiesJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBrc20Quantities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBrc20QuantitiesJSON) RawJSON() string {
	return r.raw
}

type TimestampedBrc20QuantitiesLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                                     `json:"block_height,required"`
	JSON        timestampedBrc20QuantitiesLastUpdatedJSON `json:"-"`
}

// timestampedBrc20QuantitiesLastUpdatedJSON contains the JSON metadata for the
// struct [TimestampedBrc20QuantitiesLastUpdated]
type timestampedBrc20QuantitiesLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBrc20QuantitiesLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBrc20QuantitiesLastUpdatedJSON) RawJSON() string {
	return r.raw
}
