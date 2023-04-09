package ocp

import (
	"net/http"
	"strconv"
)

func GetUserHandlerV1(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	userID,
		err := strconv.ParseInt(req.Form.Get("UserID"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	user := loadUser(userID)
	outputUser(resp, user)
}

func DeleteUserHandlerV1(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	userID,
		err := strconv.ParseInt(req.Form.Get("UserID"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	deleteUser(userID)
}
