package main

import (
	"api/database"
	"api/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {

	db := database.Connect()

	router := mux.NewRouter()
	routes.RegisterRoutes(router, db)
	// defer db.Close()

	fmt.Println("Server has started ✅")
	http.ListenAndServe(":8000", jsonContentTypeMiddleware(router))

}
