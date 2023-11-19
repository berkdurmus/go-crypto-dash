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

func (c *CoinGeckoClient) GetHistoricalPrice(coinID, currency, date string) (float64, error) {
    resp, err := c.httpClient.Get(fmt.Sprintf("%s/coins/%s/history?date=%s&localization=false", baseURL, coinID, date))
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return 0, err
    }

    marketData, ok := result["market_data"].(map[string]interface{})
    if !ok {
        return 0, fmt.Errorf("market data not available for %s", coinID)
    }

    currentPrice, ok := marketData["current_price"].(map[string]interface{})
    if !ok {
        return 0, fmt.Errorf("current price data not available for %s", coinID)
    }

    price, ok := currentPrice[currency].(float64)
    if !ok {
        return 0, fmt.Errorf("price information for %s in %s not available", coinID, currency)
    }
    return price, nil
}

