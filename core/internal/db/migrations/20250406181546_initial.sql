-- +goose Up
-- +goose StatementBegin
CREATE TABLE video (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    thumbnail_path TEXT NOT NULL,
    path TEXT NOT NULL,
    duration_seconds INT NOT NULL,
    format TEXT NOT NULL,
    type TEXT CHECK (type IN ('video', 'audio')) NOT NULL,
    tags TEXT[] DEFAULT '{}',
    views INT DEFAULT 0,
    likes INT DEFAULT 0,
    uploaded_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    status TEXT CHECK (status IN ('processing', 'ready', 'failed')) DEFAULT 'processing'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE video;
-- +goose StatementEnd
