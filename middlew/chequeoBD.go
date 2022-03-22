package middlew

import (
	"net/http"

	"github.com/victor-perez-palma/twittor/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
		}
		next.ServeHTTP(w, r)
	}
}
