// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk_test

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
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := maestrobitcoingosdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Addresses.Utxos.List(
		context.TODO(),
		"tb1qphcdyah2e4vtpxn56hsz3p6kapg90pl4x525kc",
		maestrobitcoingosdk.AddressUtxoListParams{
			Count:                maestrobitcoingosdk.F(int64(0)),
			Cursor:               maestrobitcoingosdk.F("cursor"),
			ExcludeMetaprotocols: maestrobitcoingosdk.F(true),
			FilterDust:           maestrobitcoingosdk.F(true),
			FilterDustThreshold:  maestrobitcoingosdk.F(int64(0)),
			From:                 maestrobitcoingosdk.F(int64(0)),
			Order:                maestrobitcoingosdk.F(maestrobitcoingosdk.AddressUtxoListParamsOrderAsc),
			To:                   maestrobitcoingosdk.F(int64(0)),
		},
	)
	if err != nil {
		var apierr *maestrobitcoingosdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
