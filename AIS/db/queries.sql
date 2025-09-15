-- name: CreateShip :exec
INSERT INTO ships (mmsi, ship_name)
VALUES ($1, $2);

-- name: GetShip :one
SELECT mmsi, ship_name
FROM ships
WHERE MMSI = $1;

-- name: EmptyDBTables :exec
TRUNCATE TABLE position_reports
RESTART IDENTITY CASCADE;