-- name: CreatePayment :one
INSERT INTO payments (
  account_id,
  amount,
  place
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetPayment :one
SELECT * FROM payments
WHERE id = $1 LIMIT 1;

-- name: ListPayments :many
SELECT * FROM payments
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;