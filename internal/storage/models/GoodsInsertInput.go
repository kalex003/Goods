package dbmodels

import (
	domainmodels "Goods/internal/domain/models"
)

type GoodsInsertInput struct {
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

type GoodsInsertInputs struct {
	GoodsInsertInputs []GoodsInsertInput `json:"data"`
}

func ConvertGoodFullInfoToGoodsInsertInput(info domainmodels.GoodFullInfo) GoodsInsertInput {
	return GoodsInsertInput{
		PlaceId:      info.PlaceId,
		SkuId:        info.SkuId,
		WbstickerId:  info.WbstickerId,
		Barcode:      info.Barcode,
		StateId:      info.StateId,
		ChEmployeeId: info.ChEmployeeId,
		OfficeId:     info.OfficeId,
		WhId:         info.WhId,
		TareId:       info.TareId,
		TareType:     info.TareType,
	}
}

func ConvertGoodsFullInfoToGoodsInsertInputs(infos domainmodels.GoodsFullInfo) GoodsInsertInputs {
	var goods []GoodsInsertInput
	for _, info := range infos.GoodsFullInfo {
		good := ConvertGoodFullInfoToGoodsInsertInput(info)
		goods = append(goods, good)
	}
	return GoodsInsertInputs{
		GoodsInsertInputs: goods,
	}
}
