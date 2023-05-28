package services

import (
	"encoding/json"
	"fmt"
	"gses2-btc-app/utils"
	"io"
	"net/http"
	"net/url"
)

func createUrl(fromCurr string, toCurr string) string {
	baseUrl := utils.CoingeckoApi
	path := "api/v3/simple/price"
	params := url.Values{}
	params.Add("ids", fromCurr)
	params.Add("vs_currencies", toCurr)

	u, _ := url.ParseRequestURI(baseUrl)
	u.Path = path
	u.RawQuery = params.Encode()

	url := fmt.Sprintf("%v", u)
	return url
}

func GetCurrentPrice(fromCurr string, toCurr string) (int64, error) {

	urlStr := createUrl(fromCurr, toCurr)
	resp, err := http.Get(urlStr)

	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result map[string]any
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return 0, err
	}

	getRate := result[fromCurr].(map[string]any)

	m := make(map[string]int64)
	for key, value := range getRate {
		// Each value is an `any` type, that is type asserted as a string
		m[key] = int64(value.(float64))
	}

	return m[toCurr], nil
}
