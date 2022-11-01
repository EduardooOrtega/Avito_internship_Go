package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/service/get/:id", app.serviceGetHandler)
	router.HandlerFunc(http.MethodPost, "/service/create/:name/:price", app.createServiceHandler)

	router.HandlerFunc(http.MethodGet, "/reserv/get/:id", app.getReservTempHandler)

	//id = reserv_id (create report based on reserv)
	router.HandlerFunc(http.MethodPut, "/report/create/:id", app.SuccsesPayHandler)
	router.HandlerFunc(http.MethodGet, "/report/getID/:id", app.getPaymentHandler)
	router.HandlerFunc(http.MethodGet, "/report/get/:year/:month", app.getReportHandler)
	router.HandlerFunc(http.MethodGet, "/file/:filename", app.getFileOfReportHandler)

	router.HandlerFunc(http.MethodPost, "/account/create", app.createAccountHandler)
	router.HandlerFunc(http.MethodGet, "/account/getId/:id", app.getAccountHandler)
	router.HandlerFunc(http.MethodPut, "/account/add/:id/:account_cash", app.addDepositAccountHandler)
	router.HandlerFunc(http.MethodPut, "/account/transfer/:id/:ToId/:account_cash", app.transferAccountHandler)
	router.HandlerFunc(http.MethodPut, "/account/withdrawal/:id/:account_cash", app.withDrawaDepositAccountHandler)
	router.HandlerFunc(http.MethodPut, "/account/reserv/:id_account/:id_service", app.depositReservAccountHandler)

	router.HandlerFunc(http.MethodGet, "/transaction/get/:id", app.getTransactionHandler)

	router.HandlerFunc(http.MethodGet, "/history/transactions/:id", app.getTransactionHistoryHandler)
	router.HandlerFunc(http.MethodGet, "/history/user/:id", app.getUserReportHistoryHandler)

	return router
}
