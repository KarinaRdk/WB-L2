ALTER TABLE events ADD CONSTRAINT event_date_check CHECK (event_date >= CURRENT_DATE);
ALTER TABLE events ADD CONSTRAINT event_time_check CHECK (start_time <= end_time);