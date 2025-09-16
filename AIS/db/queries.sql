-- name: CreatePositionReportTableIfNotExist :exec
CREATE TABLE IF NOT EXISTS position_reports (
                                  mmsi BIGINT PRIMARY KEY,
                                    ship_name VARCHAR(255),
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

-- name: EmptyDBTables :exec
TRUNCATE TABLE position_reports
RESTART IDENTITY CASCADE;

-- name: UpsertPositionEntry :exec
INSERT INTO position_reports (
    mmsi,
    ship_name,
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
           $1,  -- mmsi
           $2,  -- ship_name
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
           $15, --valid
            $16 --timeUtc
       )
ON CONFLICT (mmsi) DO UPDATE
    SET
        ship_name = EXCLUDED.ship_name,
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

-- name: GetPositionData :many
SELECT *
FROM position_reports
WHERE latitude  BETWEEN $1 AND $2 --minLat --maxLat
AND longitude BETWEEN $3 AND $4; --minLon --maxLon

