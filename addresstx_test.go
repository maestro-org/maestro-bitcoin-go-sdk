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

func TestAddressTxListWithOptionalParams(t *testing.T) {
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
	_, err := client.Addresses.Txs.List(
		context.TODO(),
		"tb1qphcdyah2e4vtpxn56hsz3p6kapg90pl4x525kc",
		maestrobitcoin.AddressTxListParams{
			Count:  maestrobitcoin.F(int64(0)),
			Cursor: maestrobitcoin.F("cursor"),
			From:   maestrobitcoin.F(int64(0)),
			Order:  maestrobitcoin.F(maestrobitcoin.AddressTxListParamsOrderAsc),
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
