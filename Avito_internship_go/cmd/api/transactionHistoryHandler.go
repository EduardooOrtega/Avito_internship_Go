package main

import (
	"Avito_internship_go/internal/data"
	"log"
	"net/http"
)

func (app *application) getTransactionHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParam(r)

	if err != nil || id < 1 {
		app.logger.Println(err)
		return
	}

	transaction, err := app.models.Transaction.Get(id)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"transaction": transaction}, nil)

	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) getTransactionHistoryHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		data.Filters
	}

	qs := r.URL.Query()
	var err error
	input.Filters.Page = app.readIntFromQuery(qs, "page", 1)
	input.Filters.PageSize = app.readIntFromQuery(qs, "page_size", 1)
	input.Filters.Sort, err = app.readStringFromQuery(qs, "sort", "")

	if err != nil {
		err = app.writeJSON(w, http.StatusBadRequest, envelope{"error": error.Error(err)}, nil)
		return
	}

	id, err := app.readIdParam(r)
	if err != nil || id < 1 {
		log.Println(err)
		return
	}

	tranHistory, metadata, err := app.models.TransactionHistory.Get(id, input.Filters)

	if err != nil {
		log.Println(err)
		return
	}
	err = app.writeJSON(w, http.StatusOK,
		envelope{"metadata": metadata}, nil)
	err = app.writeJSON(w, http.StatusOK,
		envelope{"User Transaction History": tranHistory}, nil)

	if err != nil {
		log.Println(err)
	}

}
