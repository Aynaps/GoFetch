package receiptService

import (
	"src/main/src/models"
	"testing"
)

func TestExample1(t *testing.T) {
	items := []models.Item{
		{ShortDescription: "Mountain Dew 12PK", Price: 6.49},
		{ShortDescription: "Emils Cheese Pizza", Price: 12.25},
		{ShortDescription: "Knorr Creamy Chicken", Price: 1.26},
		{ShortDescription: "Doritos Nacho Cheese", Price: 3.35},
		{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: 12.00},
	}
	var r = models.Receipt{Retailer: "Target", PurchaseDate: "2022-01-01", PurchaseTime: "13:01", Items: items, Total: 35.35}
	date, time, _ := parseDateTime(r.PurchaseDate, r.PurchaseTime)
	got := calculatePoints(r, date, time)
	want := 28
	if got != want {
		t.Fatalf(`Wanted %d, Got %d`, want, got)
	}
}

func TestExample2(t *testing.T) {
	items := []models.Item{
		{ShortDescription: "Gatorade", Price: 2.25},
		{ShortDescription: "Gatorade", Price: 2.25},
		{ShortDescription: "Gatorade", Price: 2.25},
		{ShortDescription: "Gatorade", Price: 2.25},
	}
	var r = models.Receipt{Retailer: "M&M Corner Market", PurchaseDate: "2022-03-20", PurchaseTime: "14:33", Items: items, Total: 9.00}
	date, time, _ := parseDateTime(r.PurchaseDate, r.PurchaseTime)
	got := calculatePoints(r, date, time)
	want := 109
	if got != want {
		t.Fatalf(`Wanted %d, Got %d`, want, got)
	}
}
func TestInvalidDateTime(t *testing.T) {
	items := []models.Item{
		{ShortDescription: "Gatorade", Price: 2.25},
		{ShortDescription: "Gatorade", Price: 2.25},
		{ShortDescription: "Gatorade", Price: 2.25},
		{ShortDescription: "Gatorade", Price: 2.25},
	}
	var r = models.Receipt{Retailer: "M&M Corner Market", PurchaseDate: "2022-03-20:12:12", PurchaseTime: "14:33", Items: items, Total: 9.00}
	_, _, err := parseDateTime(r.PurchaseDate, r.PurchaseTime)

	if err == nil {
		t.Fatalf(`Date time parser was suppose to fail on %s and %s`, r.PurchaseDate, r.PurchaseTime)
	}
}
