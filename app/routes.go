package app

import "github.com/muhdanfyan/onlineshop_golang/app/controllers"

func (server *Server) InitializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
