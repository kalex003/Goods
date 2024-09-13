package models

type GoodFullInfo struct {
	GoodsId      int64
	PlaceId      int64
	SkuId        int64
	WbstickerId  int64
	Barcode      string
	StateId      string
	ChEmployeeId int64
	OfficeId     int64
	WhId         int64
	TareId       int64
	TareType     string
}

func NewFullInfo(GoodsId int64, PlaceId int64, SkuId int64, WbstickerId int64, Barcode string, StateId string, ChEmployeeId int64, OfficeId int64, WhId int64, TareId int64, TareType string) *GoodFullInfo {
	return &GoodFullInfo{
		GoodsId:      GoodsId,
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
