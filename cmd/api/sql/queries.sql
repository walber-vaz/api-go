-- name: CreateJob :exec
INSERT INTO jobs (
  company, title, description, location, salary, role, remote
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetJob :one
SELECT * FROM jobs WHERE id = $1;

-- name: GetJobs :many
SELECT * FROM jobs ORDER BY created_at DESC;

-- name: UpdateJob :exec
UPDATE jobs 
  SET company = $2, 
  title = $3,
  description = $4, 
  location = $5,
  salary = $6,
  role = $7,
  remote = $8, 
  updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeleteJob :exec
UPDATE jobs SET deleted_at = NOW() WHERE id = $1;
