package main

import (
	"fmt"
	"strconv"
	"time"
)

type ReceiptDto struct {
	Retailer     string    `json:"retailer"`
	PurchaseDate string    `json:"purchaseDate"`
	PurchaseTime string    `json:"purchaseTime"`
	Total        string    `json:"total"`
	Items        []ItemDto `json:"items"`
}

type ItemDto struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string
	PurchaseDate time.Time
	PurchaseTime time.Time
	Total        float64
	Items        []Item
}

type Item struct {
	ShortDescription string
	Price            float64
}

func (receiptDto *ReceiptDto) toReceipt() (Receipt, error) {
	purchaseDate, err := time.Parse("2006-01-02", receiptDto.PurchaseDate)
	if err != nil {
		fmt.Println(err)
		return Receipt{}, err
	}
	purchaseTime, err := time.Parse("15:04", receiptDto.PurchaseTime)
	if err != nil {
		fmt.Println(err)
		return Receipt{}, err
	}
	total, err := strconv.ParseFloat(receiptDto.Total, 64)
	if err != nil {
		fmt.Println(err)
		return Receipt{}, err
	}
	var items []Item
	for _, item := range receiptDto.Items {
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			fmt.Println(err)
			return Receipt{}, err
		}
		items = append(items, Item{ShortDescription: item.ShortDescription, Price: price})
	}
	return Receipt{
		Retailer:     receiptDto.Retailer,
		PurchaseDate: purchaseDate,
		PurchaseTime: purchaseTime,
		Total:        total,
		Items:        items,
	}, nil
}
