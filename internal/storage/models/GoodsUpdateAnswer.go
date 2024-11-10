package dbmodels

import (
	domainmodels "Goods/internal/domain/models"
	"time"
)

type GoodsUpdateAnswer struct {
	GoodsId int64     `db:"goods_id"`
	ChDt    time.Time `db:"ch_dt"`
}

type GoodsUpdateAnswers struct {
	GoodsUpdateAnswers []GoodsUpdateAnswer
}

/*
func ConvertGoodUpdateAnswerToOneUpdateResponse(answer GoodUpdateAnswer) *goodsv1.OneUpdateResponse {
	return &goodsv1.OneUpdateResponse{
		GoodsId: good.GoodsId,
		ChDt:    timestamppb.New(good.ChDt),
	}
}


func ConvertGoodsUpdateAnswerToUpdateResponse(answers GoodUpdateAnswers) *goodsv1.UpdateResponse { //пока буду указатель отдавать
	var responses []*goodsv1.OneUpdateResponse
	for _, answer := range answers.GoodsUpdateAnswer {
		response := ConvertGoodsUpdateAnswerToOneUpdateResponse(answer)
		responses = append(responses, response)
	}

	return &goodsv1.UpdateResponse{
		Structs: responses,
	}
}
*/

func ConvertGoodsUpdateAnswerToGoodFullInfo(answer GoodsUpdateAnswer) domainmodels.GoodFullInfo {
	return domainmodels.GoodFullInfo{
		GoodsId: answer.GoodsId,
		ChDt:    answer.ChDt,
	}
}

func ConvertGoodsUpdateAnswerToGoodsFullInfo(answers GoodsUpdateAnswers) domainmodels.GoodsFullInfo { //пока буду указатель отдавать
	var infos []domainmodels.GoodFullInfo
	for _, answer := range answers.GoodsUpdateAnswers {
		info := ConvertGoodsUpdateAnswerToGoodFullInfo(answer)
		infos = append(infos, info)
	}

	return domainmodels.GoodsFullInfo{
		GoodsFullInfo: infos,
	}
}
