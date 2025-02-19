package spapi

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	SPAPIEndpointEU = "https://sellingpartnerapi-eu.amazon.com"
)

type SPAPIClient struct {
	Region   string
	Endpoint string
	LWA      *LWAClient
}

func NewSPAPIClient(region string, lwa *LWAClient) *SPAPIClient {
	return &SPAPIClient{
		Region:   region,
		Endpoint: SPAPIEndpointEU, // Adjust based on region
		LWA:      lwa,
	}
}

func (c *SPAPIClient) GetOrders(createdAfter time.Time) ([]byte, error) {
	accessToken, err := c.LWA.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get LWA token: %v", err)
	}

	url := fmt.Sprintf("%s/orders/v0/orders?MarketplaceIds=YOUR_MARKETPLACE_ID&CreatedAfter=%s",
		c.Endpoint,
		createdAfter.Format(time.RFC3339),
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add headers
	req.Header.Add("x-amz-access-token", accessToken)
	req.Header.Add("Content-Type", "application/json")

	// Sign the request with AWS SigV4
	if err := SignRequest(req, c.Region); err != nil {
		return nil, fmt.Errorf("failed to sign request: %v", err)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("SP-API error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}