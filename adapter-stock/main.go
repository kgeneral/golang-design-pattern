package main

import (
	"golang-design-pattern/adapter-stock/adapter"
	"log"
)

func main() {
	var client = adapter.Client{}
	var xml = `
        <stock>
            <name>Apple</name>
            <datetime>2022-02-23T11:00:00Z</datetime>
            <price>115.0</price>
        </stock>`
	var stockSent = client.SendStockData(xml)
	log.Printf("%v", stockSent)

}
