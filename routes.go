package main

import (
	"exercise1/db"
	"exercise1/rates"
	"net/http"

	"github.com/gorilla/mux"
)

func routes(db db.DBInterface) {
	rateRepository := rates.NewRepository(db.GetDB())
	rateUsecase := rates.NewUsecase(rateRepository)
	ratesHandler := rates.NewHandler(rateUsecase)

	r := mux.NewRouter()
	r.HandleFunc("/lastest-rate", ratesHandler.GetRates)
	r.HandleFunc("/{date}", ratesHandler.GetRates)
	r.HandleFunc("/average-currency", ratesHandler.GetRates)
	http.ListenAndServe(":8000", r)
}
