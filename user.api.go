package main

import (
	"net/http"
	"strings"
)

type UserAPI struct{}

var UserErrorOk = "E_USER_OK"
var UserErrorInvalidParameters = "E_INVALID_PARAMS"

var userApi = UserAPI{}

func (userApi *UserAPI) handle(router *http.ServeMux) {
	router.Handle("/user/", &UserAPI{})
}

func (userApi *UserAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(500)
			response := Response{
				ErrorCode: "E_GENERIC",
				Message:   "an unexpected error has occurred",
				Data:      r,
			}
			response.send(w)
		}
	}()
	switch r.Method {
	case http.MethodGet:
		userApi.get(w, r)
	}
}

func (*UserAPI) get(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(strings.TrimRight(r.URL.Path, "/"), "/")
	uuid := url[len(url)-1]
	logger.infof("retrieving user with uuid '%s'", uuid)
	if data, err := user.GetByUuid(uuid); err != nil {
		panic(err)
	} else {
		response := Response{
			ErrorCode: UserErrorOk,
			Message:   "ok",
			Data:      data,
		}
		response.send(w)
	}
}
