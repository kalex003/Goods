package models

import goods1 "github.com/kalex003/Goods_Proto/gen/go/goods"

type GoodsUpdateIsDelInput struct {
	GoodsId int64 `json:"goods_id"`
	IsDel   bool  `json:"is_del"`
}

type GoodsUpdateIsDelInputs struct {
	GoodsUpdateIsDelInputs []GoodsUpdateIsDelInput `json:"data"`
}

// Преобразование одной структуры GoodUpdateIsDelInput в OneUpdateIsDelRequest
func ConvertOneUpdateIsDelRequestToGoodUpdateIsDelInput(request *goods1.OneUpdateIsDelRequest) GoodsUpdateIsDelInput {
	return GoodsUpdateIsDelInput{
		GoodsId: request.GoodsId,
		IsDel:   request.IsDel,
	}
}

// Преобразование массива структур GoodUpdateIsDelInput в UpdateIsDelRequest
func ConvertGoodsUpdateIsDelRequestToUpdateIsDelInput(Inputs *goods1.UpdateIsDelRequest) GoodsUpdateIsDelInputs { //пока буду указатель отдавать
	var requests []GoodsUpdateIsDelInput
	for _, Input := range Inputs.GetStructs() {
		request := ConvertOneUpdateIsDelRequestToGoodUpdateIsDelInput(Input)
		requests = append(requests, request)
	}

	return GoodsUpdateIsDelInputs{
		GoodsUpdateIsDelInputs: requests,
	}
}
