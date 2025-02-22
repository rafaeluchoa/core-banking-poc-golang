// @title Account API
// @version 1.0
// @description Account API
// @host localhost:8080
// @BasePath /api/v1
package api

const (
	API_ACCOUNT = "/api/v1/account"
)

type Account struct {
	Id         string
	CustomerId string
}

// List

// @Description Request
// @Param customerId query string true "Customer ID"
type AccountListReq struct {
	CustomerId string `json:"customerId"`
}

// @Description Response
// @Param accounts body Account true "Account"
type AccountListRes struct {
	Response
	Accounts []Account `json:"accounts"`
}

// Create

// @Description Request
// @Param customerId body string true "Customer ID"
type AccountCreateReq struct {
	CustomerId string `json:"customerId"`
}

// @Description Response
// @Param accountId body string true "Account ID"
type AccountCreateRes struct {
	Response
	AccountId string `json:"accountId"`
}
