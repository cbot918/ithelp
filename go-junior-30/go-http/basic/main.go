package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port = ":8887"

func main() {

	// 這邊我們改模擬個 login api
	http.HandleFunc("/login", Login)

	// 示範怎麼快速把靜態檔案提供出去
	http.Handle("/", http.FileServer(http.Dir("public")))

	fmt.Println("listening: " + port)
	// 這邊做個錯誤捕捉, 有時候自己port被佔到不自知, 有錯誤訊息就比較省時間
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// 習慣上會先用個 struct 起來接資料
type userParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	// 由於內建沒有捕捉 method, 我們自己寫
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "method not allowed")
		return
	}

	// 這邊開始解析 body資料, 個人是滿喜歡 json decode/encode 這種方式
	user := userParam{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "data error, please try again")
		return
	}
	// 這邊 log 出來看, 資料實際有接收到
	fmt.Printf("%+v", user)

	// 這邊設定個 headers
	w.Header().Set("Content-Type", "application/json")
	// 這邊把 w 拿去做 encoder, 資料就會直接出去了
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "json encode failed")
		return
	}

}
