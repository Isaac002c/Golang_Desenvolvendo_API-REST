package main

import (

"fmt"
"net/http"
"log"
)

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Home Page")
}

func HandleRequest(){
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}


func main(){
	fmt.Println("Iniciando o Servidor Rest GO")
	HandleRequest()
}