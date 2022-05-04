package routes


import(
	"github.com/gorilla/mux"
	"github.com/brkss/booksapi/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods('POST')
	router.HandleFunc("/book/", controllers.GetBook).Method('GET')
	router.HandleFunc("/book/{bookID}", controllers.GetBookById).Method('GET')
	router.HandleFunc("/book/{bookID}", controllers.EditBook).Method('PUT')
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook).Method('DELETE')
}
