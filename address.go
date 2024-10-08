// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk

import (
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// AddressService contains methods and other services that help with interacting
// with the maestro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressService] method instead.
type AddressService struct {
	Options []option.RequestOption
	Brc20   *AddressBrc20Service
	Runes   *AddressRuneService
	Txs     *AddressTxService
	Utxos   *AddressUtxoService
}

// NewAddressService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAddressService(opts ...option.RequestOption) (r *AddressService) {
	r = &AddressService{}
	r.Options = opts
	r.Brc20 = NewAddressBrc20Service(opts...)
	r.Runes = NewAddressRuneService(opts...)
	r.Txs = NewAddressTxService(opts...)
	r.Utxos = NewAddressUtxoService(opts...)
	return
}
