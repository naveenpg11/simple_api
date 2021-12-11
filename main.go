package main

import (
    // "encoding/json"
    "fmt"
    "log"
    "net/http"
    // "github.com/gorilla/mux"
)

func ping(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!\n")
    fmt.Fprintf(w, "PONG!!!!!")
    fmt.Println("Endpoint Hit: ping")
}

// func api(w http.ResponseWriter, r *http.Request){

//     vars := mux.Vars(r)
//     id := vars["id"]
//     json.NewEncoder(w).Encode(id)
// }

// func apis(w http.ResponseWriter, r *http.Request){

//     vars := mux.Vars(r)
//     id := vars["id"]
//     json.NewEncoder(w).Encode(id)
// }

func handleRequests() {

    // myRouter := mux.NewRouter().StrictSlash(true)
    http.HandleFunc("/", ping)
    // http.HandleFunc("/api", apis)
    // http.HandleFunc("/api{id}", api)


    log.Fatal(http.ListenAndServe(":8888", nil))
}

func main() {
    fmt.Println("Starting WebServer")
    handleRequests()
}
