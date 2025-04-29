package routes

import (
	"github.com/gorilla/mux"
)

// Generate vai retornar um router com as rotas da aplicação
func Generate() *mux.Router {
	return mux.NewRouter()
}
