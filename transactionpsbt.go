// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk

import (
	"context"
	"net/http"

	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/apijson"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/requestconfig"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// TransactionPsbtService contains methods and other services that help with
// interacting with the maestro-bitcoin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionPsbtService] method instead.
type TransactionPsbtService struct {
	Options []option.RequestOption
}

// NewTransactionPsbtService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionPsbtService(opts ...option.RequestOption) (r *TransactionPsbtService) {
	r = &TransactionPsbtService{}
	r.Options = opts
	return
}

// Decode PSBT
func (r *TransactionPsbtService) Decode(ctx context.Context, body TransactionPsbtDecodeParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "transactions/psbt/decode"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type TransactionPsbtDecodeParams struct {
	Body interface{} `json:"body,required"`
}

func (r TransactionPsbtDecodeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
