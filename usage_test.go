// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package maestrobitcoingosdk_test

import (
	"context"
	"os"
	"testing"

	"github.com/maestro-org/maestro-bitcoin-go-sdk"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/internal/testutil"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

func TestUsage(t *testing.T) {
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
	timestampedBlock, err := client.Blocks.Latest.Get(context.TODO())
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", timestampedBlock.Data)
}
