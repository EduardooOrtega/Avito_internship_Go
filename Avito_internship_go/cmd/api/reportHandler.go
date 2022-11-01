package main

import (
	"Avito_internship_go/internal/data"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
)

func (app *application) getReservTempHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParam(r)

	if err != nil || id < 1 {
		app.logger.Println(err)
		err = app.writeJSON(w, http.StatusOK, envelope{"reserv": nil}, nil)
		return
	}

	reserv, err := app.models.Report.GetReserv(id)
	if err != nil {
		app.logger.Println(err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"reserv": reserv}, nil)

	if err != nil {
		app.logger.Println(err)
	}
}

// функция подтверждения пэймента
func (app *application) SuccsesPayHandler(w http.ResponseWriter, r *http.Request) {
	reservId, err := app.readIdParam(r)
	if err != nil || reservId < 1 {
		app.logger.Println(err)
		return
	}

	reserv, err := app.models.Report.GetReserv(reservId)

	if reserv == nil {
		return
	}

	if err != nil {
		app.logger.Println(err)
	}

	Service, err := app.models.Service.Get(reserv.ServiceId)
	if err != nil || reservId < 1 {
		app.logger.Println(err)
		return
	}
	Account, err := app.models.Account.Get(reserv.AccountId)
	if err != nil || reservId < 1 {
		app.logger.Println(err)
		return
	}
	app.models.Report.CreatePayment(reserv.AccountId, reserv.ServiceId, reservId, Account.AccountReservedCash-Service.ServicePrice)
	if err != nil {
		app.logger.Println(err)
	}

}

// функция получения платежа
func (app *application) getPaymentHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParam(r)

	if err != nil || id < 1 {
		app.logger.Println(err)
		return
	}

	report, err := app.models.Report.Get(id)
	if err != nil {
		app.logger.Println(err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"report": report}, nil)

	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) getReportHandler(w http.ResponseWriter, r *http.Request) {

	year, err := app.readYearParam(r)
	if err != nil {
		log.Println(err)
		return
	}

	month, err := app.readMonthParam(r)
	if err != nil {
		log.Println(err)
		return
	}

	reports, err := app.models.Report.GetMonthlyReport(year, month)
	if err != nil {
		log.Println(err)
		return
	}

	fileName := fmt.Sprintf("report%d-%d.csv", year, month)
	filePath := "reports/" + fileName

	csvFile, err := os.Create(filePath)
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	err = csvwriter.WriteAll(reports)
	if err != nil {
		log.Println("writing to CSV file error")
		http.Error(w, "cannot write to CSV file", http.StatusInternalServerError)
		return
	}
	csvwriter.Flush()

	link := fmt.Sprintf("http://localhost:4000/file/%s", fileName)
	err = app.writeJSON(w, http.StatusOK, envelope{"link": link}, nil)

	if err != nil {
		log.Println(err)
	}

}

func (app *application) getFileOfReportHandler(w http.ResponseWriter, r *http.Request) {

	fileName, _ := app.readFileNameParam(r)
	filePath := "reports/" + fileName

	http.ServeFile(w, r, filePath)
}

func (app *application) getUserReportHistoryHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		data.Filters
	}

	qs := r.URL.Query()
	var err error
	input.Filters.Page = app.readIntFromQuery(qs, "page", 1)
	input.Filters.PageSize = app.readIntFromQuery(qs, "page_size", 5)
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

	userReportHistory, metadata, err := app.models.Report.GetUserHistory(id, input.Filters)

	if err != nil {
		log.Println(err)
		return
	}
	err = app.writeJSON(w, http.StatusOK,
		envelope{"metadata": metadata}, nil)
	err = app.writeJSON(w, http.StatusOK,
		envelope{"User Buying Services History": userReportHistory}, nil)

	if err != nil {
		log.Println(err)
	}

}
