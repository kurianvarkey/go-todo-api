// Controller of the the application. This serve as the base controller and we will inject this in the routes
package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/kurianvarkey/go-todo-api/db"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

/*
// some useful time constants
const (
    Layout      = "01/02 03:04:05PM '06 -0700"
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700"
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"

    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
)
*/

const (
	date_layout      = "2006-01-02"
	date_time_layout = "2006-01-02 15:04:05"
)

// declaring App struct to hold the db connection
type App struct {
	DB *gorm.DB
}

type Response struct {
	Status   int8        `json:"status"`
	Errors   []string    `json:"errors"`
	Messages []string    `json:"messages"`
	Meta     db.Meta     `json:"meta,omitempty"`
	Results  interface{} `json:"results"`
}

// NewDb: function to create new database connection
func NewDb(db *gorm.DB) App {
	return App{db}
}

// GetInputs: this is a standard function read the form raw request or form-url-encoded
//
// Output: as url.Values
func (a App) GetInputs(r *http.Request) (url.Values, error) {
	input := url.Values{}
	content_type := r.Header.Get("Content-type")

	//we will try to set all the query string from url
	for key, val := range r.URL.Query() {
		input.Set(key, val[0])
	}

	if content_type == "application/x-www-form-urlencoded" {
		err := r.ParseForm()
		if err != nil {
			fmt.Println("ParForm error", err)
			return nil, err
		}

		input = r.Form // r.Form is already url.Values
	} else if content_type == "text/plain" || content_type == "application/json" {
		defer r.Body.Close()
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}

		// from the above code we will get the []byte array
		raw_input := make(map[string]interface{})
		err = json.Unmarshal(raw, &raw_input)
		if err != nil {
			return nil, err
		}

		// loop through the json to create url.Values
		for key, val := range raw_input {
			// val is an interface type, so use fmt.Sprint(val) to convert to string
			input.Set(key, fmt.Sprint(val)) // use Set to add key and values.
		}
	} else {
		//return nil, errors.New("Content type not supported: " + content_type)
	}

	return input, nil
}

// SendOutput: created a standard response function so that I can add more functionalities later
// If success it will send the status as 1 and data will have the result
// If error, the status will be 0 and errors array will have the error message
func (a App) SendOutput(w http.ResponseWriter, status int, result any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	response := Response{} // creating a response struct varaible
	if status == http.StatusBadRequest || status == http.StatusNotFound || status == http.StatusInternalServerError {
		response.Status = 0
		response.Results = nil
		response.Errors = []string{cast.ToString(result)}
	} else {
		// if pagination struct then handle the meta and results
		if reflect.TypeOf(result).String() == "db.Pagination" {
			/*
				// This way also you can get the values using reflect way, but it is not a good practice, just keeping for reference
				response.Meta.Limit = cast.ToInt(reflect.ValueOf(result).FieldByName("Limit").Interface())
				response.Meta.Page = cast.ToInt(reflect.ValueOf(result).FieldByName("Page").Interface())
				response.Meta.TotalRows = cast.ToInt64(reflect.ValueOf(result).FieldByName("TotalRows").Interface())
				response.Meta.TotalPages = cast.ToInt(reflect.ValueOf(result).FieldByName("TotalPages").Interface())
				response.Results = reflect.ValueOf(result).FieldByName("Results").Interface()
			*/

			// we are marshalling the struct and merging the response struct
			ja, _ := json.Marshal(result)

			// the below line will preserve the key order in the result struct. comment below both lines to see the difference
			pre_compute_records := json.RawMessage(`{"precomputed": true}`)
			response = Response{Results: &pre_compute_records} // we are doing to make effect of the above line
			response.Status = 1

			json.Unmarshal(ja, &response)
		} else {
			response.Status = 1
			response.Results = result
		}

		response.Messages = []string{"The call is success"}
	}

	json.NewEncoder(w).Encode(response)
}

// Paginate function with pagination meta
func (a App) Paginate(r *http.Request, db *gorm.DB, source interface{}, pagination *db.Pagination) func(db *gorm.DB) *gorm.DB {
	var total_rows int64
	db.Model(source).Count(&total_rows)

	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page == 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(q.Get("limit"))

	pagination.SetLimit(limit)
	pagination.SetPage(page)
	pagination.SetTotalRows(total_rows)
	pagination.SetTotalPages()

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
