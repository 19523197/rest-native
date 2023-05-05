package main

import (
	"log"
	"net/http"
	"rest-native/database"
	"rest-native/handler"
	"rest-native/middleware"
	"rest-native/repository"
	"rest-native/utils"
	"time"
)

func main() {
	db, err := database.SetupSql()
	utils.CheckError(err)

	repo := repository.BaseRepository{db}

	userRepo := repository.UserRepository{repo}
	userHandler := handler.UserHandler{userRepo}
	productRepo := repository.ProductRepository{repo}
	productHandler := handler.ProductHandler{productRepo}

	mux := new(middleware.CustomMux)
	mux.RegisterMiddleware(middleware.MiddlewareJWTAuthorization)

	mux.HandleFunc("/login", userHandler.Login)
	mux.HandleFunc("/", handler.Ping)
	mux.HandleFunc("/product", productHandler.Handle)
	mux.HandleFunc("/product/{id}", productHandler.HandleOne)

	s := &http.Server{
		Addr:         ":81",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
