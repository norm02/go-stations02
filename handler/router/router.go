package router

import (
	"database/sql"
	"fmt"
	"net/http"

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

	mux.Handle("/do-panic",middleware.Recover(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		panic("surely panic")
	})))

	mux.Handle("/os",middleware.StoreOS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		os,err := middleware.RegisterOS(r.Context())
		if err!= nil{
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Println(os)
})))

	return mux

}

