// @title Account API
// @version 1.0
// @description Account API
// @host localhost:8080
// @BasePath /api/v1
package api

const (
	APIAccount = "/api/v1/account"
)

type Account struct {
	ID         string
	CustomerID string
}

// List

// @Description Request
// @Param customerId query string true "Customer ID"
type AccountListReq struct {
	CustomerID string `json:"customerId"`
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
	CustomerID string `json:"customerId"`
}

// @Description Response
// @Param accountId body string true "Account ID"
type AccountCreateRes struct {
	Response
	AccountID string `json:"accountId"`
}
