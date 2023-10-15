package main

import (
    "Red_Pocket_Game_Server/handlers"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    // Define routing rules
    r.HandleFunc("/login", login.LoginHandler)

    // Start HTTP server
    http.Handle("/", r)  //Bind router r to HTTP root path ("/")
    http.ListenAndServe(":8080", nil)  //Listen on port 8080 and use the default route
}