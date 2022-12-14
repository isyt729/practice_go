package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
)

func main(){
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "<html><body>Hello</body></html>")
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request){
		val := 	r.FormValue("say")
		fmt.Fprint(w, "<html><body></body></html>",
					html.EscapeString(val))
	})

	err := http.ListenAndServe("localhost:8080",nil)

	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
