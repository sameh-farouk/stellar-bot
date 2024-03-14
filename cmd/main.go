package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sameh-farouk/stellar-bot/internal/config"
	"github.com/sameh-farouk/stellar-bot/internal/monitoring"
	"github.com/sameh-farouk/stellar-bot/internal/stellar"
	"github.com/stellar/go/protocols/horizon"
	"github.com/stellar/go/xdr"
)

func main() {
	// create context
	ctx := context.Background()
	// Load configuration
	err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	// carete selling and buying assets from loaded GlobalConfig values
	var sellingAssetCode xdr.AssetCode4
	copy(sellingAssetCode[:], []byte(config.GlobalConfig.SellingAssetCode))

	var sellingAssetIssuer xdr.AccountId
    err = sellingAssetIssuer.SetAddress(config.GlobalConfig.SellingAssetIssuer)
    if err != nil {
		log.Fatalf("Error setting issuer address: %v", err)		
    }

	var buyingAssetCode xdr.AssetCode4
	copy(buyingAssetCode[:], []byte(config.GlobalConfig.SellingAssetCode))

	var buyingAssetIssuer xdr.AccountId
    err = buyingAssetIssuer.SetAddress(config.GlobalConfig.SellingAssetIssuer)
    if err != nil {
		log.Fatalf("Error setting issuer address: %v", err)		
    }

	sellingAsset := xdr.Asset{
		Type: xdr.AssetTypeAssetTypeCreditAlphanum4,
		AlphaNum4: &xdr.AlphaNum4{
			AssetCode: sellingAssetCode,
			Issuer: sellingAssetIssuer,
		},
	}

	buyingAsset := xdr.Asset{
		Type: xdr.AssetTypeAssetTypeCreditAlphanum4,
		AlphaNum4: &xdr.AlphaNum4{
			AssetCode: buyingAssetCode,
			Issuer: buyingAssetIssuer,
		},
	}


	fmt.Println("Trading bot configuration is loaded...")

	// Initialize Stellar client
	stellarClient, err := stellar.NewClient()
	if err != nil {
		log.Fatalf("Error initializing Stellar client: %v", err)
	}
	fmt.Println("stellar client is initialized...")

	fmt.Println("Trading bot for Stellar network is starting...")

	// Start price monitoring
	f := func(orderBook horizon.OrderBookSummary) {
		fmt.Println("Order book summary:", orderBook.Bids, orderBook.Asks)
	} 
	monitoring.StartPriceMonitoring(ctx, stellarClient, sellingAsset, buyingAsset, f)
}
