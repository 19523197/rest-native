package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type HandlerCont interface {
	Handle(w http.ResponseWriter, r *http.Request)
	HandleOne(w http.ResponseWriter, r *http.Request)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	var data string = "ping"
	log.Println(r.URL)
	res, _ := json.Marshal(data)
	w.Write(res)
}
