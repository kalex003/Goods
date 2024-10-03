package models

import (
	goodsv1 "github.com/kalex003/Goods_Proto/gen/go/goods"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type GoodsUpdateIsDelAnswer struct {
	GoodsId int64     `db:"goods_id"`
	ChDt    time.Time `db:"ch_dt"`
	IsDel   bool      `db:"is_del"`
}

type GoodsUpdateIsDelAnswers struct {
	GoodsUpdateIsDelAnswers []GoodsUpdateIsDelAnswer
}

// Преобразование одной структуры GoodUpdateIsDelAnswer в OneUpdateIsDelResponse
func ConvertGoodUpdateIsDelAnswerToOneUpdateIsDelResponse(answer GoodsUpdateIsDelAnswer) *goodsv1.OneUpdateIsDelResponse {
	return &goodsv1.OneUpdateIsDelResponse{
		GoodsId: answer.GoodsId,
		ChDt:    timestamppb.New(answer.ChDt),
		IsDel:   answer.IsDel,
	}
}

// Преобразование массива структур GoodUpdateIsDelAnswer в UpdateIsDelResponse
func ConvertGoodsUpdateIsDelAnswerToUpdateIsDelResponse(answers GoodsUpdateIsDelAnswers) *goodsv1.UpdateIsDelResponse { //пока буду указатель отдавать
	var responses []*goodsv1.OneUpdateIsDelResponse
	for _, answer := range answers.GoodsUpdateIsDelAnswers {
		response := ConvertGoodUpdateIsDelAnswerToOneUpdateIsDelResponse(answer)
		responses = append(responses, response)
	}

	return &goodsv1.UpdateIsDelResponse{
		Structs: responses,
	}
}
