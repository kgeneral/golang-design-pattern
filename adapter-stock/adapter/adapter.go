package adapter

import (
	"encoding/json"
	"encoding/xml"
	"log"
)

// Stock Name : name of the stock
// Stock Datetime : RFC 3339 format (2022-02-23T11:00:00Z)
// Stock Price : US Dollar
type Stock struct {
	XMLName  xml.Name `xml:"stock"`
	Name     string   `xml:"name"`
	Datetime string   `xml:"datetime"`
	Price    float64  `xml:"price"`
}

type ClientInterface interface {
	SendStockData(stockXmlString string) Stock
}

type StockService interface {
	Send(jsonString string) Stock
}

type StockJsonService struct {
	//...
}

func (s StockJsonService) Send(jsonString string) Stock {
	// doing some complex things here maybe?
	var stock Stock
	err := json.Unmarshal([]byte(jsonString), &stock)
	if err != nil {
		log.Fatalln("json parsing failed : ", err)
	}
	return stock
}

type XmlToJsonAdapter struct {
	stockService StockService
}

func (a XmlToJsonAdapter) SendStockData(stockXmlString string) Stock {
	a.stockService = StockJsonService{}
	var stock Stock
	xmlParseErr := xml.Unmarshal([]byte(stockXmlString), &stock)
	if xmlParseErr != nil {
		log.Fatalln("xml parsing failed : ", xmlParseErr)
	}
	jsonString, jsonParseErr := json.Marshal(stock)
	if jsonParseErr != nil {
		log.Fatalln("json parsing failed : ", xmlParseErr)
	}
	return a.stockService.Send(string(jsonString))
}

type Client struct {
	stockDataSendingAdapter ClientInterface
}

func (c Client) SendStockData(xml string) Stock {
	c.stockDataSendingAdapter = XmlToJsonAdapter{}
	return c.stockDataSendingAdapter.SendStockData(xml)
}
