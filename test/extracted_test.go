package test

import (
	"HidingIoBehindInterfaces/internal/domain"
	"HidingIoBehindInterfaces/internal/extracted"
	"testing"
)

func TestExtractedCalculateOrderTotal(t *testing.T) {
	var expected float64 = 3476
	provider := domain.JsonOrderProvider{FilePath: "../data/order_items.json"}
	actual := extracted.CalculateOrderTotal(provider)

	if actual != expected {
		t.Fatalf("expected %.2f, actual %.2f", expected, actual)
	}
}
