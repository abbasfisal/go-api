package main

import (
	"fmt"
	"github.com/abbasfisal/go-api/helper"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	fmt.Println("start server")

	router := httprouter.New()
	router.GET("/ping", pingHandler)
	router.GET("/", homeHandler)

	helper.PanicIfError(http.ListenAndServe("localhost:8080", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "welcome to Home Page")
}

func pingHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "ping check")
}
