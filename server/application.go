package server

import (
	"encoding/json"
	"net/http"

	"github.com/leonmaia/newmotion-golang/parser"
)

type Application struct{}

func (a *Application) TransactionHandler(w http.ResponseWriter, req *http.Request) {
	var decoder = json.NewDecoder(req.Body)
	transaction := new(parser.Transaction)
	err := decoder.Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	errValidation := transaction.IsValid()
	if err != nil {
		http.Error(w, errValidation.Error(), http.StatusInternalServerError)
		return
	}
}
