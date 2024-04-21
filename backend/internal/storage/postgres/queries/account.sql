-- name: CreateAccount :one
INSERT INTO accounts (
  number, 
  password,
  first_name,
  name,
  last_name,
  balance
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetAccountByID :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountByNumber :one
SELECT * FROM accounts
WHERE number = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccountBalance :one
UPDATE accounts
  set balance = $2
WHERE id = $1
RETURNING *;

-- name: UpdateAccountPassword :one
UPDATE accounts
  set password = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;