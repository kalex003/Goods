package dbmodels

//поч импорт пакета целиком не работает?
import (
	domainmodels "Goods/internal/domain/models"
)

type GoodsUpdateInput struct {
	GoodsId      int64   `json:"goods_id"`
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

type GoodsUpdateInputs struct {
	GoodsUpdateInputs []GoodsUpdateInput `json:"data"`
}

func ConvertGoodsFullInfoToGoodsUpdateInput(info domainmodels.GoodFullInfo) GoodsUpdateInput {
	return GoodsUpdateInput{
		GoodsId:      info.GoodsId,
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

func ConvertGoodsFullInfoToGoodsUpdateInputs(infos domainmodels.GoodsFullInfo) GoodsUpdateInputs {
	var goods []GoodsUpdateInput
	for _, info := range infos.GoodsFullInfo {
		good := ConvertGoodsFullInfoToGoodsUpdateInput(info)
		goods = append(goods, good)
	}
	return GoodsUpdateInputs{
		GoodsUpdateInputs: goods,
	}
}
