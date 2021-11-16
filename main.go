package main

import (
	"rate-watcher/config"
)

type TickerPrice struct {
	Symbol	string	`json:"symbol"`
	Price 	string	`json:"price"`
}

func main() {
	
	config.ParseFlags()

	// req, errNewRequest := http.NewRequest("GET", "https://api1.binance.com/api/v3/ticker/price", nil)
	// if errNewRequest != nil {
	// 	fmt.Printf("0x01")
	// 	os.Exit(1)
	// }

	// query := req.URL.Query()
	// query.Add("symbol", "ETHUSDT")
	// req.URL.RawQuery = query.Encode()
	
	// client := &http.Client{}
	// res, errClientRequest := client.Do(req)
	// if errClientRequest != nil {
	// 	fmt.Printf("0x02")
	// 	os.Exit(2)
	// }
	// defer res.Body.Close()

	
	// if res.StatusCode != http.StatusOK {
	// 	log.Printf("0x03")
	// 	os.Exit(3)
	// }
	
	
	// bodyBytes, errReadBody := ioutil.ReadAll(res.Body)
	// if errReadBody != nil {
	// 	log.Printf("0x03")
	// 	os.Exit(4)
	// }
	
	// var tickerPrice TickerPrice
	// if err := json.Unmarshal(bodyBytes, &tickerPrice); err != nil {
	// 	log.Printf("0x04")
	// 	os.Exit(5)
	// }

	// price, _ := strconv.ParseFloat(tickerPrice.Price, 64)
	// fmt.Printf("%.2f", price)
}