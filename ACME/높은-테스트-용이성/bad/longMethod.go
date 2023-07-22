package acme_bad

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Person struct {
	Id    int64
	Name  string
	Phone string
}

func longMethod(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		resp.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	userId, err := strconv.ParseInt(req.Form.Get("userId"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	row := DB.QeuryRow("SELECT name FROM users WHERE id = ?", userId)

	person := &Person{}
	err = row.Scan(person.Id, person.Name, person.Phone)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(resp)
	err = encoder.Encode(person)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
}
