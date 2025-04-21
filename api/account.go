package api

import (
	"cushonTechTest/models"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"gorm.io/gorm"
)

type CurrencyType string

const (
	GBP CurrencyType = "GBP"
	USD CurrencyType = "USD"
)

type AccountResponse struct {
	ID      string `json:"id"`
	Balance string `json:"balance"`
	Fund    string `json:"fund,omitempty"`
	User    string `json:"user,omitempty"`
}

func (api *API) getAccount(writer http.ResponseWriter, request *http.Request) {
	slog.Info("starting getAccount process")

	account, err := api.findAccount(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	accountResponse := AccountResponse{
		ID:      account.ID.String(),
		Fund:    *account.Fund.Name,
		User:    fmt.Sprintf("%s %s", *account.User.FirstName, *account.User.LastName),
		Balance: formatCurrency(*account.Balance, GBP),
	}

	json.NewEncoder(writer).Encode(accountResponse)
}

func (api *API) findAccount(request *http.Request) (*models.Account, error) {
	fundID := request.URL.Query().Get("fundId")
	userID := request.URL.Query().Get("userId")

	var account models.Account
	err := api.DB.
		Preload("User.Type").
		Preload("Fund").
		Where("fund_id = ? AND user_id = ?", fundID, userID).
		First(&account).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("account not found")
		}
		return nil, err
	}

	return &account, nil
}

func formatCurrency(amount int, currencyType CurrencyType) string {
	switch currencyType {
	case GBP:
		return fmt.Sprintf("£%.2f", float64(amount)/100)
	case USD:
		return fmt.Sprintf("$%.2f", float64(amount)/100)
	default:
		return fmt.Sprintf("%.2f %s", float64(amount)/100, currencyType)
	}
}
