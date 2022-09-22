// main_test.go
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kurianvarkey/go-todo-api/controllers"
	"github.com/kurianvarkey/go-todo-api/db"
	"github.com/kurianvarkey/go-todo-api/routes"
)

var router *mux.Router
var req *http.Request
var err error
var resp_rec *httptest.ResponseRecorder

func init() {
	log.Println("Kapi Todos test started")
	error_env := godotenv.Load(".env_test")
	if error_env != nil {
		panic("Failed to load the env test file, please check")
	}
}

func setup() {
	log.Println("Kapi Todos running tests")

	router = mux.NewRouter()

	routes.Routes(router)
}

func TruncateTable() {
	app := controllers.NewDb(db.GetConnection())

	app.DB.Exec("TRUNCATE TABLE todos")
}

func TestEndpoints(t *testing.T) {
	setup()

	//Testing get the default endpoint
	/* req, err = http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("Creating 'GET /' request failed!")
	}

	router.ServeHTTP(resp_rec, req)
	if resp_rec.Code != http.StatusOK {
		t.Fatal("Server error: Returned ", resp_rec.Code, " instead of ", http.StatusOK)
	} */

	json_headers := map[string]string{"Content-Type": "application/json"}
	end_points := []struct {
		name              string
		method            string
		uri               string
		payload           []byte
		headers           map[string]string
		status_code       int
		skip_status_check bool
	}{
		{"Get /", "GET", "/", nil, nil, http.StatusOK, true},
		{"Create new todo", "POST", "/todos", []byte(`{"user_id":1,"title":"Test title","description":"Description","task_date":"2022-09-22"}`), json_headers, http.StatusCreated, false},
		{"Get created todo", "GET", "/todos/1", nil, nil, http.StatusOK, false},
		{"Update todo", "PUT", "/todos/1", []byte(`{"user_id":1,"title":"Test title updated","description":"Description updated","task_date":"2022-09-22","is_completed":1,"completed_date":"2022-09-22 15:11:25"}`), json_headers, http.StatusOK, false},
		{"Get created todo", "GET", "/todos/1", nil, nil, http.StatusOK, false},
		{"Get todos", "GET", "/todos", nil, nil, http.StatusOK, false},
		{"Delete todo", "DELETE", "/todos/1", nil, nil, http.StatusOK, false},
	}

	TruncateTable() // clearing all the entries before testing

	for _, tc := range end_points {
		t.Run(tc.name, func(t *testing.T) {
			var payload io.Reader
			if tc.payload != nil {
				payload = bytes.NewBuffer(tc.payload)
			}

			req, err := http.NewRequest(tc.method, tc.uri, payload)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			for key, value := range tc.headers {
				req.Header.Set(key, value)
			}

			//The response recorder used to record HTTP responses
			resp_rec = httptest.NewRecorder()

			router.ServeHTTP(resp_rec, req)
			response := controllers.Response{}
			//json.Unmarshal([]byte(resp_rec.Body.String()), &response)
			json.Unmarshal(resp_rec.Body.Bytes(), &response)

			if resp_rec.Code != tc.status_code {
				t.Fatalf("expected status %v; got %v; code %v", tc.status_code, resp_rec.Result().StatusCode, resp_rec.Code)
			}

			if !tc.skip_status_check && response.Status != 1 {
				t.Fatalf("Eventhough the response was success, there are errors %v", strings.Join(response.Errors, ". "))
			}
		})
	}

}
