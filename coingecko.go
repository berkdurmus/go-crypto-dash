package crypto_price_check_dashboard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const baseURL = "https://api.coingecko.com/api/v3"

type CoinGeckoClient struct {
	httpClient *http.Client
}

func NewCoinGeckoClient() *CoinGeckoClient {
	return &CoinGeckoClient{httpClient: &http.Client{}}
}

func (c *CoinGeckoClient) GetCoinsPrice(coinIDs []string, currency string) (map[string]float64, error) {
	ids := strings.Join(coinIDs, ",")
	resp, err := c.httpClient.Get(fmt.Sprintf("%s/simple/price?ids=%s&vs_currencies=%s", baseURL, ids, currency))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	prices := make(map[string]float64)
	for _, id := range coinIDs {
		price, ok := result[id][currency]
		if !ok {
			return nil, fmt.Errorf("price information for %s not available", id)
		}
		prices[id] = price
	}
	return prices, nil
}
