package crypto_price_check_dashboard

import (
	"encoding/csv",
	"encoding/json",
	"fmt"
	"log"
	"os",
	"net/http",
        "github.com/gorilla/websocket"
)

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"strings"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}
func handleConnections(w http.ResponseWriter, r *http.Request) {
    // Upgrade initial GET request to a websocket
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close()

    // TODO: Consider listening for messages from the WebSocket, if necessary.

    // Main loop for sending data
    for {
        prices, _ := // fetch the latest prices
        err := ws.WriteJSON(prices)
        if err != nil {
            log.Printf("error: %v", err)
            break
        }
        time.Sleep(time.Second * 10) // Update interval
    }
}

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

func printPricesJSON(prices map[string]float64) {
    file, err := os.Create("prices.json")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "    ") // Optional: for pretty-printing the JSON

    if err := encoder.Encode(prices); err != nil {
        log.Fatal("Cannot write to file", err)
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

	 // Set up WebSocket route
	    http.HandleFunc("/ws", handleConnections)
	
	    // Start the server on localhost and log errors
	    log.Println("http server started on :8000")
	    err := http.ListenAndServe(":8000", nil)
	    if err != nil {
	        log.Fatal("ListenAndServe: ", err)
	    }
}

func exportData(prices map[string]float64, format string) {
	switch format {
	case "csv":
		printPricesCSV(prices)
	case "json":
       		printPricesJSON(prices) // New JSON export functionality
	default:
		log.Fatalf("Unsupported export format: %s", format)
	}
}
