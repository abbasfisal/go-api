package main

import (
	"fmt"
	"github.com/abbasfisal/go-api/helper"
	"net/http"
)

func main() {
	fmt.Println("start server")

	helper.PanicIfError(http.ListenAndServe("localhost:8080", nil))
}
