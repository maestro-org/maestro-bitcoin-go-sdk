// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk

import (
	"context"
	"net/http"

	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/requestconfig"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// BlockLatestService contains methods and other services that help with
// interacting with the maestro-bitcoin API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockLatestService] method instead.
type BlockLatestService struct {
	Options []option.RequestOption
}

// NewBlockLatestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlockLatestService(opts ...option.RequestOption) (r *BlockLatestService) {
	r = &BlockLatestService{}
	r.Options = opts
	return
}

// Information about the latest block on the chain.
func (r *BlockLatestService) Get(ctx context.Context, opts ...option.RequestOption) (res *TimestampedBlock, err error) {
	opts = append(r.Options[:], opts...)
	path := "blocks/latest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
