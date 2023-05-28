package main

import (
	"encoding/json"
	"gses2-btc-app/services"
	"gses2-btc-app/storage"
	"gses2-btc-app/utils"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/rate", GetRateHandler)
	http.HandleFunc("/api/subscribe", SubscribeHandler)
	//http.HandleFunc("/api/sendEmails", SendEmailsHandler)

	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// SubscribeHandler is function for handle query for adding new email for subscription
func SubscribeHandler(writer http.ResponseWriter, request *http.Request) {
	reqString := request.URL
	email := reqString.Query().Get("email")

	// additional check if user input is an email
	// return status 405 with a message about non-valid email
	if !utils.IsEmailValid(email) {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(writer).Encode("Email is not valid, try again!")
		if err != nil {
			return
		}
	}

	// check for email in emails list
	if !storage.CheckForEmail(email) {
		storage.AddEmail(email)
	} else {
		// return code 409 if email is present in subscription list
		writer.WriteHeader(http.StatusConflict)
		writer.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(writer).Encode("Email is in subscriptions list!")
		if err != nil {
			return
		}
	}
}

func GetRateHandler(writer http.ResponseWriter, request *http.Request) {
	price, err := services.GetCurrentPrice(utils.FromCurr, utils.ToCurr)
	if price == 0 {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(price)
	if err != nil {
		return
	}
}
