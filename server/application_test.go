package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestTransactionHandler(c *C) {
	jsonprep := `{"customerId": "john","startTime": "2014-10-28T09:34:17Z","endTime": "2014-10-28T16:45:13Z","volume": 32.03 }`

	jsonStr := []byte(jsonprep)
	req, _ := http.NewRequest("POST", "/api/transactions", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	app := &Application{}
	app.TransactionHandler(response, req)

	c.Assert(response.Code, Equals, http.StatusOK)
}

func (s *TestSuite) TestTransactionHandlerJsonWithError(c *C) {
	jsonprep := `{"customerId": "","startTime": "2014-10-28T09:34:17Z","endTime": "2014-10-28T16:45:13Z","volume": 32.03 }`
	jsonStr := []byte(jsonprep)

	req, _ := http.NewRequest("POST", "/api/transactions", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	app := &Application{}
	app.TransactionHandler(response, req)

	c.Assert(response.Code, Equals, http.StatusBadRequest)
}

func (s *TestSuite) TestTransactionHandlerWithInvalidBody(c *C) {
	jsonprep := "invalid body"
	jsonStr := []byte(jsonprep)

	req, _ := http.NewRequest("POST", "/api/transactions", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	app := &Application{}
	app.TransactionHandler(response, req)

	c.Assert(response.Code, Equals, http.StatusInternalServerError)
}
