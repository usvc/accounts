package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Server class to hold the server
type Server struct {
	options *ServerOptions
	router  *http.ServeMux
}

// ServerOptions to initialize the server with
type ServerOptions struct {
	Port      string
	Interface string
}

var (
	// ServerErrorOK to indicate all is well
	ServerErrorOK = "E_SERVER_OK"
)

func (server *Server) init(opts *ServerOptions) {
	server.options = opts
	server.router = http.NewServeMux()
	// wire up the user layer
	logger.Info("registering user api...")
	userAPI := UserAPI{}
	userAPI.Handle(server.router)
	// wire up the security layer
	logger.Info("registering security api...")
	securityAPI := SecurityAPI{}
	securityAPI.Handle(server.router)
	// wire up the sessions layer
	logger.Info("registering session api...")
	sessionAPI := SessionAPI{}
	sessionAPI.Handle(server.router)
	// let it go wild
	server.handle(server.router)
	server.listen()
}

func (server *Server) listen() {
	server.prelisten()
	bindingInterface := fmt.Sprintf("%v:%v", server.options.Interface, server.options.Port)
	logger.Infof("[server] listening on %v", bindingInterface)
	logger.Error(http.ListenAndServe(bindingInterface, server.router))
}

func (server *Server) prelisten() {
	if server.options == nil {
		panic("[server] init() was not called before attempting to listen()")
	} else if len(server.options.Interface) == 0 {
		panic("[server] interface was not provided")
	} else if len(server.options.Port) == 0 {
		panic("[server] port was not provided")
	} else if value, err := strconv.Atoi(server.options.Port); err != nil {
		panic("[server] port is not a number")
	} else if value < 1000 || value >= 1<<16 {
		panic("[server] port is invalid (1000 <= port <= 65536)")
	}
}

func (server *Server) handle(router *http.ServeMux) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := APIResponse{
			Code:    ServerErrorOK,
			Message: "ok",
			Data:    "ok",
		}
		response.send(w)
	})
}
