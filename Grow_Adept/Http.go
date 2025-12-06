package main

import (
	"encoding/json"
	"net/http"
)

type Json struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/hello", helloHandle)
	http.ListenAndServe(":8080", nil)
}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	mensagem := Json{
		Message: "OIIII Gatinho, agora em json"}
	w.Header().Set("Contensdsdt-Type", "applicdsation/json")

	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(mensagem)
}
