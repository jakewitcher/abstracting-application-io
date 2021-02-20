package embedded

import (
	"HidingIoBehindInterfaces/internal/domain"
	"encoding/json"
	"io/ioutil"
	"log"
)

func CalculateOrderTotal() float64 {
	filePath := "data/order_items.json"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("unable to read file %s", filePath)
	}

	var lineItems []domain.LineItem
	err = json.Unmarshal(file, &lineItems)
	if err != nil {
		log.Fatalf("unable to parse json for file %s", filePath)
	}

	var orderTotal float64
	for _, lineItem := range lineItems {
		orderTotal += lineItem.Price
	}

	return orderTotal
}
