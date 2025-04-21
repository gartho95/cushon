package api

import (
	"fmt"
	"log/slog"
	"net/http"
)

var methodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	slog.Error(fmt.Sprintf("Method Not Allowed: %s %s", r.Method, r.URL.Path))
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
})
