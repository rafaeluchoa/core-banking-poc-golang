package account

import (
	"nk/account/api"
	"nk/account/test/e2e"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateAccount(t *testing.T, customerID string) string {
	// Setup
	e2e.Setup()

	req := &api.AccountCreateReq{
		CustomerID: customerID,
	}

	// Act
	res := e2e.Post[api.AccountCreateRes](api.APIAccount, req)

	// Assert
	assert.Empty(t, res.Message)
	assert.NotEmpty(t, res.AccountID)

	return res.AccountID
}
