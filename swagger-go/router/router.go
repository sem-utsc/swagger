package router

import (
	"net/http"
	"swagger-go/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/characters", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetCharactersHandler(w, r)
		case http.MethodPost:
			handlers.CreateCharacterHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/characters/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			handlers.DeleteCharacterHandler(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
