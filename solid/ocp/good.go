package ocp

import (
	"net/http"
)

func GetUserHandlerV2(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, err := extractUserID(req.Form)
	if err != nil {
		resp.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	user := loadUser(userID)
	outputUser(resp, user)
}

func DeleteUserHandlerV2(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, err := extractUserID(req.Form)
	if err != nil {
		resp.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	deleteUser(userID)
}
