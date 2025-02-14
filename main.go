package main

import (

        "encoding/json"
        "log"
        "net/http"

)

type Message struct {

        Text string `json:"text"`
}


func helloHandler(w http.ResponseWriter, r *http.Request) {


        w.Header().Set("Content-Type", "application/json")

        response := Message{Text:"Hello from Go backend!"}

        json.NewEncoder(w).Encode(response)
}

func main () {

        http.HandleFunc("/api/hello", helloHandler)

        log.Println("Server started on : 8080")

        log.Fatal(http.ListenAndServe(":8080", nil))
}
