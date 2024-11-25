-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS import;

--потом сделаю  таблицу и крончик
CREATE TABLE IF NOT EXISTS import.goods_imported
(
    import_dt      TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    goods_id       BIGINT                      NOT NULL,
    place_id       BIGINT                      NOT NULL,
    sku_id         BIGINT                      NULL,
    wbsticker_id   BIGINT                      NULL,
    barcode        VARCHAR(30)                 NULL,
    state_id       CHAR(3)                     NULL,
    ch_employee_id INT                         NOT NULL,
    ch_dt          TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    office_id      INT                         NULL,
    wh_id          INT                         NULL,
    tare_id        BIGINT                      NULL,
    tare_type      CHAR(3)                     NULL,
    is_del         BOOLEAN                    NOT NULL
    )
CREATE INDEX IF NOT EXISTS ix_goodsimported_import_dt
    ON goods.goodslog (import_dt);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
