package evepraisal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Response is the response structure from a call to the EvePraisal API
type Response struct {
	Appraisal appraisal `json:"appraisal,omitempty"`
}

type request struct {
	Items      []*AppraisalItem `json:"items"`
	MarketName string           `json:"market_name"`
}

// AppraisalItem is an item with quantity you want
type AppraisalItem struct {
	Name     string `json:"name,omitempty"`
	TypeID   uint32 `json:"type_id,omitempty"`
	Quantity uint32 `json:"quantity"`
}

type appraisal struct {
	ID         string `json:"id,omitempty"`
	Created    int32  `json:"created,omitempty"`
	MarketName string `json:"market_name,omitempty"`
	Totals     totals `json:"totals,omitempty"`
}

type totals struct {
	Buy  float64 `json:"buy,omitempty"`
	Sell float64 `json:"sell,omitempty"`
}

// AppraiseSingle appraises a given item at the specified market
func (evepraisal Client) AppraiseSingle(item string, market string) (*Response, error) {
	uri := fmt.Sprintf("https://evepraisal.com/appraisal.json?market=%s&raw_textarea=%s", market, item)
	request, error := http.NewRequest("POST", uri, nil)
	if error != nil {
		log.Println("Error creating new request: ", error)
		return nil, error
	}

	response, error := evepraisal.client.Do(request)
	if error != nil {
		log.Println("Error pulling market totals from evepraisal: ", error)
		return nil, error
	}

	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		log.Println("Invalid response from EvePraisal: ", error)
		return nil, error
	}

	var appraisal Response
	error = json.Unmarshal(body, &appraisal)
	if error != nil {
		log.Println("Error parsing appraisal")
		return nil, error
	}

	return &appraisal, nil
}

// AppraiseAll appraises a full list of items as the specified market
func (evepraisal Client) AppraiseAll(items []*AppraisalItem, market string) (*Response, error) {
	uri := fmt.Sprintf("https://evepraisal.com/appraisal/structured.json?market=%s", market)
	body, error := json.Marshal(&request{Items: items, MarketName: market})
	request, error := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	if error != nil {
		log.Println("Error creating new request: ", error)
		return nil, error
	}

	response, error := evepraisal.client.Do(request)
	if error != nil {
		log.Println("Error pulling market totals from evepraisal: ", error)
		return nil, error
	}

	body, error = ioutil.ReadAll(response.Body)
	if error != nil {
		log.Println("Invalid response from EvePraisal: ", error)
		return nil, error
	}

	var appraisal Response
	error = json.Unmarshal(body, &appraisal)
	if error != nil {
		log.Println("Error parsing appraisal")
		return nil, error
	}

	return &appraisal, nil
}
