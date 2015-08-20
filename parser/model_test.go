package parser

import (
	"encoding/json"
	"time"

	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestTransactionCreationWithSuccess(c *C) {
	jsonprep := `{"customerId": "john","startTime": "2014-10-28T09:34:17Z","endTime": "2014-10-28T16:45:13Z","volume": 32.03 }`
	jsonStr := []byte(jsonprep)
	var transaction Transaction
	startDate := time.Date(2014, 10, 28, 9, 34, 17, 0, time.UTC)
	endDate := time.Date(2014, 10, 28, 16, 45, 13, 0, time.UTC)

	json.Unmarshal(jsonStr, &transaction)

	c.Assert(transaction.IsValid(), Equals, nil)
	c.Assert(transaction.CustomerID, Equals, "john")
	c.Assert(transaction.StartTime, Equals, startDate)
	c.Assert(transaction.EndTime, Equals, endDate)
}
func (s *TestSuite) TestTransactionCreationWithErrors(c *C) {
	jsonprep := `{"customerId": "","startTime": "2014-10-28T09:34:17Z","endTime": "2014-10-28T16:45:13Z","volume": 32.03 }`
	jsonStr := []byte(jsonprep)
	var transaction Transaction

	json.Unmarshal(jsonStr, &transaction)

	c.Assert(transaction.IsValid(), ErrorMatches, "Customer ID should not be empty")
}
