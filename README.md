
# Cryptocurrency Price Checker with Visualization

A Go-based application that fetches and displays cryptocurrency prices using the CoinGecko API. This project demonstrates HTTP requests, JSON parsing, caching mechanisms, and data visualization in Go.

## Features

- Fetch current prices of various cryptocurrencies.
- Support for multiple cryptocurrencies and fiat currencies.
- Caching of price data to reduce API calls.
- Visualization of cryptocurrency prices as a CSV file or a bar chart image.
- Command-line interface for easy and flexible use.

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You'll need Go installed on your machine (version 1.15 or higher). Additionally, for visualization, the Gonum Plot library is used:

```bash
go get -u gonum.org/v1/plot/...
```

### Installing

Clone the repository to your local machine:

```bash
git clone https://github.com/[your-username]/crypto-price-checker.git
```

Navigate to the project directory:

```bash
cd crypto-price-checker
```

### Running the Application

Run the server using:

```bash
go run . bitcoin,ethereum usd
```

Replace `bitcoin,ethereum` with the cryptocurrencies of your choice and `usd` with your preferred fiat currency.

## Usage

The application can be used via command-line interface. Here are some example usages:

- To fetch prices for Bitcoin (BTC) and Ethereum (ETH) in USD:
  
  ```bash
  go run . bitcoin,ethereum usd
  ```

- The application will generate a CSV file (`prices.csv`) or a bar chart image (`prices.png`) based on the uncommented line in the `main.go` file.

## Visualization

This application supports two types of visualizations:

1. **CSV Output**: Generates a CSV file with cryptocurrency prices, which can be imported into tools like Excel or Google Sheets for custom visualizations.
2. **Bar Chart Image**: Generates a bar chart image displaying the prices of the specified cryptocurrencies.

## Built With

- [Go](https://golang.org/) - The Go Programming Language
- [Gonum Plot](https://gonum.org/v1/plot) - A plotting library for Go

## Contributing

Please read [CONTRIBUTING.md](https://github.com/[your-username]/crypto-price-checker/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

- **[Your Name](https://github.com/[your-username])** - *Initial work*

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

- Hat tip to the [CoinGecko API](https://www.coingecko.com/en/api) for providing cryptocurrency data.
- Inspired by the Go community and their support for open-source development.

