package acme_bad

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestLongMethod(t *testing.T) {
	// 요청에 대한 빌드를 수행
	request := &http.Request{}
	request.PostForm = url.Values{}
	request.PostForm.Add("userId", "1")

	// 데이터베이스에 대한 모의를 구현
	var mockDB sqlmock.Sqlmock
	var err error

	DB, mockDB, err = sqlmock.New()
	require.NoError(t, err)
	mockDB.ExpectQuery("SELECT name FROM users WHERE id = ").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone"}).AddRow(1, "John", "555-1234"))

	// 등답에 대한 빌드를 수행
	response := httptest.NewRecorder()

	// 함수를 호출
	logMethod(response, request)

	// 등답에 대한 유효성 검사 수행
	require.Equal(t, http.StatusOK, response.Code)

	// JSON에 대한 유효성 검사 수행
	responseBytes, err := ioutil.ReadAll(response.Body)
	require.NoError(t, err)

	expectedJSON := `{"id":1,"name":"John","phone":"555-1234"}`
	assert.Equal(t, expectedJSON, string(responseBytes))
}
