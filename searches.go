package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SearchController struct {
	db *sql.DB
}

func (s *SearchController) Init(db *sql.DB) {
	s.db = db
}

// -- Handlers --

func (s *SearchController) HandlerSearchFlexible(w http.ResponseWriter, r *http.Request) {

	var search FlexibleSearch

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "failed to read request body with error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &search)
	if err != nil {
		fmt.Fprintf(w, "failed to unmarshal request body with error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	go s.InsertSearchToDB(search)

	//TODO Cached request?

	var flexibleSearchResponse FlexibleSearchResponse

	periods, err := s.FlexibleSearch(search)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if periods == nil {
		fmt.Fprintf(w, "no tests found within that search period - Please refine and try again")
		w.WriteHeader(http.StatusOK)
		return
	}

	flexibleSearchResponse := FlexibleSearchResponse{
		SearchParameters: search,
		HasFailures:      false,
		UptimePercentage: 100,
		Periods:          periods,
	}

	totalTimeSearch := search.TimeEnd - search.TimeStart
	for x in periods (

		)



}

// -- Database Functions --

func (s *SearchController) FlexibleSearch(search FlexibleSearch) (periods []TestPeriods,err error) {
	statement := `
SELECT id, timeended, failure
FROM tests
WHERE environment=$1 AND (timeended BETWEEN $2 AND $3) AND service=$4
`
	rows, err := s.db.Query(statement, search.Environment, search.TimeStart, search.TimeEnd, search.Service)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		




	}
}

func (s *SearchController) InsertSearchToDB(search FlexibleSearch) {
	statement := `
INSERT INTO incomingsearches(timestart, timeend, service, environment) 
VALUES($1,$2,$3,$4)
RETURNING id
`
	var id int64
	_ = s.db.QueryRow(statement, search.TimeStart, search.TimeEnd, search.Service, search.Environment).Scan(&id)
}