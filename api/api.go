package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type API struct {
	DB *gorm.DB
}

func (api *API) RunWebServer() error {
	router := api.setupRouter()
	port := os.Getenv("WEBSERVER_PORT")
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	slog.Info(fmt.Sprintf("Starting server on %s", port))
	return http.ListenAndServe(port, router)
}
func (api *API) setupRouter() *mux.Router {
	router := mux.NewRouter()

	api.registerRoutes(router)

	router.Use(jsonMiddleware)

	router.MethodNotAllowedHandler = methodNotAllowedHandler

	return router
}

func (api *API) registerRoutes(router *mux.Router) {
	router.HandleFunc("/funds", api.getFunds).Methods("GET", "OPTIONS")
	router.HandleFunc("/account", api.getAccount).Methods("GET", "OPTIONS")
	router.HandleFunc("/deposit", api.depositFunds).Methods("POST", "OPTIONS")
}
func enableCors(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		enableCors(writer)
		writer.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
