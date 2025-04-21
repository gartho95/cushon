package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDepositType(t *testing.T) {
	api := &API{}
	t.Run("should return RetailDeposit for userType 'retail'", func(t *testing.T) {
		strategy, err := api.getDepositType("retail")
		assert.NoError(t, err)

		_, ok := strategy.(RetailDeposit)
		assert.True(t, ok, "Expected RetailDeposit type")
	})

	t.Run("should return EmployerDeposit for userType 'employer'", func(t *testing.T) {
		strategy, err := api.getDepositType("employer")
		assert.NoError(t, err)

		_, ok := strategy.(EmployerDeposit)
		assert.True(t, ok, "Expected EmployerDeposit type")
	})

	t.Run("should return error for unknown userType", func(t *testing.T) {
		strategy, err := api.getDepositType("other")

		assert.Nil(t, strategy)
		assert.Error(t, err)
		assert.Equal(t, "unknown user type", err.Error())
	})
}

func TestDecodeDepositRequestValidJSON(t *testing.T) {
	api := &API{}
	body := `{"fundBalances":[{"fundId":"abc","value":100}], "userId":"user1"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	responseRecorder := httptest.NewRecorder()

	result, err := api.decodeDepositRequest(responseRecorder, req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.UserID != "user1" {
		t.Errorf("Expected userID user1, got %s", result.UserID)
	}
	if len(result.FundBalances) != 1 {
		t.Errorf("Expected 1 fundBalance, got %d", len(result.FundBalances))
	}
	if result.FundBalances[0].FundID != "abc" || result.FundBalances[0].Value != 100 {
		t.Errorf("Unexpected fund balance values: %+v", result.FundBalances[0])
	}
}
