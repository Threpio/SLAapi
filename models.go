package main

//Test
//Tests that are placed within the database
type Test struct {
	TimeStarted int64
	TimeEnded   int64
	TestBody    string
	Environment string
	Failure     bool
}

func (t Test) Validate() (success bool) {

	if (t.TimeStarted > 160000) && (t.TimeEnded > 160000) {
		if t.TestBody != "" && t.Environment != ""{
			return true

		}
	}

	return false
}

//FlexibleSearch
//An incoming search Request
type FlexibleSearch struct {
	TimeStart   int64
	TimeEnd     int64
	Service     string
	Environment string
}

//FlexibleSearchResponse
//The api response for a FlexibleSearch api request
type FlexibleSearchResponse struct {
	SearchParameters FlexibleSearch
	HasFailures 	 bool
	UptimePercentage float32
	Periods          []TestPeriods
}

//TestPeriods
//Denotes a period of time of ongoing
type TestPeriods struct {
	TimeStart int64
	TimeEnd   int64
	Failures  bool
}
