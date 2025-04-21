package api

import (
	"cushonTechTest/models"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func (api *API) getFunds(writer http.ResponseWriter, request *http.Request) {
	slog.Info("starting getFunds process")

	var funds []models.Fund
	if err := api.DB.Find(&funds).Error; err != nil {
		http.Error(writer, "Error retrieving funds", http.StatusInternalServerError)
		return
	}
	var fundResponse []FundResponse
	for _, fund := range funds {
		fundResponse = append(fundResponse, FundResponse{
			ID:   fund.ID.String(),
			Name: *fund.Name,
		})
	}
	slog.Info(fmt.Sprintf("found %v funds", len(funds)))

	json.NewEncoder(writer).Encode(fundResponse)
}

type FundResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
