package routes


import(
	"github.com/gorilla/mux"
	"github.com/iamajraj/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func (router *mux.Router)  {
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}