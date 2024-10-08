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

// TransactionService contains methods and other services that help with
// interacting with the maestro-bitcoin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	Options []option.RequestOption
	Psbt    *TransactionPsbtService
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r *TransactionService) {
	r = &TransactionService{}
	r.Options = opts
	r.Psbt = NewTransactionPsbtService(opts...)
	return
}

// Information about a specific transaction.
func (r *TransactionService) Get(ctx context.Context, txHash string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if txHash == "" {
		err = errors.New("missing required tx_hash parameter")
		return
	}
	path := fmt.Sprintf("transactions/%s", txHash)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Submits a signed transaction to the network.
func (r *TransactionService) Submit(ctx context.Context, body TransactionSubmitParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := "transactions/submit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TransactionSubmitParams struct {
	Body interface{} `json:"body,required"`
}

func (r TransactionSubmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
