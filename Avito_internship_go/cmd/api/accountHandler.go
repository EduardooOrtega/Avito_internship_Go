package main

import (
	"fmt"
	"net/http"
)

func (app *application) createAccountHandler(w http.ResponseWriter, r *http.Request) {
	app.models.Account.Create()
}

func (app *application) getAccountHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParam(r)

	if err != nil || id < 1 {
		app.logger.Println(err)
		return
	}

	account, err := app.models.Account.Get(id)
	if err != nil {
		app.logger.Println(err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)

	if err != nil {
		app.logger.Println(err)
	}
}

// функция перевода денег из кошелька в резерв
func (app *application) depositReservAccountHandler(w http.ResponseWriter, r *http.Request) {
	id_account, id_service, err := app.readDepositReservParam(r)

	if err != nil {
		app.logger.Println(err)
		return
	}

	account, err := app.models.Account.Get(id_account)
	service, err := app.models.Service.Get(id_service)
	var price = service.ServicePrice
	if err != nil {
		app.logger.Println(err)
		return
	}

	if account == nil || service.ServiceId == 0 {
		app.logger.Println("Service or account not found")
		return
	}
	account.AccountId = id_account
	if account.AccountCash-price < 0 {
		app.logger.Println("not enough money")
		return
	}
	account.AccountCash -= price
	account.AccountReservedCash += price

	err = app.models.Account.UpdateFull(account)
	if err != nil {
		app.logger.Println(err)
	}

	app.models.Report.CreateReserv(id_account, id_service)
	if err != nil {
		app.logger.Println("Can not create temp_report")
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)
	err = app.writeJSON(w, http.StatusOK, envelope{"service": service}, nil)

	if err != nil {
		app.logger.Println(err)
	}
}

// функция добавления денег на аккаунт
func (app *application) addDepositAccountHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	cash, err := app.readCashParam(r)

	if err != nil {
		app.logger.Println(err)
		return
	}

	account, err := app.models.Account.Get(id)
	if err != nil || account == nil {
		app.models.Account.CreateId(id)
		app.logger.Println(err)
	}
	account, err = app.models.Account.Get(id)

	account.AccountId = id
	account.AccountCash += cash

	err = app.models.Account.Update(account)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)
	if err != nil {
		app.logger.Println(err)
	}
	app.models.Transaction.Create(account.AccountId, account.AccountId, 1, cash)
}

// функция перевода денег с аккаунта на аккаунт
func (app *application) transferAccountHandler(w http.ResponseWriter, r *http.Request) {
	FromId, err := app.readIdParam(r)
	ToId, err := app.readToIdParam(r)
	cash, err := app.readCashParam(r)
	if cash < 0 {
		return
	}

	accountFrom, err := app.models.Account.Get(FromId)
	if err != nil {
		app.logger.Println(err)
		return
	}

	accountTo, err := app.models.Account.Get(ToId)
	if err != nil {
		app.logger.Println(err)
		return
	}
	var newCash float64 = accountFrom.AccountCash - cash
	if newCash < 0 {
		return
	} else {
		fmt.Printf("%2f", newCash)
		accountFrom.AccountCash = newCash
		newCash = accountTo.AccountCash + cash
		accountTo.AccountCash = newCash
		app.models.Account.Update(accountFrom)
		app.models.Account.Update(accountTo)
		app.models.Transaction.Create(accountFrom.AccountId, accountTo.AccountId, 3, cash)
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"accountFrom": accountFrom}, nil)
	err = app.writeJSON(w, http.StatusOK, envelope{"accountTo": accountTo}, nil)
}

// функция снятия денег с аккаунта
func (app *application) withDrawaDepositAccountHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	cash, err := app.readCashParam(r)

	if err != nil {
		app.logger.Println(err)
		return
	}

	account, err := app.models.Account.Get(id)
	if err != nil {
		app.logger.Println(err)
		return
	}
	if account.AccountId == 0 {
		app.logger.Println("Account not found")
		return
	}
	if account.AccountCash-cash < 0 {
		return
	}
	account.AccountId = id
	account.AccountCash -= cash

	err = app.models.Account.Update(account)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)
	if err != nil {
		app.logger.Println(err)
	}
	app.models.Transaction.Create(account.AccountId, account.AccountId, 2, cash)
}
