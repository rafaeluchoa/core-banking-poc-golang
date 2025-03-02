package account

import (
	"nk/account/api"
	"nk/account/internal/repo"
	"nk/account/test/e2e"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccountSuccess(t *testing.T) {
	// Setup
	e2e.Setup()

	// Arrange
	customerID := repo.UUID()

	req := &api.AccountCreateReq{
		CustomerID: customerID,
	}

	// Act
	res := e2e.Post[api.AccountCreateRes](api.APIAccount, req)

	// Assert
	assert.Empty(t, res.Message)
	assert.NotEmpty(t, res.AccountID)
}
