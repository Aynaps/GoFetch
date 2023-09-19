package receiptService

import (
	"fmt"
	"math"
	"net/http"
	"src/main/src/models"
	"strings"
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
		ctx.JSON(http.StatusCreated, gin.H{"id": uuid.String()})
	} else {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
	}
}

func HandleGetRewardPoints(ctx *gin.Context) {
	userUUID := uuid.MustParse(ctx.Param("id"))

	if mockDataBase[userUUID] == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Unknown UUID"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"points": mockDataBase[userUUID]})
	}
}

func processReceipt(json models.Receipt) (uuid.UUID, error) {
	purchaseDate, purchaseTime, err := parseDateTime(json.PurchaseDate, json.PurchaseTime)
	if err != nil {
		return uuid.Nil, err
	}
	var pointSum = calculatePoints(json, purchaseDate, purchaseTime)
	var newUUID = uuid.New()
	mockDataBase[newUUID] = pointSum
	return newUUID, nil
}

func calculatePoints(r models.Receipt, purchaseDate *time.Time, purchaseTime *time.Time) int {

	var sum = 0

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
	sum += (len(r.Items) / 2) * 5
	for _, item := range r.Items {
		if (len(strings.TrimSpace(item.ShortDescription)) % 3) == 0 {
			sum += int(math.Ceil(item.Price * 0.2))
		}
	}
	if (purchaseDate.Day() % 2) != 0 {
		sum += 6
	}
	if purchaseTime.After(TWOPM) && purchaseTime.Before(FOURPM) {
		sum += 10
	}
	return sum
}

func parseDateTime(date string, clockTime string) (*time.Time, *time.Time, error) {
	//validate purchaseDate is yyyy-mm-dd
	parsedDate, err := time.Parse(time.DateOnly, date)

	if err != nil {
		return nil, nil, err
	}
	//validate purchaseTime is hh:mm
	time, err := time.Parse(HHMM, clockTime)

	if err != nil {
		return nil, nil, err
	}
	return &parsedDate, &time, nil
}
