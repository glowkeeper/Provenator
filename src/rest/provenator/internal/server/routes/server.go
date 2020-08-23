package routes

import (
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(`{"text": "method disallowed"}`))
}

func home(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"text": "Provenator RESTful Server"}`))
}
