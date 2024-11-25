-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE import.goods_imported
ALTER COLUMN import_dt DROP DEFAULT;

ALTER TABLE import.goods_imported
    ALTER COLUMN import_dt SET DEFAULT CURRENT_TIMESTAMP AT TIME ZONE 'Europe/Moscow';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE import.goods_imported
    ALTER COLUMN import_dt DROP DEFAULT;

ALTER TABLE import.goods_imported
    ALTER COLUMN import_dt SET DEFAULT CURRENT_TIMESTAMP;
-- +goose StatementEnd
