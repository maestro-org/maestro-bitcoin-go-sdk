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

func TestTransactionGet(t *testing.T) {
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
	err := client.Transactions.Get(context.TODO(), "tx_hash")
	if err != nil {
		var apierr *maestrobitcoin.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionSubmit(t *testing.T) {
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
	_, err := client.Transactions.Submit(context.TODO(), maestrobitcoin.TransactionSubmitParams{
		Body: map[string]interface{}{},
	})
	if err != nil {
		var apierr *maestrobitcoin.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
