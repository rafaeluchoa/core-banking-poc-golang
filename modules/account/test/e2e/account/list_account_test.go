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
	customerID := repo.UUID()
	expectedAccountID := CreateAccount(t, customerID)

	req := &api.AccountListReq{
		CustomerID: customerID,
	}

	// Act
	res := e2e.Get[api.AccountListRes](api.APIAccount, req)

	// Assert
	assert.NotNil(t, res.Accounts)
	assert.Equal(t, res.Accounts[0].ID, expectedAccountID)
}
