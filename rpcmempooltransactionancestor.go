// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk

import (
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// RpcMempoolTransactionAncestorService contains methods and other services that
// help with interacting with the Maestro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRpcMempoolTransactionAncestorService] method instead.
type RpcMempoolTransactionAncestorService struct {
	Options []option.RequestOption
}

// NewRpcMempoolTransactionAncestorService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRpcMempoolTransactionAncestorService(opts ...option.RequestOption) (r *RpcMempoolTransactionAncestorService) {
	r = &RpcMempoolTransactionAncestorService{}
	r.Options = opts
	return
}
