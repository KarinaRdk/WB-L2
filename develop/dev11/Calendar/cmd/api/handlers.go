package main 

import (
	"fmt"
	"net/http"
	"time"
	"github.com/KarinaRdk/WB-L2/tree/main/develop/dev11/Calendar/internal/domain"

)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) createEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create an event")
	if r.Method != "POST" {
		//  If you don’t call w.WriteHeader() explicitly, then the first call to
        // w.Write() will automatically send a 200 OK status code to the user
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}
	/*
	    user_id INTEGER NOT NULL,
    event_name TEXT[] NOT NULL,
    event_date TIMESTAMP NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL
	*/
	stmt := `INSERT INTO events (User_id, Event_name, Event_date, Start_time, End_time) VALUES(?, ?, ?, ?,?)`
	event := data.Event{
		ID:          12345,
		User_id:     67890,
		Event_name:  "Conference",
		Event_date:  time.Now(),
		Start_time:  time.Date(2024, 5, 18, 9, 0, 0, 0, time.UTC),
		End_time:    time.Date(2024, 5, 18, 17, 0, 0, 0, time.UTC),
	}
	err := app.writeJSON(w, http.StatusOK, event, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		}

	}
func (app *application) getEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get an event")
}