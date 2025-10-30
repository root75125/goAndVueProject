package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// リクエスト用の構造体
type Request struct {
	Name string `json:"name"`
}

// レスポンス用の構造体
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

	// ハンドラーの登録
	http.HandleFunc("/api/greet", greetHandler)

	log.Println("Server starting on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("server stop")
	}
}

// ハンドラー関数を定義
func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reach backend")
	// CORSの設定
	// すべてのオリジン（ドメイン）からのアクセスを許可
	// 本番環境では特定のオリジンのみを許可することを推奨
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 許可するHTTPメソッドを指定
	// POST: データの送信
	// OPTIONS: プリフライトリクエスト（ブラウザが実際のリクエスト前に送信する確認リクエスト）
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	// 許可するHTTPヘッダーを指定
	// Content-Type: リクエストのデータ形式を指定するヘッダー
	// 例：application/json などの指定を許可
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// リクエストボディの解析
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// レスポンスの作成
	response := Response{
		Message: "こんにちは、" + req.Name + "さん！",
	}

	// JSONレスポンスの返却
	json.NewEncoder(w).Encode(response)
}

// curlコマンド例:curl  http://localhost:8080/api/greet -X POST -d '{"name":"ran"}'
