package stellar

import (
	"context"
	"fmt"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/xdr"
)

// Client represents a Stellar client
type Client struct {
	inner *horizonclient.Client
}

// NewClient creates a new Stellar client
func NewClient() (*Client, error) {
	// Perform stellar go sdk client initialization here based on configuration
	client := horizonclient.DefaultPublicNetClient
	return &Client{client}, nil
}

// StreamOrderBook streams the order book for the specified assets
func (c *Client) StreamOrderBook(ctx context.Context, sellingAsset xdr.Asset, buyingAsset xdr.Asset, f horizonclient.OrderBookHandler) error {
	// create horizonclient.OrderBookRequest
	orderBookRequest := horizonclient.OrderBookRequest{
		SellingAssetType: horizonclient.AssetType(sellingAsset.Type),
		SellingAssetCode: sellingAsset.GetCode(),
		SellingAssetIssuer: sellingAsset.GetIssuer(),
		BuyingAssetType: horizonclient.AssetType(buyingAsset.Type),
		BuyingAssetCode: buyingAsset.GetCode(),
		BuyingAssetIssuer: buyingAsset.GetIssuer(),
	}
	
	c.inner.StreamOrderBooks(ctx, orderBookRequest, f)
	fmt.Println("Example function using the Stellar client")
	return nil
}