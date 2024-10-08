// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk

import (
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// RpcMempoolService contains methods and other services that help with interacting
// with the maestro-bitcoin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRpcMempoolService] method instead.
type RpcMempoolService struct {
	Options      []option.RequestOption
	Info         *RpcMempoolInfoService
	Transactions *RpcMempoolTransactionService
}

// NewRpcMempoolService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRpcMempoolService(opts ...option.RequestOption) (r *RpcMempoolService) {
	r = &RpcMempoolService{}
	r.Options = opts
	r.Info = NewRpcMempoolInfoService(opts...)
	r.Transactions = NewRpcMempoolTransactionService(opts...)
	return
}
