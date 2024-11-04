// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk

import (
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// RpcTransactionService contains methods and other services that help with
// interacting with the Maestro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRpcTransactionService] method instead.
type RpcTransactionService struct {
	Options []option.RequestOption
}

// NewRpcTransactionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRpcTransactionService(opts ...option.RequestOption) (r *RpcTransactionService) {
	r = &RpcTransactionService{}
	r.Options = opts
	return
}
