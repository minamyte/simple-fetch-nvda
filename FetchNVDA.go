package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ChartResponse struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency           string  `json:"currency"`
				RegularMarketPrice float64 `json:"regularMarketPrice"`
			} `json:"meta"`
		} `json:"result"`
	} `json:"chart"`
}

func getPrice() {
	url := "https://query2.finance.yahoo.com/v8/finance/chart/NVDA?period1=1754244000&period2=1754416800&interval=1m"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0") // ✅ Add User-Agent

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)

	var result ChartResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("JSON parse error:", err)
		return
	}

	if len(result.Chart.Result) > 0 {
		chart := result.Chart.Result[0]
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] NVDA Price: %.2f %s\n", timestamp, chart.Meta.RegularMarketPrice, chart.Meta.Currency)
	} else {
		fmt.Println("No data returned.")
	}
}

func main() {
	for {
		getPrice()
		time.Sleep(1 * time.Second)
	}
}
