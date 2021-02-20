package domain

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type LineItem struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type OrderProvider interface {
	GetLineItems() []LineItem
}

type JsonOrderProvider struct {
	FilePath string
}

func (r JsonOrderProvider) GetLineItems() []LineItem {
	file, err := ioutil.ReadFile(r.FilePath)
	if err != nil {
		log.Fatalf("unable to read file %s", r.FilePath)
	}

	var lineItems []LineItem
	err = json.Unmarshal(file, &lineItems)
	if err != nil {
		log.Fatalf("unable to parse json for file %s", r.FilePath)
	}

	return lineItems
}

type InMemoryOrderProvider struct {
	LineItems []LineItem
}

func (r InMemoryOrderProvider) GetLineItems() []LineItem {
	return r.LineItems
}

type ProviderFunc func() []LineItem

func (f ProviderFunc) GetLineItems() []LineItem {
	return f()
}
