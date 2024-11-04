// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk

import (
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// RpcMempoolTransactionDescendantService contains methods and other services that
// help with interacting with the Maestro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRpcMempoolTransactionDescendantService] method instead.
type RpcMempoolTransactionDescendantService struct {
	Options []option.RequestOption
}

// NewRpcMempoolTransactionDescendantService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRpcMempoolTransactionDescendantService(opts ...option.RequestOption) (r *RpcMempoolTransactionDescendantService) {
	r = &RpcMempoolTransactionDescendantService{}
	r.Options = opts
	return
}
