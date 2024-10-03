package models

//поч импорт пакета целиком не работает?
import (
	goodsv1 "github.com/kalex003/Goods_Proto/gen/go/goods"
	"google.golang.org/protobuf/types/known/timestamppb"
)

import "time"

type GoodsUpdateAnswer struct {
	GoodsId int64     `db:"goods_id"`
	ChDt    time.Time `db:"ch_dt"`
}

type GoodsUpdateAnswers struct {
	GoodsUpdateAnswers []GoodsUpdateAnswer
}

// Преобразование одной структуры GoodsUpdateAnswer в OneUpdateResponse
func ConvertGoodsUpdateAnswerToOneUpdateResponse(good GoodsUpdateAnswer) *goodsv1.OneUpdateResponse {
	return &goodsv1.OneUpdateResponse{
		GoodsId: good.GoodsId,
		ChDt:    timestamppb.New(good.ChDt),
	}
}

// Преобразование массива структур GoodsUpdateAnswers в UpdateResponse
func ConvertGoodsUpdateAnswersToUpdateResponse(goods GoodsUpdateAnswers) *goodsv1.UpdateResponse { //пока буду указатель отдавать
	var responses []*goodsv1.OneUpdateResponse
	for _, good := range goods.GoodsUpdateAnswers {
		response := ConvertGoodsUpdateAnswerToOneUpdateResponse(good)
		responses = append(responses, response)
	}

	return &goodsv1.UpdateResponse{
		Structs: responses,
	}
}
