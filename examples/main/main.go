package main

import (
	"context"
	"fmt"

	maestrobitcoingosdk "github.com/maestro-org/maestro-bitcoin-go-sdk"
	"github.com/maestro-org/maestro-bitcoin-go-sdk/option"
)

func main() {
	client := maestrobitcoingosdk.NewClient(
		option.WithAPIKey("My API Key"), // defaults to os.LookupEnv("API_KEY")
	)
	timestampedBlockchainInfo, err := client.General.Info.Get(context.TODO())
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", timestampedBlockchainInfo.Data)
}
