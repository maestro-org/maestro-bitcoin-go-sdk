// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoin

import (
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// AssetService contains methods and other services that help with interacting with
// the maestro-bitcoin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetService] method instead.
type AssetService struct {
	Options []option.RequestOption
	Brc20   *AssetBrc20Service
	Runes   *AssetRuneService
}

// NewAssetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAssetService(opts ...option.RequestOption) (r *AssetService) {
	r = &AssetService{}
	r.Options = opts
	r.Brc20 = NewAssetBrc20Service(opts...)
	r.Runes = NewAssetRuneService(opts...)
	return
}
