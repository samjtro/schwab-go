package quotes

import (
	"fmt"
	"net/http"
	"github.com/samjtro/go-tda/handler"
	"github.com/go-gota/gota"
)

type QUOTE struct {
	DATETIME	string
	TICKER		string
	BID		int
	ASK		int
	MARK		int
	VOLUME		int
	VOLATILITY	int
}

type HISTORY struct {
	DATETIME	string
	OPEN		int
	HIGH		int
	LOW		int
	CLOSE		int
	VOLUME		int
}

// realTime takes one parameter:
// ticker = "AAPL", etc.
func realTime(ticker string) *QUOTE {
	url := fmt.Sprintf(endpoint_realtime,ticker)
	req,_ := http.NewRequest("GET",url,nil)
	body := handler(req)
	
	return body
}

// priceHistory takes five parameters:
// ticker = "AAPL", etc.
// periodType = "day", "month", "year", "ytd" - default is "day"
// period = the number of periods to show
// frequencyType = the type of frequency with which each candle is formed; valid fTypes by pType
// "day": "minute"
// "month": "daily", "weekly"
// "year": "daily", "weekly", "monthly"
// "ytd": "daily", "weekly"
// frequency = the number of the frequencyType included in each candle; valid freqs by fType
// "minute": 1,5,10,15,30
// "daily": 1
// "weekly": 1
// "monthly": 1
func priceHistory(ticker,periodType,period,frequencyType,frequency string) DataFrame {
	url := fmt.Sprintf(endpoint_pricehistory,ticker)
	req,_ := http.NewRequest("GET",url,nil)
	q := req.URL.Query()
	q.Add("periodType",periodType)
	q.Add("period",period)
	q.Add("frequencyType",frequencyType)
	q.Add("frequency",frequency)
	req.URL.RawQuery = q.Encode()
	body := handler(req)

	df := dataframe.New(body)
	return df
}

