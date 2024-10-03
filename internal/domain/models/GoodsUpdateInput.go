package models

//поч импорт пакета целиком не работает?
import (
	goodsv1 "github.com/kalex003/Goods_Proto/gen/go/goods"
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

// Преобразование одной структуры OneUpdateRequest в GoodsUpdateInput
func ConvertOneUpdateRequestToGoodsUpdateInput(good *goodsv1.OneUpdateRequest) GoodsUpdateInput {
	return GoodsUpdateInput{
		GoodsId:      good.GoodsId,
		PlaceId:      good.PlaceId,
		SkuId:        good.SkuId,
		WbstickerId:  good.WbstickerId,
		Barcode:      good.Barcode,
		StateId:      good.StateId,
		ChEmployeeId: good.ChEmployeeId,
		OfficeId:     good.OfficeId,
		WhId:         good.WhId,
		TareId:       good.TareId,
		TareType:     good.TareType,
	}
}

// Преобразование массива структур GoodsUpdateInputs в Updateinput
func ConvertUpdateRequestToGoodsUpdateInputs(goods *goodsv1.UpdateRequest) GoodsUpdateInputs { //пока буду указатель отдавать
	var inputs []GoodsUpdateInput
	for _, good := range goods.GetStructs() {
		input := ConvertOneUpdateRequestToGoodsUpdateInput(good)
		inputs = append(inputs, input)
	}

	return GoodsUpdateInputs{
		GoodsUpdateInputs: inputs,
	}
}
