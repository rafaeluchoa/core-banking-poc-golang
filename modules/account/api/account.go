// @title Account API
// @version 1.0
// @description Account API
// @host localhost:8080
// @BasePath /
package api

const (
	API_ACCOUNT = "/account"
)

type Account struct {
	Id         string
	CustomerId string
}

// List

// @Description Structure that contains the parameters for listing accounts.
// @Param customerId query string true "Customer ID"
type AccountListReq struct {
	CustomerId string `json:"customerId"`
}

type AccountListRes struct {
	Response
	Accounts []Account `json:"accounts"`
}

// Create

type AccountCreateReq struct {
	CustomerId string `json:"customerId"`
}

type AccountCreateRes struct {
	Response
	AccountId string `json:"accountId"`
}
