package account

import (
	"nk/account/internal/api"
	"nk/account/test/e2e"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateAccount(t *testing.T, customerId string) string {
	// Setup
	e2e.Setup()

	req := &api.AccountCreateReq{
		CustomerId: customerId,
	}

	// Act
	res := e2e.Post[api.AccountCreateRes](api.API_ACCOUNT, req)

	// Assert
	assert.Empty(t, res.Message)
	assert.NotEmpty(t, res.AccountId)

	return res.AccountId
}
