package acme_good

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func shortMethods(resp http.ResponseWriter, req *http.Request) {
	userId, err := extractUserId(req)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	person, err := loadPerson(userId)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	outputPerson(resp, person)
}

func extractUserId(req *http.Request) (int64, error) {
	err := req.ParseForm()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(req.Form.Get("userId"), 10, 64)
}

func loadPerson(userId int64) (*Person, error) {
	row := DB.QeuryRow("SELECT name FROM users WHERE id = ?", userId)

	person := &Person{}
	err := row.Scan(person.Id, person.Name, person.Phone)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func outputPerson(resp http.ResponseWriter, person *Person) {
	encoder := json.NewEncoder(resp)
	err = encoder.Encode(person)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
}
