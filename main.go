package main

import (
	"fmt"
	"log"
	"net/http"
)

// cookieの設定+標準出力へプリント
func setCookies(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "hoge",
		Value: "bar",
	}
	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "Hello Cookie!")

	cookies := r.Cookies()

	fmt.Println("start")
	fmt.Println(cookies)
	fmt.Println("end")
}

// メイン
func main() {
	http.HandleFunc("/", setCookies)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
