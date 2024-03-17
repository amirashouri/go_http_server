-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    email,
    age,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- -- name: UpdateUser :one
-- UPDATE users
-- SET
--   hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
--   password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
--   email = COALESCE(sqlc.narg(email), email),
--   is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified),
--   updated_at = COALESCE(sqlc.narg(updated_at), updated_at)
-- WHERE
--   username = sqlc.arg(username)
-- RETURNING *;

-- -- name: DeleteUser :exec
-- DELETE FROM users
-- WHERE id = $1;