package main

import (
	"HidingIoBehindInterfaces/internal/abstracted"
	"HidingIoBehindInterfaces/internal/domain"
	"HidingIoBehindInterfaces/internal/embedded"
	"HidingIoBehindInterfaces/internal/extracted"
	"HidingIoBehindInterfaces/internal/hof"
	"fmt"
)

func main() {
	CalculateOrderTotalUsingEmbeddedSolution()
	CalculateOrderTotalUsingExtractedSolution()
	CalculateOrderTotalUsingAbstractedSolution()
	CalculateOrderTotalHigherOrderFunction()
}

func CalculateOrderTotalUsingEmbeddedSolution() {
	orderTotal := embedded.CalculateOrderTotal()
	fmt.Printf("Your total comes to %.2f", orderTotal)
}

func CalculateOrderTotalUsingExtractedSolution() {
	provider := domain.JsonOrderProvider{FilePath: "data/order_items.json"}
	orderTotal := extracted.CalculateOrderTotal(provider)
	fmt.Printf("Your total comes to %.2f", orderTotal)
}

func CalculateOrderTotalUsingAbstractedSolution() {
	jsonProvider := domain.JsonOrderProvider{FilePath: "data/order_items.json"}
	orderTotal := abstracted.CalculateOrderTotal(jsonProvider)
	fmt.Printf("Your total comes to %.2f", orderTotal)

	inMemoryProvider := domain.InMemoryOrderProvider{
		LineItems: []domain.LineItem{
			{Description: "Leather Recliner", Price: 2499},
			{Description: "End Table", Price: 249},
		},
	}

	orderTotal = abstracted.CalculateOrderTotal(inMemoryProvider)
	fmt.Printf("Your total comes to %.2f", orderTotal)
}

func CalculateOrderTotalHigherOrderFunction() {
	fileBasedProviderFunc := hof.FileBasedLineItemsProviderFuncFactory("data/order_items.json")
	orderTotal := hof.CalculateOrderTotal(fileBasedProviderFunc)
	fmt.Printf("Your total comes to %.2f", orderTotal)

	orderTotal = abstracted.CalculateOrderTotal(fileBasedProviderFunc)
	fmt.Printf("Your total comes to %.2f", orderTotal)

	lineItems := []domain.LineItem{
		{Description: "Leather Recliner", Price: 2499},
		{Description: "End Table", Price: 249},
	}

	inMemoryProviderFunc := hof.InMemoryLineItemsProviderFuncFactory(lineItems)
	orderTotal = hof.CalculateOrderTotal(inMemoryProviderFunc)
	fmt.Printf("Your total comes to %.2f", orderTotal)

	orderTotal = abstracted.CalculateOrderTotal(inMemoryProviderFunc)
	fmt.Printf("Your total comes to %.2f", orderTotal)
}
