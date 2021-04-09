package main

import (
	"awesomeProject/pkg/transport"
	"fmt"
	"net/http"
)

func main() {
	r := transport.Router()
	fmt.Println(http.ListenAndServe(":8080", r))
}
