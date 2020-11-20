-- name: GetTMByID :one
SELECT * FROM telemetry
WHERE id = $1
LIMIT 1;

-- name: GetTMByDeviceID :many
SELECT * FROM telemetry
WHERE deviceid = $1
ORDER BY createdat
LIMIT $2
OFFSET $3;

-- name: CreateTM :one
INSERT INTO telemetry (
    deviceid,
    createdat,
    latitude,
    longitude,
    value,
    value2,
    value3,
    value4
) VALUES (
    $1, now(), $2, $3, $4, $5, $6, $7
)
RETURNING *;


