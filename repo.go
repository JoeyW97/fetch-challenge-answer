package main

import (
	"fmt"
	"github.com/google/uuid"
)

var receipts = map[string]Receipt{}

func addReceipt(newReceipt Receipt) string {
	id := uuid.New().String()
	receipts[id] = newReceipt
	fmt.Printf("Added new receipt %v\n", newReceipt)
	return id
}

func getReceiptById(id string) (Receipt, bool) {
	receipt, ok := receipts[id]
	return receipt, ok
}
