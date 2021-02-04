package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

func helloWorldHandler(res http.ResponseWriter, req *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	callback := req.URL.Query().Get("callback")
	if callback != "" {
		req.Header.Set("Content-Type", "application/javascript")
		fmt.Fprintf(res, "%s(%s)", callback, string(data))
	} else {
		fmt.Fprintf(res, string(data))
	}
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
