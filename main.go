package main

import (
		// "fmt"
    "net/http"
    "encoding/json"
		"google.golang.org/appengine"
)

type Response struct{
	Count int `json:"count"`
}
var count int
func main() {
		count = 0
		http.HandleFunc( "/static/" , func(w http.ResponseWriter, r *http.Request){
			http.ServeFile(w, r , r.URL.Path[1:])
		})
		http.HandleFunc("/app/countUp" , countUpHandler )
		http.HandleFunc("/app/clear" , clearHandler )

    appengine.Main()
}


func countUpHandler( w http.ResponseWriter , r * http.Request ){
	count++ //カウントアップ！
	json.NewEncoder(w).Encode( Response{ Count: count })
}

func clearHandler( w http.ResponseWriter , r * http.Request ){
	count = 0  //リセット
	json.NewEncoder(w).Encode( Response{ Count: count })
}
