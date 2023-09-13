package receiptService

import (
	"fmt"
	"math"
	"net/http"
	"src/main/src/models"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var mockDataBase = make(map[uuid.UUID]int)

const HHMM = "15:04"

var TWOPM, _ = time.Parse(HHMM, "14:00")
var FOURPM, _ = time.Parse(HHMM, "16:00")

func HandleProcessReceipt(ctx *gin.Context) {
	var json models.Receipt

	if err := ctx.ShouldBindJSON(&json); err == nil {
		uuid, parseError := processReceipt(json)
		if parseError != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "purchaseDate must be yyyy-mm-dd and purchaseTime must be HH:MM"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": uuid.String()})
	} else {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
	}
}

func HandleGetRewardPoints(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "not yet implemented",
	})
}

func processReceipt(json models.Receipt) (uuid.UUID, error) {
	fmt.Println(json.String())
	//validate purchaseDate is yyyy-mm-dd
	date, err := time.Parse(time.DateOnly, json.PurchaseDate)
	fmt.Print(date)
	if err != nil {
		return uuid.Nil, err
	}
	//validate purchaseTime is hh:mm
	time, err := time.Parse(HHMM, json.PurchaseTime)
	fmt.Print(time)
	if err != nil {
		return uuid.Nil, err
	}
	var pointSum = AccumulatePoints(json)
	fmt.Printf("points: %d\n", pointSum)
	return uuid.New(), nil
}

func AccumulatePoints(r models.Receipt) int {
	var sum = 0
	/*
		1. One point for every alphanumeric character in the retailer name.
		2. 50 points if the total is a round dollar amount with no cents.
		3. 25 points if the total is a multiple of 0.25.
		4. 5 points for every two items on the receipt.
		5. If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
		6. 6 points if the day in the purchase date is odd.
		7. 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	*/

	for _, char := range r.Retailer {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			sum++
		}
	}
	if r.Total == math.Trunc(r.Total) {
		sum += 50
	}
	if math.Mod(r.Total, 0.25) == 0 {
		sum += 25
	}
	return sum
}
