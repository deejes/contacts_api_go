package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    //"encoding/json"
)

func main(){
    router := mux.NewRouter()
    log.Fatal(http.ListenAndServe(":8000",router))
}
