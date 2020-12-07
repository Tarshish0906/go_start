package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//ハンドラの登録
	http.HandleFunc("/", helloHandler)

	fmt.Println("hrrp://lpcalhost:8080 で起動中...")
	//HTTPサーバーを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World!!!"
	fmt.Fprint(w, msg)
}
