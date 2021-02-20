package test

import (
	"HidingIoBehindInterfaces/internal/domain"
	"HidingIoBehindInterfaces/internal/hof"
	"testing"
)

var hofTestCases = []struct {
	lineItems []domain.LineItem
	expected  float64
}{
	{
		lineItems: []domain.LineItem{
			{Description: "A", Price: 85},
			{Description: "B", Price: 15},
		},
		expected: 100,
	},
	{
		lineItems: []domain.LineItem{
			{Description: "A", Price: 35.25},
			{Description: "B", Price: 95.5},
		},
		expected: 130.75,
	},
}

func TestHigherOrderFunctionCalculateOrderTotal(t *testing.T) {
	for _, test := range hofTestCases {
		providerFunc := hof.InMemoryLineItemsProviderFuncFactory(test.lineItems)

		if actual := hof.CalculateOrderTotal(providerFunc); actual != test.expected {
			t.Fatalf("expected %.2f, actual %.2f", test.expected, actual)
		}
	}
}