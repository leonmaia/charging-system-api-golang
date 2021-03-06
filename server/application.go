package server

import (
	"encoding/json"
	"net/http"

	"github.com/leonmaia/newmotion-golang/keydb"
	"github.com/leonmaia/newmotion-golang/parser"
)

type Application struct {
	DB *keydb.DB
}

func (a *Application) TransactionHandler(w http.ResponseWriter, req *http.Request) {
	var decoder = json.NewDecoder(req.Body)
	transaction := new(parser.Transaction)
	if err := decoder.Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if errValidation := transaction.IsValid(); errValidation != nil {
		http.Error(w, errValidation.Error(), http.StatusBadRequest)
		return
	}

	a.DB.Set(transaction.CustomerID, transaction)
	w.WriteHeader(http.StatusCreated)
}

func (a *Application) OverviewHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
