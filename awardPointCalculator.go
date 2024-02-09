package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func calculateAwardPoint(receipt Receipt) int {
	point := 0
	retailerPoint := getRetailerNamePoint(receipt)
	fmt.Printf("%d points - retailer name has %d characters\n", retailerPoint, len(receipt.Retailer))
	point += retailerPoint
	totalAmountPoint := getTotalAmountPoint(receipt)
	fmt.Printf("%d points - total is a round dollar amount\n", totalAmountPoint)
	point += totalAmountPoint
	totalOf25Point := getTotalOf25Point(receipt)
	fmt.Printf("%d points - total is a multiple of 0.25\n", totalOf25Point)
	point += totalOf25Point
	numberOfItemPoint := getNumberOfItemPoint(receipt)
	fmt.Printf("%d points - %d items\n", numberOfItemPoint, len(receipt.Items))
	point += numberOfItemPoint
	shortDescriptionPoint := getShortDescriptionPoint(receipt)
	fmt.Printf("%d points - short descriptions\n", shortDescriptionPoint)
	point += shortDescriptionPoint
	dayPoint := getPurchaseDayPoint(receipt)
	fmt.Printf("%d points - purchase day\n", dayPoint)
	point += dayPoint
	timePoint := getPurchaseTimePoint(receipt)
	fmt.Printf("%d points - purchase time\n", timePoint)
	point += timePoint
	return point
}

func getRetailerNamePoint(receipt Receipt) int {
	point := 0
	for _, c := range receipt.Retailer {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			point++
		}
	}
	return point
}

func getTotalAmountPoint(receipt Receipt) int {
	point := 0
	if math.Trunc(receipt.Total) == receipt.Total {
		point = 50
	}
	return point
}

func getTotalOf25Point(receipt Receipt) int {
	point := 0
	if math.Mod(receipt.Total, 0.25) == 0 {
		point = 25
	}
	return point
}

func getNumberOfItemPoint(receipt Receipt) int {
	return len(receipt.Items) / 2 * 5
}

func getShortDescriptionPoint(receipt Receipt) int {
	point := 0
	for _, item := range receipt.Items {
		desc := item.ShortDescription
		price := item.Price
		if len(strings.TrimSpace(desc))%3 == 0 {
			point += int(math.Ceil(0.2 * price))
		}
	}
	return point
}

func getPurchaseDayPoint(receipt Receipt) int {
	date := receipt.PurchaseDate
	day := date.Day()
	point := 0
	if day%2 == 1 {
		point = 6
	}
	return point
}

func getPurchaseTimePoint(receipt Receipt) int {
	point := 0
	time := receipt.PurchaseTime
	if time.Hour() >= 14 && time.Hour() <= 16 {
		point = 10
	}
	return point
}
