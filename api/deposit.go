package api

import (
	"cushonTechTest/models"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

type DepositRequest struct {
	FundBalances []FundBalance `json:"fundBalances"`
	UserID       string        `json:"userId"`
}

type FundBalance struct {
	FundID string `json:"fundId"`
	Value  int    `json:"value"`
}

func (api *API) getDepositType(userType string) (DepositStrategy, error) {
	switch userType {
	case "retail":
		return RetailDeposit{api: api}, nil
	case "employer":
		return EmployerDeposit{api: api}, nil
	default:
		return nil, errors.New("unknown user type")
	}
}

func (api *API) depositFunds(writer http.ResponseWriter, request *http.Request) {
	slog.Info("starting depositFunds process")

	depositRequest, err := api.decodeDepositRequest(writer, request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	slog.Info(fmt.Sprintf("processing request for userID %s and fundIDs %v", depositRequest.UserID, depositRequest.FundBalances))

	accounts, err := api.findAccountsFromDepositRequest(depositRequest)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	depositMap := api.getBalancesForFunds(depositRequest)
	transaction := api.DB.Begin()
	slog.Info("beginning transaction to handle deposits")

	for _, account := range accounts {
		if account.User.Type.Name == nil {
			http.Error(writer, "User type is missing", http.StatusInternalServerError)
			transaction.Rollback()
			return
		}

		depositAmount, ok := depositMap[account.FundID.String()]
		if !ok {
			http.Error(writer, fmt.Sprintf("No balance found for fund %s", account.FundID), http.StatusInternalServerError)
			transaction.Rollback()
			return
		}

		strategy, err := api.getDepositType(*account.User.Type.Name)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			transaction.Rollback()
			return
		}

		if err := strategy.Deposit(account.ID.String(), *account.Balance, depositAmount); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			transaction.Rollback()
			return
		}
	}
	slog.Info("committing transaction to handle deposits")
	fmt.Println(api.DB)
	transaction.Commit()
	fmt.Println(api.DB)
}
func (api *API) decodeDepositRequest(writer http.ResponseWriter, request *http.Request) (DepositRequest, error) {
	var depositRequestBody DepositRequest
	if err := json.NewDecoder(request.Body).Decode(&depositRequestBody); err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return DepositRequest{}, err
	}

	return depositRequestBody, nil
}

func (api *API) depositFundsToAccount(accountID string, newBalance int) error {

	if err := api.DB.Model(&models.Account{}).
		Where("id = ?", accountID).
		Update("balance", newBalance).Error; err != nil {
		return err
	}
	return nil
}

func (api *API) findAccountsFromDepositRequest(depositRequest DepositRequest) ([]models.Account, error) {
	var fundIDs []string
	for _, fundBalance := range depositRequest.FundBalances {
		fundIDs = append(fundIDs, fundBalance.FundID)
	}

	var accounts []models.Account

	if err := api.DB.
		Preload("User.Type").
		Preload("Fund").
		Where("user_id = ? AND fund_id IN ?", depositRequest.UserID, fundIDs).
		Find(&accounts).Error; err != nil {
		return nil, err
	}

	if len(accounts) == 0 {
		return nil, errors.New("no accounts found for user")
	}

	return accounts, nil

}

func (api *API) getBalancesForFunds(depositRequest DepositRequest) map[string]int {
	funds := make(map[string]int)
	for _, fund := range depositRequest.FundBalances {
		funds[fund.FundID] = fund.Value
	}
	return funds
}

type RetailDeposit struct {
	api *API
}

func (retailDeposit RetailDeposit) Deposit(accountID string, currentBalance int, depositAmount int) error {
	newBalance := currentBalance + depositAmount
	return retailDeposit.api.depositFundsToAccount(accountID, newBalance)
}

type EmployerDeposit struct {
	api *API
}

func (employerDeposit EmployerDeposit) Deposit(accountID string, currentBalance int, depositAmount int) error {
	newBalance := currentBalance + depositAmount
	//employer specific logic
	return employerDeposit.api.depositFundsToAccount(accountID, newBalance)
}
