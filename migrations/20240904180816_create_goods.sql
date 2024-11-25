-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS goods;

CREATE TABLE IF NOT EXISTS goods.goods
(
    goods_id       BIGSERIAL                                             NOT NULL,
    place_id       BIGINT                                                NOT NULL,
    sku_id         BIGINT                                                NULL,
    wbsticker_id   BIGINT                                                NULL,
    barcode        VARCHAR(30)                                           NULL,
    state_id       CHAR(3)                                               NULL,
    ch_employee_id INT                                                   NOT NULL,
    ch_dt          TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    office_id      INT                                                   NULL,
    wh_id          INT                                                   NULL,
    tare_id        BIGINT                                                NULL,
    tare_type      CHAR(3)                                               NULL,
    is_del         BOOLEAN DEFAULT FALSE                                 NOT NULL,
    CONSTRAINT pk_goods PRIMARY KEY (goods_id)
);

--потом сделаю партиционироанную таблицу и крончик
CREATE TABLE IF NOT EXISTS goods.goodslog
(
    log_dt         TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
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
    is_del         BOOLEAN                     NOT NULL
) PARTITION BY RANGE (log_dt);

CREATE INDEX IF NOT EXISTS ix_goodslog_log_dt
    ON goods.goodslog (log_dt);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA IF EXISTS goods CASCADE;
-- +goose StatementEnd

