package models

import "fmt"

type Receipt struct {
	Retailer     string  `json:"retailer" binding:"required"`
	PurchaseDate string  `json:"purchaseDate" binding:"required"`
	PurchaseTime string  `json:"purchaseTime" binding:"required"`
	Items        []Item  `json:"items" binding:"required"`
	Total        float64 `json:"total,string" binding:"required"`
}

func (r Receipt) String() string {
	fmt.Println("Items:")
	for i := 0; i < len(r.Items); i++ {
		fmt.Println(r.Items[i])
	}
	return fmt.Sprintf("retailer: %s\n purchaseDate: %s\n purchaseTime: %s\n Total:%f", r.Retailer, r.PurchaseDate, r.PurchaseTime, r.Total)
}

type Item struct {
	ShortDescription string  `json:"shortDescription" binding:"required"`
	Price            float64 `json:"price,string" binding:"required"`
}

func (i Item) String() string {
	return fmt.Sprintf("desc: %s, price: %f", i.ShortDescription, i.Price)
}
