-- name: CreateUserAddress :one
INSERT INTO user_address(
   user_id,
   address_line1,
   address_line2,
   city,
   postal_code,
   country,
   phone_number,
   created_at,
   modified_at
   ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING *;

-- name: DeleteUserAddress :one
DELETE 
FROM user_address
WHERE id = $1 RETURNING *;

-- name: ListUserAddresses :many
SELECT * 
FROM user_address
WHERE user_id = $1;

