package routes

import (
	"github.com/brkss/booksapi/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.EditBook).Methods("PUT")
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
