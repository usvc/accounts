package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	options *serverOptions
	handler *http.Handler
}

type serverOptions struct {
	Port      string
	Interface string
}

var server = Server{}

func (server *Server) init(opts *serverOptions) {
	server.options = opts
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	server.listen()
}

func (server *Server) listen() {
	server.prelisten()
	bindingInterface := fmt.Sprintf("%v:%v", server.options.Interface, server.options.Port)
	logger.infof("listening on %v", bindingInterface)
	logger.error(http.ListenAndServe(bindingInterface, nil))
}

func (server *Server) prelisten() {
	if server.options == nil {
		panic("init() was not called before attempting to listen()")
	} else if len(server.options.Interface) == 0 {
		panic("interface was not provided")
	} else if len(server.options.Port) == 0 {
		panic("port was not provided")
	} else if value, err := strconv.Atoi(server.options.Port); err != nil {
		panic("port is not a number")
	} else if value < 1000 || value >= 1<<16 {
		panic("port is invalid (1000 <= port <= 65536)")
	}
}
