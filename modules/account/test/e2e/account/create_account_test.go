package account

import (
	"nk/account/internal/api"
	"nk/account/internal/repo"
	"nk/account/test/e2e"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccountSuccess(t *testing.T) {
	// Setup
	e2e.Setup()

	// Arrange
	customerId := repo.UUID()

	req := &api.AccountCreateReq{
		CustomerId: customerId,
	}

	// Act
	res := e2e.Post[api.AccountCreateRes](api.API_ACCOUNT, req)

	// Assert
	assert.Empty(t, res.Message)
	assert.NotEmpty(t, res.AccountId)
}
