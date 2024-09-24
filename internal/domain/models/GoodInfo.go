package models

type GoodInfo struct {
	PlaceId      int64  `db:"place_id"`
	SkuId        int64  `db:"sku_id"`
	WbstickerId  int64  `db:"wbsticker_id"`
	Barcode      string `db:"barcode"`
	StateId      string `db:"state_id"`
	ChEmployeeId int64  `db:"ch_employee_id"`
	OfficeId     int64  `db:"office_id"`
	WhId         int64  `db:"wh_id"`
	TareId       int64  `db:"tare_id"`
	TareType     string `db:"tare_type"`
}

func NewInfo(PlaceId int64, SkuId int64, WbstickerId int64, Barcode string, StateId string, ChEmployeeId int64, OfficeId int64, WhId int64, TareId int64, TareType string) *GoodInfo {
	return &GoodInfo{
		PlaceId:      PlaceId,
		SkuId:        SkuId,
		WbstickerId:  WbstickerId,
		Barcode:      Barcode,
		StateId:      StateId,
		ChEmployeeId: ChEmployeeId,
		OfficeId:     OfficeId,
		WhId:         WhId,
		TareId:       TareId,
		TareType:     TareType,
	}
}
