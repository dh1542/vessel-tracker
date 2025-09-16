-- name: CreateShipTableIfNotExist :exec
CREATE TABLE IF NOT EXISTS ships  (
                       mmsi BIGINT PRIMARY KEY,
                       ship_name VARCHAR(255)
);

-- name: CreatePositionReportTableIfNotExist :exec
CREATE TABLE IF NOT EXISTS position_reports (
                                  mmsi BIGINT PRIMARY KEY REFERENCES ships(mmsi) ON DELETE CASCADE,
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


-- name: CreateShip :exec
INSERT INTO ships (mmsi, ship_name)
VALUES ($1, $2)
ON CONFLICT (mmsi) DO NOTHING;

-- name: GetShip :one
SELECT mmsi, ship_name
FROM ships
WHERE MMSI = $1;

-- name: EmptyDBTables :exec
TRUNCATE TABLE position_reports
RESTART IDENTITY CASCADE;

-- name: UpsertPositionEntry :exec
INSERT INTO position_reports (
    mmsi,
    latitude,
    longitude,
    cog,
    sog,
    true_heading,
    navigational_status,
    position_accuracy,
    communication_state,
    rate_of_turn,
    special_manoeuvre_indicator,
    repeat_indicator,
    message_id,
    valid,
    time_utc
)
VALUES (
           $1,  -- id
           $2,  -- mmsi
           $3,  -- latitude
           $4,  -- longitude
           $5,  -- cog
           $6,  -- sog
           $7,  -- true_heading
           $8,  -- navigational_status
           $9,  -- position_accuracy
           $10, -- communication_state
           $11, -- rate_of_turn
           $12, -- special_manoeuvre_indicator
           $13, -- repeat_indicator
           $14, -- message_id
           $15 -- valid
       )
ON CONFLICT (mmsi) DO UPDATE
    SET
        latitude = EXCLUDED.latitude,
        longitude = EXCLUDED.longitude,
        cog = EXCLUDED.cog,
        sog = EXCLUDED.sog,
        true_heading = EXCLUDED.true_heading,
        navigational_status = EXCLUDED.navigational_status,
        position_accuracy = EXCLUDED.position_accuracy,
        communication_state = EXCLUDED.communication_state,
        rate_of_turn = EXCLUDED.rate_of_turn,
        special_manoeuvre_indicator = EXCLUDED.special_manoeuvre_indicator,
        repeat_indicator = EXCLUDED.repeat_indicator,
        message_id = EXCLUDED.message_id,
        valid = EXCLUDED.valid,
        time_utc = EXCLUDED.time_utc;