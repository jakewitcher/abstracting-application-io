package abstracted

import "HidingIoBehindInterfaces/internal/domain"

func CalculateOrderTotal(provider domain.OrderProvider) float64 {
	lineItems := provider.GetLineItems()

	var orderTotal float64
	for _, lineItem := range lineItems {
		orderTotal += lineItem.Price
	}

	return orderTotal
}
