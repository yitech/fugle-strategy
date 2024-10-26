package main

import (
	"context"
	"fmt"
	"os"

	fuglego "github.com/yitech/fugle-go"
)

func main() {
	symbol := "2330"         // string |
	fromDate := "2024-10-01" // string |
	toDate := "2024-10-26"   // string |
	resolution := "D"        // string |  (optional) (default to "D")

	configuration := fuglego.NewConfiguration()
	configuration.Servers = fuglego.ServerConfigurations{
		{
			URL:         "http://localhost:8000",
			Description: "Local server",
		},
	}
	apiClient := fuglego.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.GetCandlesApiV1HistoricalCandlesGet(context.Background()).Symbol(symbol).FromDate(fromDate).ToDate(toDate).Resolution(resolution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.GetCandlesApiV1HistoricalCandlesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCandlesApiV1HistoricalCandlesGet`: KLinesResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.GetCandlesApiV1HistoricalCandlesGet`: %v\n", resp)
}
