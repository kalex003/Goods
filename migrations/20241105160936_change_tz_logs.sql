-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE goods.goodslog
    ALTER COLUMN log_dt DROP DEFAULT;

ALTER TABLE  goods.goodslog
    ALTER COLUMN log_dt SET DEFAULT CURRENT_TIMESTAMP AT TIME ZONE 'Europe/Moscow';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE goods.goodslog
    ALTER COLUMN log_dt DROP DEFAULT;

ALTER TABLE goods.goodslog
    ALTER COLUMN log_dt SET DEFAULT CURRENT_TIMESTAMP;
-- +goose StatementEnd
