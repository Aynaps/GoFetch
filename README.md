# Introduction

A take home assessment by Fetch Rewards written in Golang 
using Gin web framework.


# Pre-Requisites
Version go 1.21.1

# How to Run
1. Pull down the repo with ssh or https
2. run ```go get .``` to install the dependencies in the go.mod file
2. cd into the src directory
3. run the main.go file with `go run main.go`
4. now the two API endpoints `/receipts/process` and `:id/points` are open for requests listening on port 8080


# Example CURL requests

example CURL for `/receipts/process

`curl --location 'localhost:8080/receipts/process' \
--header 'Content-Type: application/json' \
--data '{
  "retailer": "Target",
  "purchaseDate": "2006-01-02",
  "purchaseTime":"16:01",
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
}'`

example curl for /:id/points

`curl --location 'localhost:8080/{uuid-goes-here}/points'`


# Running Test Cases
I've written two test files `main_test.go` which tests the http requests

and `receipt_service_test.go` which tests the helper functions i've written

to run all the files be in the parent directory with the main file and run 
`go test -v ./...` this will run through every directory and run each test file

to run a  **specific** `cd` into its respective directory with the file inside
and run with `go test receipt_service_test.go` or `go test main_test.go` 