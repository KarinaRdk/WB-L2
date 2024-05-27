package domain

import (
	"time"
)

type Event struct {
	ID int64
	User_id int64
	Event_name string
	Event_date time.Time
	Start_time time.Time
	End_time time.Time
}