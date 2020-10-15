package controllers

import "github.com/phapli/go-kit/api/middlewares"

func (server *Server) initializeRoutes() {

	// Home Route
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")

	// News routes
	server.Router.HandleFunc("/news", middlewares.SetMiddlewareJSON(server.CreateNews)).Methods("POST")
	server.Router.HandleFunc("/news", middlewares.SetMiddlewareJSON(server.GetAllNews)).Methods("GET")
	server.Router.HandleFunc("/news/{id}", middlewares.SetMiddlewareJSON(server.GetANews)).Methods("GET")
	server.Router.HandleFunc("/news/{id}", middlewares.SetMiddlewareJSON(server.UpdateNews)).Methods("PUT")
	server.Router.HandleFunc("/news/{id}", middlewares.SetMiddlewareJSON(server.DeleteNews)).Methods("DELETE")
}
