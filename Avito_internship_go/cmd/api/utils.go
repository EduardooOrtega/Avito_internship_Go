package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type envelope map[string]interface{}

func (app *application) readServiceParam(r *http.Request) (string, float64, error) {

	params := httprouter.ParamsFromContext(r.Context())
	name := params.ByName("name")
	price, err := strconv.ParseFloat(params.ByName("price"), 64)

	if err != nil || price < 0 {
		return "", 0, errors.New("invalid parameter")
	}

	return name, price, nil
}

func (app *application) readIdParam(r *http.Request) (int64, error) {

	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) readToIdParam(r *http.Request) (int64, error) {

	params := httprouter.ParamsFromContext(r.Context())
	ToId, err := strconv.ParseInt(params.ByName("ToId"), 10, 64)

	if err != nil || ToId < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return ToId, nil
}

func (app *application) readDepositReservParam(r *http.Request) (int64, int64, error) {

	params := httprouter.ParamsFromContext(r.Context())
	id_account, err := strconv.ParseInt(params.ByName("id_account"), 10, 64)
	id_service, err := strconv.ParseInt(params.ByName("id_service"), 10, 64)

	if err != nil || id_account < 1 || id_service < 1 {
		return 0, 0, errors.New("invalid parameter")
	}

	return id_account, id_service, nil
}

func (app *application) readYearParam(r *http.Request) (int64, error) {

	params := httprouter.ParamsFromContext(r.Context())
	year, err := strconv.ParseInt(params.ByName("year"), 10, 64)

	if err != nil || year < 1 {
		return 0, errors.New("invalid year parameter")
	}

	return year, nil
}

func (app *application) readMonthParam(r *http.Request) (int64, error) {

	params := httprouter.ParamsFromContext(r.Context())
	month, err := strconv.ParseInt(params.ByName("month"), 10, 64)

	if err != nil || month < 1 || month > 12 {
		return 0, errors.New("invalid month parameter")
	}

	return month, nil
}

func (app *application) readCashParam(r *http.Request) (float64, error) {

	params := httprouter.ParamsFromContext(r.Context())

	accountCash, err := strconv.ParseFloat(params.ByName("account_cash"), 64)
	if err != nil || accountCash < 0 {
		return 0, errors.New("invalid id parameter")
	}

	return accountCash, nil
}

func (app *application) readFileNameParam(r *http.Request) (string, error) {

	params := httprouter.ParamsFromContext(r.Context())
	filename := params.ByName("filename")

	return filename, nil
}
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {

	js, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {

		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	return nil
}

func (app *application) readStringFromQuery(qs url.Values, key string, defaultValue string) (string, error) {
	s := qs.Get(key)

	if s == "" {
		return defaultValue, nil
	}

	if s != "date" && s != "sum" && s != "-date" && s != "-sum" {
		return "", errors.New("Invalid sort parametrs")
	}

	return s, nil
}

func (app *application) readIntFromQuery(qs url.Values, key string, defaultValue int) int {

	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	//parseInt
	i, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}

	if i < 1 {
		return defaultValue
	}

	return i
}

func (app *application) readFloatFromQuery(qs url.Values, key string, defaultValue float64) float64 {

	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultValue
	}

	return i
}
