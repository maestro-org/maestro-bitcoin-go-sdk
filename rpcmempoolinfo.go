// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk

import (
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// RpcMempoolInfoService contains methods and other services that help with
// interacting with the Maestro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRpcMempoolInfoService] method instead.
type RpcMempoolInfoService struct {
	Options []option.RequestOption
}

// NewRpcMempoolInfoService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRpcMempoolInfoService(opts ...option.RequestOption) (r *RpcMempoolInfoService) {
	r = &RpcMempoolInfoService{}
	r.Options = opts
	return
}
