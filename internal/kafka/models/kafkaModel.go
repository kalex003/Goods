package models

import "time"

type KafkaModel struct {
	LogDt        time.Time `db:"log_dt"`
	GoodsId      int64     `db:"goods_id"`
	PlaceId      int64     `db:"place_id"`
	SkuId        *int64    `db:"sku_id"`
	WbstickerId  *int64    `db:"wbsticker_id"`
	Barcode      *string   `db:"barcode"`
	StateId      *string   `db:"state_id"`
	ChEmployeeId int64     `db:"ch_employee_id"`
	OfficeId     *int64    `db:"office_id"`
	WhId         *int64    `db:"wh_id"`
	TareId       *int64    `db:"tare_id"`
	TareType     *string   `db:"tare_type"`
	ChDt         time.Time `db:"ch_dt"`
	IsDel        bool      `db:"is_del"`
}
