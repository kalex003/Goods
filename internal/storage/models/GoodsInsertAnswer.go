package dbmodels

import (
	domainmodels "Goods/internal/domain/models"
	"time"
)

type GoodsInsertAnswer struct {
	GoodsId int64     `db:"goods_id"`
	ChDt    time.Time `db:"ch_dt"`
}

type GoodsInsertAnswers struct {
	GoodsInsertAnswers []GoodsInsertAnswer
}

func ConvertGoodsInsertAnswerToGoodFullInfo(good GoodsInsertAnswer) domainmodels.GoodFullInfo {
	return domainmodels.GoodFullInfo{
		GoodsId: good.GoodsId,
		ChDt:    good.ChDt,
	}
}

func ConvertGoodsinsertAnswersToGoodsFullInfo(goods GoodsInsertAnswers) domainmodels.GoodsFullInfo { //пока буду указатель отдавать
	var infos []domainmodels.GoodFullInfo
	for _, good := range goods.GoodsInsertAnswers {
		info := ConvertGoodsInsertAnswerToGoodFullInfo(good)
		infos = append(infos, info)
	}

	return domainmodels.GoodsFullInfo{
		GoodsFullInfo: infos,
	}
}
