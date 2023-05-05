package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"rest-native/repository"
	"rest-native/utils"
	"strconv"
)

type ProductHandler struct {
	Repo repository.ProductRepository
}

func (h *ProductHandler) Handle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		product, err := h.Repo.GetAllProduct()
		if err != nil {
			response, err := json.Marshal(err)
			log.Println(err)
			w.WriteHeader(500)
			w.Write(response)
		}
		response, err := json.Marshal(product)
		w.WriteHeader(200)
		w.Write(response)
	case "POST":

	default:
		data := "Soon to be implemented"
		response, _ := json.Marshal(data)
		w.Write(response)
	}
}

func (h *ProductHandler) HandleOne(w http.ResponseWriter, r *http.Request) {
	url, err := url.Parse(r.URL.String())
	utils.CheckError(err)
	id, err := strconv.ParseInt(url.Path, 0, 0)
	log.Println(id)
	switch r.Method {
	case "GET":
		product, err := h.Repo.GetOneProduct(id)
		if err != nil {
			response, err := json.Marshal(err)
			log.Println(err)
			w.WriteHeader(500)
			w.Write(response)
		}
		response, err := json.Marshal(product)
		w.WriteHeader(200)
		w.Write(response)
	case "PUT":
	case "DELETE":

	default:
		data := "Soon to be implemented"
		response, _ := json.Marshal(data)
		w.Write(response)
	}
}
