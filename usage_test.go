// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoin_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/maestro-bitcoin-go"
	"github.com/stainless-sdks/maestro-bitcoin-go/internal/testutil"
	"github.com/stainless-sdks/maestro-bitcoin-go/option"
)

func TestUsage(t *testing.T) {
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
	timestampedBlockchainInfo, err := client.General.Info.Get(context.TODO())
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", timestampedBlockchainInfo.Data)
}
