
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    event_name TEXT[] NOT NULL,
    event_date TIMESTAMP NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL
);
