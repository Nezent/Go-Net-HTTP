package routes

import (
	"database/sql"
	"net/http"

	"github.com/Nezent/Go-Net/controller"
)


func PageRoute(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("content-type","application/json")
		switch r.Method {
		case "GET":
			controller.GetUser(db,w)
		case "POST":
			controller.CreateUser(db,w,r)
		case "PATCH":
			controller.GetUser(db,w)
		case "DELETE":
			controller.GetUser(db,w)
		}
	}
}

