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

func TestAssetRuneUtxoListWithOptionalParams(t *testing.T) {
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
	_, err := client.Assets.Runes.Utxos.List(
		context.TODO(),
		"2519999:31",
		maestrobitcoingosdk.AssetRuneUtxoListParams{
			Count:  maestrobitcoingosdk.F(int64(0)),
			Cursor: maestrobitcoingosdk.F("cursor"),
			From:   maestrobitcoingosdk.F(int64(0)),
			Order:  maestrobitcoingosdk.F(maestrobitcoingosdk.AssetRuneUtxoListParamsOrderAsc),
			To:     maestrobitcoingosdk.F(int64(0)),
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
