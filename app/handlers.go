package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vietbui1502/RestAPIGolang/dto"
	"github.com/vietbui1502/RestAPIGolang/logger"
	"github.com/vietbui1502/RestAPIGolang/service"
)

// type Customer struct {
// 	Name    string `json:"full_name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	Zipcode string `json:"zip_code" xml:"zip_code"`
// }

type CustomerHandlers struct {
	customerService service.CustomerService
	accountSerivce  service.AccountService
}

func (ch *CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.customerService.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.customerService.GetCustomer(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		logger.Error("Customer handler get customer error" + err.Error())
		fmt.Println(w, err.Error())
	} else {
		if r.Header.Get("Content-Type") == "application/xml" {
			w.Header().Add("Content-Type", "application/xml")
			xml.NewEncoder(w).Encode(customer)
		} else {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(customer)
		}
	}

}

func (ch *CustomerHandlers) creatAccount(w http.ResponseWriter, r *http.Request) {
	var mux = mux.Vars(r)
	customer_id := mux["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		logger.Error("Account handler creatAccount error " + err.Error())
	} else {
		request.CustomerId = customer_id
		account, err := ch.accountSerivce.NewAccount(request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error("Account handler creatAccount error " + err.Error())
		} else {
			w.WriteHeader(http.StatusCreated)
			logger.Info("AccountID: " + account.AccountId)
		}
	}
}
