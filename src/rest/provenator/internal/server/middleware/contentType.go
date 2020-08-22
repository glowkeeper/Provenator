package middleware

import (
	"net/http"
    "github.com/gorilla/mux"
)

// ContentType - set the default content type
func ContentType() mux.MiddlewareFunc {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Content-Type", "application/json")
            next.ServeHTTP(w, r)
        })
    }
}
