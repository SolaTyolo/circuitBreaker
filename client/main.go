package main

import (
	"fmt"
	"net/http"

	"github.com/SolaTyolo/circuitBreaker/client/config"
	"github.com/SolaTyolo/circuitBreaker/client/handler"
	"github.com/SolaTyolo/circuitBreaker/client/request"
)

func main() {

	circuitBreaker := config.CircuitBreakerConfig()
	request := request.NewHttpRequest(circuitBreaker, http.Client{}, "http://localhost:8082/api/v1/ping")
	handler := handler.NewHandler(*request)

	http.HandleFunc("/api/v1/ping", handler.Ping)
	fmt.Println("Client Application running .")
	http.ListenAndServe(":8080", nil)
}
