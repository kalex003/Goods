package models

import goodsv1 "github.com/kalex003/Goods_Proto/gen/go/goods"

type GoodInfo struct {
	PlaceId      int64   `json:"place_id"`
	SkuId        *int64  `json:"sku_id"`
	WbstickerId  *int64  `json:"wbsticker_id"`
	Barcode      *string `json:"barcode"`
	StateId      *string `json:"state_id"`
	ChEmployeeId int64   `json:"ch_employee_id"`
	OfficeId     *int64  `json:"office_id"`
	WhId         *int64  `json:"wh_id"`
	TareId       *int64  `json:"tare_id"`
	TareType     *string `json:"tare_type"`
}

type GoodsInfo struct {
	GoodsInfo []GoodInfo `json:"data"`
}

func NewInfo(PlaceId int64, SkuId *int64, WbstickerId *int64, Barcode *string, StateId *string, ChEmployeeId int64, OfficeId *int64, WhId *int64, TareId *int64, TareType *string) *GoodInfo {
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

func ConvertOneInsertRequestToGoodInfo(req *goodsv1.OneInsertRequest) GoodInfo {
	return GoodInfo{
		PlaceId:      req.PlaceId,
		SkuId:        req.SkuId,
		WbstickerId:  req.WbstickerId,
		Barcode:      req.Barcode,
		StateId:      req.StateId,
		ChEmployeeId: req.ChEmployeeId,
		OfficeId:     req.OfficeId,
		WhId:         req.WhId,
		TareId:       req.TareId,
		TareType:     req.TareType,
	}
}

// Преобразование массива указателей на OneInsertRequest в массив GoodInfo
func ConvertInsertRequestToGoodInfo(reqs *goodsv1.InsertRequest) GoodsInfo {
	var goods []GoodInfo
	for _, req := range reqs.GetStructs() {
		good := ConvertOneInsertRequestToGoodInfo(req)
		goods = append(goods, good)
	}
	return GoodsInfo{
		GoodsInfo: goods,
	}
}
