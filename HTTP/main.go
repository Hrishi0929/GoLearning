package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error Occured:", err)
		os.Exit(1)
	}
	fmt.Println(response)
}
