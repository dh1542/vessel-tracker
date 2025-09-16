
CREATE TABLE position_reports (
    mmsi BIGINT PRIMARY KEY,
    ship_name VARCHAR,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    cog INTEGER,
    sog INTEGER,
    true_heading INTEGER,
    navigational_status INTEGER,
    position_accuracy BOOLEAN,
    communication_state BIGINT,
    rate_of_turn INTEGER,
    special_manoeuvre_indicator INTEGER,
    repeat_indicator INTEGER,
    message_id INTEGER,
    valid BOOLEAN,
    time_utc TIMESTAMP
);
