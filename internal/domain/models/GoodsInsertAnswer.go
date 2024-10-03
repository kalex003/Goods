package models

import (
	goodsv1 "github.com/kalex003/Goods_Proto/gen/go/goods"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type GoodsInsertAnswer struct {
	GoodsId int64     `db:"goods_id"`
	ChDt    time.Time `db:"ch_dt"`
}

type GoodsInsertAnswers struct {
	GoodsInsertAnswers []GoodsInsertAnswer
}

// Преобразование одной структуры GoodInsertAnswer в OneInsertResponse
func ConvertGoodInsertAnswerToOneInsertResponse(answer GoodsInsertAnswer) *goodsv1.OneInsertResponse {
	return &goodsv1.OneInsertResponse{
		GoodsId: answer.GoodsId,
		ChDt:    timestamppb.New(answer.ChDt),
	}
}

// Преобразование массива структур GoodInsertAnswer в InsertResponse
func ConvertGoodsInsertAnswerToInsertResponse(answers GoodsInsertAnswers) *goodsv1.InsertResponse { //пока буду указатель отдавать
	var responses []*goodsv1.OneInsertResponse
	for _, answer := range answers.GoodsInsertAnswers {
		response := ConvertGoodInsertAnswerToOneInsertResponse(answer)
		responses = append(responses, response)
	}

	return &goodsv1.InsertResponse{
		Structs: responses,
	}
}
