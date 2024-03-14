package monitoring

import (
	"context"
	"fmt"

	stellar "github.com/sameh-farouk/stellar-bot/internal/stellar"
	"github.com/stellar/go/protocols/horizon"
	"github.com/stellar/go/xdr"
)

// StartPriceMonitoring starts monitoring the order book for the specified assets
func StartPriceMonitoring(ctx context.Context, client *stellar.Client, sellingAsset xdr.Asset, buyingAsset xdr.Asset, f func(orderBook horizon.OrderBookSummary)) error {

	err := client.StreamOrderBook(ctx, sellingAsset, buyingAsset, f)

	if err != nil {
		fmt.Println("Error streaming order book:", err)
		return err
	}
	return nil
}


