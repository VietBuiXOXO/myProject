package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/vietbui1502/RestAPIGolang/domain"
	"github.com/vietbui1502/RestAPIGolang/logger"
	"github.com/vietbui1502/RestAPIGolang/service"
)

func Start() {

	//router := http.NewServeMux()
	router := mux.NewRouter()

	//Creat DBclient
	dbClient := getDbClient()
	//Create customer service
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	//accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	ch := CustomerHandlers{customerService: service.NewCustomerService(customerRepositoryDb)}

	//Define route
	router.HandleFunc("/customer", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account", ch.creatAccount).Methods(http.MethodPost)

	//Starting server
	log.Fatal(http.ListenAndServe(":8088", router))
}

func getDbClient() *sqlx.DB {
	dbClient, err := sqlx.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
	if err != nil {
		logger.Error(err.Error())
	}
	// See "Important settings" section.
	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)
	return dbClient
}
