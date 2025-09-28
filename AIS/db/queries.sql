-- name: CreatePositionReportTableIfNotExist :exec
CREATE TABLE IF NOT EXISTS position_reports
(
    mmsi                        BIGINT PRIMARY KEY NOT NULL,
    ship_name                   VARCHAR            NOT NULL,
    latitude                    DOUBLE PRECISION   NOT NULL,
    longitude                   DOUBLE PRECISION   NOT NULL,
    cog                         INTEGER            NOT NULL,
    sog                         INTEGER            NOT NULL,
    true_heading                INTEGER            NOT NULL,
    navigational_status         INTEGER            NOT NULL,
    position_accuracy           BOOLEAN            NOT NULL,
    communication_state         BIGINT             NOT NULL,
    rate_of_turn                INTEGER            NOT NULL,
    special_manoeuvre_indicator INTEGER            NOT NULL,
    repeat_indicator            INTEGER            NOT NULL,
    message_id                  INTEGER            NOT NULL,
    valid                       BOOLEAN            NOT NULL,
    time_utc                    TIMESTAMP          NOT NULL,
    destination                 VARCHAR
);

-- name: CreateImagesTableIfNotExist :exec
CREATE TABLE IF NOT EXISTS images
(
    ship_name VARCHAR NOT NULL,
    image_url VARCHAR NOT NULL
);

-- name: EmptyDBTables :exec
TRUNCATE TABLE position_reports
    RESTART IDENTITY CASCADE;
TRUNCATE TABLE images
    RESTART IDENTITY CASCADE;

-- name: UpsertPositionEntry :exec
INSERT INTO position_reports (mmsi,
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
                              time_utc)
VALUES ($1, -- mmsi
        $2, -- ship_name
        $3, -- latitude
        $4, -- longitude
        $5, -- cog
        $6, -- sog
        $7, -- true_heading
        $8, -- navigational_status
        $9, -- position_accuracy
        $10, -- communication_state
        $11, -- rate_of_turn
        $12, -- special_manoeuvre_indicator
        $13, -- repeat_indicator
        $14, -- message_id
        $15, --valid
        $16)
ON CONFLICT (mmsi) DO UPDATE
    SET ship_name                   = EXCLUDED.ship_name,
        latitude                    = EXCLUDED.latitude,
        longitude                   = EXCLUDED.longitude,
        cog                         = EXCLUDED.cog,
        sog                         = EXCLUDED.sog,
        true_heading                = EXCLUDED.true_heading,
        navigational_status         = EXCLUDED.navigational_status,
        position_accuracy           = EXCLUDED.position_accuracy,
        communication_state         = EXCLUDED.communication_state,
        rate_of_turn                = EXCLUDED.rate_of_turn,
        special_manoeuvre_indicator = EXCLUDED.special_manoeuvre_indicator,
        repeat_indicator            = EXCLUDED.repeat_indicator,
        message_id                  = EXCLUDED.message_id,
        valid                       = EXCLUDED.valid,
        time_utc                    = EXCLUDED.time_utc,
        destination                 = EXCLUDED.destination;

-- name: GetPositionData :many
SELECT *
FROM position_reports
WHERE latitude BETWEEN $1 AND $2 --minLat --maxLat
  AND longitude BETWEEN $3 AND $4;
--minLon --maxLon


-- name: UpdateShipDestination :exec
UPDATE position_reports
SET destination = $2
WHERE ship_name = $1;

-- name: HasImage :one
SELECT
    (image_url <> '') AS has_image
FROM images
WHERE ship_name = $1;

-- name: SetImageForShip :exec
INSERT INTO images(ship_name, image_url)
VALUES ($1, $2);

