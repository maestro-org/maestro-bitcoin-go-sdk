// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoin_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/maestro-org/maestro-bitcoin-go-sdk"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/testutil"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

func TestAddressUtxoListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := maestrobitcoin.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Addresses.Utxos.List(
		context.TODO(),
		"tb1qphcdyah2e4vtpxn56hsz3p6kapg90pl4x525kc",
		maestrobitcoin.AddressUtxoListParams{
			Count:                maestrobitcoin.F(int64(0)),
			Cursor:               maestrobitcoin.F("cursor"),
			ExcludeMetaprotocols: maestrobitcoin.F(true),
			FilterDust:           maestrobitcoin.F(true),
			FilterDustThreshold:  maestrobitcoin.F(int64(0)),
			From:                 maestrobitcoin.F(int64(0)),
			Order:                maestrobitcoin.F(maestrobitcoin.AddressUtxoListParamsOrderAsc),
			To:                   maestrobitcoin.F(int64(0)),
		},
	)
	if err != nil {
		var apierr *maestrobitcoin.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
