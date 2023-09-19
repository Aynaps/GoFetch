package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"src/main/src/models"
	"src/main/src/receiptService"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
)

type PostResponse struct {
	Id uuid.UUID
}
type GetResponse struct {
	Points int
}

func setUpRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/:id/points", receiptService.HandleGetRewardPoints)
	router.POST("/receipts/process", receiptService.HandleProcessReceipt)
	return router
}

func TestPostRoute(t *testing.T) {
	r := setUpRouter()
	items := []models.Item{
		{ShortDescription: "Gatorade", Price: 2.25},
		{ShortDescription: "Gatorade", Price: 2.25},
		{ShortDescription: "Gatorade", Price: 2.25},
		{ShortDescription: "Gatorade", Price: 2.25},
	}
	var receipt = models.Receipt{Retailer: "M&M Corner Market", PurchaseDate: "2022-03-20", PurchaseTime: "14:33", Items: items, Total: 9.00}
	json, _ := json.Marshal(receipt)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(json))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEqual(t, w.Body, nil)
}

func TestGetRoute(t *testing.T) {
	r := setUpRouter()
	items := []models.Item{
		{ShortDescription: "Mountain Dew 12PK", Price: 6.49},
		{ShortDescription: "Emils Cheese Pizza", Price: 12.25},
		{ShortDescription: "Knorr Creamy Chicken", Price: 1.26},
		{ShortDescription: "Doritos Nacho Cheese", Price: 3.35},
		{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: 12.00},
	}
	var receipt = models.Receipt{Retailer: "Target", PurchaseDate: "2022-01-01", PurchaseTime: "13:01", Items: items, Total: 35.35}
	jsonVal, _ := json.Marshal(receipt)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonVal))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var postObj PostResponse
	json.Unmarshal(w.Body.Bytes(), &postObj)

	assert.NotEqual(t, postObj, nil)
	idURI := "/" + postObj.Id.String() + "/points"
	fmt.Println(postObj)
	w = httptest.NewRecorder()

	getReq, _ := http.NewRequest("GET", idURI, nil)
	r.ServeHTTP(w, getReq)

	var getObj GetResponse
	fmt.Println(w.Body)
	json.Unmarshal(w.Body.Bytes(), &getObj)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, getObj.Points, 28)
}
