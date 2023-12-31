package router

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	healthHandler := handler.NewHealthzHandler()
	mux.HandleFunc("/healthz", healthHandler.ServeHTTP)

	todoService := service.NewTODOService(todoDB)
	todoHandler := handler.NewTODOHandler(todoService)
	mux.HandleFunc("/todos", todoHandler.ServeHTTP)

	/*
		//"/do-panic"にアクセスすると、panicするmutex
		mux.HandleFunc("/do-panic",http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			panic("surely panic")
		}))
	*/

	//"/do-panic"にアクセスしても、recoverしてpanicしないmutex
	mux.Handle("/do-panic", middleware.Recover(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("recover panic")
	})))

	mux.Handle("/os", middleware.StoreOS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		os, err := middleware.CtxOS(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Println(os)
	})))

	mux.Handle("/accesslog", middleware.StoreOS(middleware.AccessLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("accesslog is written")
	}))))

	mux.Handle("/userauth", middleware.Basicauth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte("basicauth is permissioned"))
	}), "ユーザー名とパスワードを入力してください。"))

	mux.Handle("/gracefulshutdown", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 5)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte("graceful shutdown"))
	}))

	return mux

}
