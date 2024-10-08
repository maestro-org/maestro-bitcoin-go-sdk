// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoin_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/maestro-bitcoin-go"
	"github.com/stainless-sdks/maestro-bitcoin-go/internal/testutil"
	"github.com/stainless-sdks/maestro-bitcoin-go/option"
)

func TestAssetRuneUtxoListWithOptionalParams(t *testing.T) {
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
	_, err := client.Assets.Runes.Utxos.List(
		context.TODO(),
		"2519999:31",
		maestrobitcoin.AssetRuneUtxoListParams{
			Count:  maestrobitcoin.F(int64(0)),
			Cursor: maestrobitcoin.F("cursor"),
			From:   maestrobitcoin.F(int64(0)),
			Order:  maestrobitcoin.F(maestrobitcoin.AssetRuneUtxoListParamsOrderAsc),
			To:     maestrobitcoin.F(int64(0)),
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
