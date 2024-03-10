package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// 명언 구조체 정의
type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func randomQuote() Quote {
	// 명언을 가져올 외부 API URL
	url := "https://type.fit/api/quotes"

	// 외부 API에 GET 요청 보내기
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching quotes:", err)
		return Quote{}
	}
	defer response.Body.Close()

	// 응답을 JSON으로 파싱
	var quotes []Quote
	err = json.NewDecoder(response.Body).Decode(&quotes)
	if err != nil {
		fmt.Println("Error parsing quotes:", err)
		return Quote{}
	}

	// 명언 중에서 랜덤으로 선택
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	return quotes[randomIndex]
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	quote := randomQuote()

	// JSON 형식으로 응답 반환
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

func main() {
	http.HandleFunc("/quote", quoteHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
