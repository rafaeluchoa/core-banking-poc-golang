package account

import (
	"nk/account/api"
	"nk/account/internal/repo"
	"nk/account/test/e2e"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAccountSuccess(t *testing.T) {
	// Setup
	e2e.Setup()

	// Arrange
	customerId := repo.UUID()

	req := &api.AccountListReq{
		CustomerId: customerId,
	}

	// Act
	// TODO: get
	res := e2e.Post[api.AccountListRes](api.API_ACCOUNT, req)

	// Assert
	assert.NotNil(t, res.Accounts)
}
