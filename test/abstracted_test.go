package test

import (
	"HidingIoBehindInterfaces/internal/abstracted"
	"HidingIoBehindInterfaces/internal/domain"
	"testing"
)

var abstractedTestCases = []struct {
	provider domain.InMemoryOrderProvider
	expected float64
}{
	{
		provider: domain.InMemoryOrderProvider{
			LineItems: []domain.LineItem{
				{Description: "A", Price: 85},
				{Description: "B", Price: 15},
			},
		},
		expected: 100,
	},
	{
		provider: domain.InMemoryOrderProvider{
			LineItems: []domain.LineItem{
				{Description: "A", Price: 35.25},
				{Description: "B", Price: 95.5},
			},
		},
		expected: 130.75,
	},
}

func TestAbstractedCalculateOrderTotal(t *testing.T) {
	for _, test := range abstractedTestCases {
		if actual := abstracted.CalculateOrderTotal(test.provider); actual != test.expected {
			t.Fatalf("expected %.2f, actual %.2f", test.expected, actual)
		}
	}
}
