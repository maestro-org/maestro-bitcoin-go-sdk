package main

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/maestro-bitcoin-go"
	"github.com/stainless-sdks/maestro-bitcoin-go/option"
)

func main() {
	client := maestrobitcoin.NewClient(
		option.WithAPIKey("My API Key"), // defaults to os.LookupEnv("API_KEY")
	)
	timestampedBlockchainInfo, err := client.General.Info.Get(context.TODO())
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", timestampedBlockchainInfo.Data)
}
