// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk

import (
	"context"
	"net/http"

	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/apijson"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/requestconfig"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

// GeneralInfoService contains methods and other services that help with
// interacting with the Maestro Dapp Platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeneralInfoService] method instead.
type GeneralInfoService struct {
	Options []option.RequestOption
}

// NewGeneralInfoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGeneralInfoService(opts ...option.RequestOption) (r *GeneralInfoService) {
	r = &GeneralInfoService{}
	r.Options = opts
	return
}

// Information about the state of the blockchain.
func (r *GeneralInfoService) Get(ctx context.Context, opts ...option.RequestOption) (res *TimestampedBlockchainInfo, err error) {
	opts = append(r.Options[:], opts...)
	path := "general/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TimestampedBlockchainInfo struct {
	Data        TimestampedBlockchainInfoData        `json:"data,required"`
	LastUpdated TimestampedBlockchainInfoLastUpdated `json:"last_updated,required"`
	JSON        timestampedBlockchainInfoJSON        `json:"-"`
}

// timestampedBlockchainInfoJSON contains the JSON metadata for the struct
// [TimestampedBlockchainInfo]
type timestampedBlockchainInfoJSON struct {
	Data        apijson.Field
	LastUpdated apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBlockchainInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBlockchainInfoJSON) RawJSON() string {
	return r.raw
}

type TimestampedBlockchainInfoData struct {
	// The hash of the currently best block
	Bestblockhash string `json:"bestblockhash,required"`
	// The current number of blocks processed in the server
	Blocks int64 `json:"blocks,required"`
	// Current network name as defined in BIP70 (main, test, regtest)
	Chain string `json:"chain,required"`
	// Total amount of work in active chain, in hexadecimal
	Chainwork string `json:"chainwork,required"`
	// The current difficulty
	Difficulty float64 `json:"difficulty,required"`
	// The current number of headers we have validated
	Headers int64 `json:"headers,required"`
	// Estimate of whether this node is in Initial Block Download mode
	Initialblockdownload bool `json:"initialblockdownload,required"`
	// Median time for the current best block
	Mediantime int64 `json:"mediantime,required"`
	// If the blocks are subject to pruning
	Pruned bool `json:"pruned,required"`
	// The estimated size of the block and undo files on disk
	SizeOnDisk int64 `json:"size_on_disk,required"`
	// Status of softforks in progress
	Softforks map[string]TimestampedBlockchainInfoDataSoftfork `json:"softforks,required"`
	// Estimate of verification progress [0..1]
	Verificationprogress float64 `json:"verificationprogress,required"`
	// Any network and blockchain warnings.
	Warnings string `json:"warnings,required"`
	// Whether automatic pruning is enabled (only present if pruning is enabled)
	AutomaticPruning bool `json:"automatic_pruning,nullable"`
	// Lowest-height complete block stored (only present if pruning is enabled)
	PruneHeight int64 `json:"prune_height,nullable"`
	// The target size used by pruning (only present if automatic pruning is enabled)
	PruneTargetSize int64                             `json:"prune_target_size,nullable"`
	JSON            timestampedBlockchainInfoDataJSON `json:"-"`
}

// timestampedBlockchainInfoDataJSON contains the JSON metadata for the struct
// [TimestampedBlockchainInfoData]
type timestampedBlockchainInfoDataJSON struct {
	Bestblockhash        apijson.Field
	Blocks               apijson.Field
	Chain                apijson.Field
	Chainwork            apijson.Field
	Difficulty           apijson.Field
	Headers              apijson.Field
	Initialblockdownload apijson.Field
	Mediantime           apijson.Field
	Pruned               apijson.Field
	SizeOnDisk           apijson.Field
	Softforks            apijson.Field
	Verificationprogress apijson.Field
	Warnings             apijson.Field
	AutomaticPruning     apijson.Field
	PruneHeight          apijson.Field
	PruneTargetSize      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TimestampedBlockchainInfoData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBlockchainInfoDataJSON) RawJSON() string {
	return r.raw
}

type TimestampedBlockchainInfoDataSoftfork struct {
	Active bool                                       `json:"active"`
	Bip9   TimestampedBlockchainInfoDataSoftforksBip9 `json:"bip9,nullable"`
	Height int64                                      `json:"height,nullable"`
	Type   TimestampedBlockchainInfoDataSoftforksType `json:"type,nullable"`
	JSON   timestampedBlockchainInfoDataSoftforkJSON  `json:"-"`
}

// timestampedBlockchainInfoDataSoftforkJSON contains the JSON metadata for the
// struct [TimestampedBlockchainInfoDataSoftfork]
type timestampedBlockchainInfoDataSoftforkJSON struct {
	Active      apijson.Field
	Bip9        apijson.Field
	Height      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBlockchainInfoDataSoftfork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBlockchainInfoDataSoftforkJSON) RawJSON() string {
	return r.raw
}

type TimestampedBlockchainInfoDataSoftforksBip9 struct {
	Since      int64                                                `json:"since,required"`
	StartTime  int64                                                `json:"start_time,required"`
	Status     TimestampedBlockchainInfoDataSoftforksBip9Status     `json:"status,required"`
	Timeout    int64                                                `json:"timeout,required"`
	Bit        int64                                                `json:"bit,nullable"`
	Statistics TimestampedBlockchainInfoDataSoftforksBip9Statistics `json:"statistics,nullable"`
	JSON       timestampedBlockchainInfoDataSoftforksBip9JSON       `json:"-"`
}

// timestampedBlockchainInfoDataSoftforksBip9JSON contains the JSON metadata for
// the struct [TimestampedBlockchainInfoDataSoftforksBip9]
type timestampedBlockchainInfoDataSoftforksBip9JSON struct {
	Since       apijson.Field
	StartTime   apijson.Field
	Status      apijson.Field
	Timeout     apijson.Field
	Bit         apijson.Field
	Statistics  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBlockchainInfoDataSoftforksBip9) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBlockchainInfoDataSoftforksBip9JSON) RawJSON() string {
	return r.raw
}

type TimestampedBlockchainInfoDataSoftforksBip9Status string

const (
	TimestampedBlockchainInfoDataSoftforksBip9StatusDefined  TimestampedBlockchainInfoDataSoftforksBip9Status = "defined"
	TimestampedBlockchainInfoDataSoftforksBip9StatusStarted  TimestampedBlockchainInfoDataSoftforksBip9Status = "started"
	TimestampedBlockchainInfoDataSoftforksBip9StatusLockedIn TimestampedBlockchainInfoDataSoftforksBip9Status = "locked_in"
	TimestampedBlockchainInfoDataSoftforksBip9StatusActive   TimestampedBlockchainInfoDataSoftforksBip9Status = "active"
	TimestampedBlockchainInfoDataSoftforksBip9StatusFailed   TimestampedBlockchainInfoDataSoftforksBip9Status = "failed"
)

func (r TimestampedBlockchainInfoDataSoftforksBip9Status) IsKnown() bool {
	switch r {
	case TimestampedBlockchainInfoDataSoftforksBip9StatusDefined, TimestampedBlockchainInfoDataSoftforksBip9StatusStarted, TimestampedBlockchainInfoDataSoftforksBip9StatusLockedIn, TimestampedBlockchainInfoDataSoftforksBip9StatusActive, TimestampedBlockchainInfoDataSoftforksBip9StatusFailed:
		return true
	}
	return false
}

type TimestampedBlockchainInfoDataSoftforksBip9Statistics struct {
	Count     int64                                                    `json:"count,required"`
	Elapsed   int64                                                    `json:"elapsed,required"`
	Period    int64                                                    `json:"period,required"`
	Possible  bool                                                     `json:"possible,required"`
	Threshold int64                                                    `json:"threshold,required"`
	JSON      timestampedBlockchainInfoDataSoftforksBip9StatisticsJSON `json:"-"`
}

// timestampedBlockchainInfoDataSoftforksBip9StatisticsJSON contains the JSON
// metadata for the struct [TimestampedBlockchainInfoDataSoftforksBip9Statistics]
type timestampedBlockchainInfoDataSoftforksBip9StatisticsJSON struct {
	Count       apijson.Field
	Elapsed     apijson.Field
	Period      apijson.Field
	Possible    apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBlockchainInfoDataSoftforksBip9Statistics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBlockchainInfoDataSoftforksBip9StatisticsJSON) RawJSON() string {
	return r.raw
}

type TimestampedBlockchainInfoDataSoftforksType string

const (
	TimestampedBlockchainInfoDataSoftforksTypeBuried TimestampedBlockchainInfoDataSoftforksType = "buried"
	TimestampedBlockchainInfoDataSoftforksTypeBip9   TimestampedBlockchainInfoDataSoftforksType = "bip9"
)

func (r TimestampedBlockchainInfoDataSoftforksType) IsKnown() bool {
	switch r {
	case TimestampedBlockchainInfoDataSoftforksTypeBuried, TimestampedBlockchainInfoDataSoftforksTypeBip9:
		return true
	}
	return false
}

type TimestampedBlockchainInfoLastUpdated struct {
	// The hash of the block
	BlockHash string `json:"block_hash,required"`
	// The height of the block in the blockchain
	BlockHeight int64                                    `json:"block_height,required"`
	JSON        timestampedBlockchainInfoLastUpdatedJSON `json:"-"`
}

// timestampedBlockchainInfoLastUpdatedJSON contains the JSON metadata for the
// struct [TimestampedBlockchainInfoLastUpdated]
type timestampedBlockchainInfoLastUpdatedJSON struct {
	BlockHash   apijson.Field
	BlockHeight apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimestampedBlockchainInfoLastUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timestampedBlockchainInfoLastUpdatedJSON) RawJSON() string {
	return r.raw
}
