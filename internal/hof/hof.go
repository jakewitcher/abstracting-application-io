package hof

import (
	"HidingIoBehindInterfaces/internal/domain"
	"encoding/json"
	"io/ioutil"
	"log"
)

func CalculateOrderTotal(providerFunc func() []domain.LineItem) float64 {
	lineItems := providerFunc()

	var orderTotal float64
	for _, lineItem := range lineItems {
		orderTotal += lineItem.Price
	}

	return orderTotal
}

func FileBasedLineItemsProviderFuncFactory(filepath string) domain.ProviderFunc {
	return func() []domain.LineItem {
		file, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Fatalf("unable to read file %s", filepath)
		}

		var lineItems []domain.LineItem
		err = json.Unmarshal(file, &lineItems)
		if err != nil {
			log.Fatalf("unable to parse json for file %s", filepath)
		}

		return lineItems
	}
}

func InMemoryLineItemsProviderFuncFactory(lineItems []domain.LineItem) domain.ProviderFunc {
	return func() []domain.LineItem {
		return lineItems
	}
}