package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// 启用静态服务
	h := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", h))
	http.HandleFunc("/", showGuess)
	err := http.ListenAndServe(":12345", nil)
	checkErr(err)
}

func showGuess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("guess.html")
		checkErr(err)
		err = t.Execute(w, nil)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("%#v \n", err)
	}
}
