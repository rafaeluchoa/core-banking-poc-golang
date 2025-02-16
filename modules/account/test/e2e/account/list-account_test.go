package account

import (
	"nk/account/api"
	"nk/account/test/e2e"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAccountSuccess(t *testing.T) {
	// Setup
	e2e.Setup()

	// Arrange
	customerId := e2e.UUID()

	req := &api.AccountListReq{
		CustomerId: customerId,
	}

	// Act
	res := e2e.Post[api.AccountCreateRes](api.API_ACCOUNT, req)

	// Assert
	assert.NotNil(t, res.AccountId)
}
