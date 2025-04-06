-- name: CreateVideo :one
INSERT INTO video (
    title,
    description,
    thumbnail_path,
    path,
    duration_seconds,
    format,
    type,
    tags
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;