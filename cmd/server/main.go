package main

import (
	"fmt"
	"net/http"

	"github.com/pkg/browser"
)

func main() {
	port := ":9090"
	url := "http://localhost" + port
	fmt.Println("Starting Server in", url)
	browser.OpenURL(url)

	if err := http.ListenAndServe(port, http.FileServer(http.Dir("../../assets"))); err != nil {
		fmt.Println("Failed to start server, err: ", err)
	}
}
