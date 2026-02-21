-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS  urls (
    id BIGSERIAL PRIMARY KEY,
    short_code varchar(10) UNIQUE NOT NULL,
    original_url TEXT NOT NULL
);

CREATE INDEX idx_short_code ON urls(short_code);
CREATE INDEX idx_original_url ON urls(original_url);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS urls;
-- +goose StatementEnd
