package router 

import(
	"go-postgres-yt/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.ROuter{
	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}", midddleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.GETAllStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newstock", midddleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}",midddleware.UpdateStock).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/deletestock/{id}",middleware.Deletestock).Methods("DELETE","OPTIONS")
}