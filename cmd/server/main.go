package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":9090"
	fmt.Println("Starting Server in http://localhost" + port)
	if err := http.ListenAndServe(port, http.FileServer(http.Dir("../../assets"))); err != nil {
		fmt.Println("Failed to start server, err: ", err)
	}
}
