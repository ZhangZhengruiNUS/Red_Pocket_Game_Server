package login

import (
    "net/http"
	"encoding/json"
	"fmt"
	"context"
	"Red_Pocket_Game_Server/db/sqlc"
	"database/sql"
	_"github.com/lib/pq"
	"io/ioutil"
)

// temporary
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:admin123@localhost:5432/red_pocket_game?sslmode=disable"
)

type LoginData struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in the LoginHandler")

    // Parsing incoming data
	var loginData LoginData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginData); err != nil {
		fmt.Println("decode error")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("Request Body:", string(body))
    username := loginData.Username
    password := loginData.Password
	fmt.Println("password=" + password)
	fmt.Println("username=" + username)
	// Establishing a database connection
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		fmt.Println("database error")
		return
	}

	queries := db.New(conn);  //Create an Queries structure instance

    // Execute validation logic
	account, err := queries.GetAccountByName(context.Background(), username)  //Calling the GetAccountByName method

	if err == nil {
		// Handle success
		fmt.Println("query success")
		
			if account.Password == password {
				fmt.Println("password right")
				successResponse := map[string]interface{}{
					"message": "Login success",
					}
				jsonResponse, err := json.Marshal(successResponse)  //Convert map to JSON format using encoding/JSON
				if err != nil {
					http.Error(w, "JSON encoding error", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(jsonResponse)
			} else {
				fmt.Println("password wrong")
				successResponse := map[string]interface{}{
					"message": "Login failed",
					}
				jsonResponse, err := json.Marshal(successResponse)
				if err != nil {
					http.Error(w, "JSON encoding error", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(jsonResponse)
			}
	} else {
		// Handle error
		fmt.Println("query error")
		errorResponse := map[string]interface{}{
			"error": "query error",
		}
		jsonResponse, err := json.Marshal(errorResponse)
		if err != nil {
			http.Error(w, "JSON encoding error", http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonResponse)
	}
	
}