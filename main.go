package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Quote struct {
    Body   string `json:"body"`
    Author string `json:"author"`
}

func goHandler(w http.ResponseWriter, r *http.Request) {
    // 명언 데이터 생성
    quote := Quote{
        Body:   "명언 내용",
        Author: "명언 작가",
    }

    // JSON 형식으로 응답 반환
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(quote)
}

func main() {
    http.HandleFunc("/go-endpoint", goHandler)
    http.ListenAndServe(":8080", nil)
}
