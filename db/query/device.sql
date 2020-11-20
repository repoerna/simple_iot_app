-- name: GetDevices :many
SELECT * FROM device
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: GetDevice :one
SELECT * FROM device
WHERE id = $1 LIMIT 1;


-- name: CreateDevice :one
INSERT INTO device (
  name,
  shortname,
  enabled,
  createdat,
  updatedat
) VALUES (
  $1, $2, $3, now(), now()
)
RETURNING *;


-- name: UpdateDevice :one
UPDATE device SET 
name = $2,
shortname = $3,
enabled = $4,
updatedat = now()
WHERE id = $1
RETURNING *;

-- name: DeleteDevice :exec
DELETE FROM device
WHERE id = $1;


