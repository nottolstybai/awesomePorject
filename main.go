package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ClientAndMoney struct {
	senderId           int
	senderBalanceId    int
	recipientBalanceId int
	sentMoney          int
}

func SendMoney(w http.ResponseWriter, r *http.Request) {
	var cl ClientAndMoney

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод запрещен!", 405)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&cl)
	fmt.Println(cl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/add_money", SendMoney)

	log.Println("Запуск веб-сервера на http://127.0.0.1:1234")
	err := http.ListenAndServe(":1234", mux)
	log.Fatal(err)
}
