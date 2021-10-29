package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TestController struct {
	db *sql.DB
}

func (t *TestController) Init(db *sql.DB) {
	t.db = db
}

// -- Handlers --

func (t *TestController) HandlerInsertTest(w http.ResponseWriter, r *http.Request) {

	var te Test

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "failed to read request body with error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &te)
	if err != nil {
		fmt.Fprintf(w, "failed to unmarshal request body with error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !te.Validate() {
		fmt.Fprintf(w, "validation of test body failed")
	}

	err = t.InsertTestToDB(te)
	if err != nil {
		fmt.Fprintf(w, "failed to place test into database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "placed into database successfully")
	w.WriteHeader(http.StatusCreated)
	return
}

// -- Database Functions --

//InsertTestToDB places a test into the db - Validate test first!
func (t *TestController) InsertTestToDB(test Test) (err error) {
	statement := `
INSERT INTO tests(timestarted, timeended, testbody, failure, environment) 
VALUES($1, $2, $3, $4, $5)
RETURNING id
`
	var id int64
	err = t.db.QueryRow(statement, test.TimeStarted, test.TimeEnded, test.TestBody, test.Environment).Scan(&id)
	return err
}
