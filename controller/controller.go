package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Nezent/Go-Net/models"
	_ "github.com/lib/pq"
)

func GetUser(db *sql.DB, w http.ResponseWriter) {

	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()


	users := []models.User{}
	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.NAME, &user.EMAIL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func CreateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.NAME, user.EMAIL,user.PASSWORD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

