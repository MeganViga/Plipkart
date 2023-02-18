-- name: CreateUser :one
INSERT INTO user_data(
   username,
   password,
   first_name,
   last_name,
   role,
   phone_number,
   created_at,
   modified_at
   ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING *;

-- name: GetUserByName :one
SELECT *
FROM user_data
WHERE username = $1;