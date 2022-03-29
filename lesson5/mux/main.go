package main

import (
	"PicusBootcamp/lesson5/mux/httpErrors"
	"context"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func main() {
	r := mux.NewRouter()
	handlers.AllowedOrigins([]string{"https://www.example.com"})
	handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH"})

	r.Use(loggingMiddleware)
	//r.Use(authenticationMiddleware)

	s := r.PathPrefix("/products").Subrouter()
	s.HandleFunc("/{name}", ProductNameHandler)

	p := r.PathPrefix("/user").Subrouter()
	p.HandleFunc("/", userCreate).Methods(http.MethodPost)
	srv := &http.Server{
		Addr:         "localhost:8090",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	ShutdownServer(srv, time.Second*10)
}

type ApiResponse struct {
	Data interface{} `json:"data"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ProductNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//r.URL.Query().Get("param")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	d := ApiResponse{Data: vars["name"]}

	resp, _ := json.Marshal(d)
	w.Write(resp)
}

func userCreate(w http.ResponseWriter, r *http.Request) {
	var u User
	if r.Header.Get("Content-Type") != "application/json" {
		data, _ := json.Marshal(httpErrors.NewRestError(http.StatusBadRequest, httpErrors.ContentTypeError.Error()))
		w.Write(data)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		data, _ := json.Marshal(httpErrors.NewRestError(http.StatusBadRequest, httpErrors.CannotDecodeError.Error()))
		w.Write(data)
		return
	}

	personData, err := json.Marshal(u)
	if err != nil {
		data, _ := json.Marshal(httpErrors.NewRestError(http.StatusBadRequest, httpErrors.CannotDecodeError.Error()))
		w.Write(data)
		return
	}
	w.Write(personData)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Query())

		next.ServeHTTP(w, r)
	})
}

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if strings.HasPrefix(r.URL.Path, "/products") {
			if token != "" {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Token not found", http.StatusUnauthorized)
			}
		} else {
			next.ServeHTTP(w, r)
		}

		next.ServeHTTP(w, r)
	})
}

// ShutdownServer -> Graceful Shutdown
func ShutdownServer(srv *http.Server, timeout time.Duration) {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	log.Println("shutting down...")

	srv.Shutdown(ctx)
	os.Exit(0)
}
