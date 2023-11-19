package crypto_price_check_dashboard

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"strings"
)

func printPricesCSV(prices map[string]float64) {
	file, err := os.Create("prices.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for coin, price := range prices {
		err := writer.Write([]string{coin, fmt.Sprintf("%.2f", price)})
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}

func plotPrices(prices map[string]float64) {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Cryptocurrency Prices"
	p.X.Label.Text = "Cryptocurrency"
	p.Y.Label.Text = "Price"

	pts := make(plotter.XYs, len(prices))
	i := 0
	for coin, price := range prices {
		pts[i].X = float64(i + 1)
		pts[i].Y = price
		i++
	}

	bars, err := plotter.NewBarChart(pts, vg.Points(15))
	if err != nil {
		log.Fatal(err)
	}
	p.Add(bars)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "prices.png"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . [cryptocurrency] [currency]")
		return
	}

	coinIDs := strings.Split(os.Args[1], ",")
	currency := os.Args[2]

	client := NewCoinGeckoClient()
	cache := NewCache()

	cacheKey := fmt.Sprintf("%s-%s", os.Args[1], currency)
	if prices, found := cache.Get(cacheKey); found {
		fmt.Println("Cached Prices:")
		printPricesCSV(prices) // Export to CSV
		// plotPrices(prices) // Uncomment to create a plot
		return
	}

	prices, err := client.GetCoinsPrice(coinIDs, currency)
	if err != nil {
		log.Fatalf("Error fetching prices: %v", err)
	}

	cache.Set(cacheKey, prices)
	printPricesCSV(prices) // Export to CSV
	// plotPrices(prices) // Uncomment to create a plot
}

func exportData(prices map[string]float64, format string) {
	switch format {
	case "csv":
		printPricesCSV(prices)
	case "json":
		// Implement JSON export functionality
	default:
		log.Fatalf("Unsupported export format: %s", format)
	}
}
