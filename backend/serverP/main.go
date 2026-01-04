package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	http.HandleFunc("/api/greet", greetHandler)
	fmt.Println("run server")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("error server:", err)
	}
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if http.MethodOptions == r.Method {
		w.WriteHeader(http.StatusOK)
		return
	}
	if http.MethodPost != r.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	req, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error1:", err)
		return
	}
	var reqData Request
	if err := json.Unmarshal(req, &reqData); err != nil {
		fmt.Println("erorr2", err)
		return
	}
	res := Response{
		Message: "test" + reqData.Name,
	}
	if err = json.NewEncoder(w).Encode(&res); err != nil {
		fmt.Println("error3", err)
		return
	}
}
