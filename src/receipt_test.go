package main

import (
	"src/main/src/models"
	"src/main/src/receiptService"
	"testing"
)

/*
		type Receipt struct {
			Retailer     string  `json:"retailer" binding:"required"`
			PurchaseDate string  `json:"purchaseDate" binding:"required"`
			PurchaseTime string  `json:"purchaseTime" binding:"required"`
			Items        []*Item `json:"items" binding:"required"`
			Total        float64 `json:"total,string" binding:"required"`
		}
		{
	  "retailer": "Target",
	  "purchaseDate": "2022-01-01",
	  "purchaseTime": "13:01",
	  "items": [
	    {
	      "shortDescription": "Mountain Dew 12PK",
	      "price": "6.49"
	    },{
	      "shortDescription": "Emils Cheese Pizza",
	      "price": "12.25"
	    },{
	      "shortDescription": "Knorr Creamy Chicken",
	      "price": "1.26"
	    },{
	      "shortDescription": "Doritos Nacho Cheese",
	      "price": "3.35"
	    },{
	      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
	      "price": "12.00"
	    }
	  ],
	  "total": "35.35"
	}
*/
func TestPointsTotal(t *testing.T) {
	items := []models.Item{
		{ShortDescription: "Mountain Dew 12PK", Price: 6.49},
		{ShortDescription: "Emils Cheese Pizza", Price: 12.25},
		{ShortDescription: "Knorr Creamy Chicken", Price: 1.26},
		{ShortDescription: "Doritos Nacho Cheese", Price: 3.35},
		{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: 12.00},
	}
	var r = models.Receipt{Retailer: "Target", PurchaseDate: "2022-01-01", PurchaseTime: "13:01", Items: items, Total: 35.35}
	got := receiptService.AccumulatePoints(r)
	want := 28
	if got != want {
		t.Fatalf(`Wanted %d, Got %d`, want, got)
	}
}
