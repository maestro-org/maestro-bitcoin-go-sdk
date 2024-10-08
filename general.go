// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk

import (
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// GeneralService contains methods and other services that help with interacting
// with the maestro API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeneralService] method instead.
type GeneralService struct {
	Options []option.RequestOption
	Info    *GeneralInfoService
}

// NewGeneralService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGeneralService(opts ...option.RequestOption) (r *GeneralService) {
	r = &GeneralService{}
	r.Options = opts
	r.Info = NewGeneralInfoService(opts...)
	return
}
