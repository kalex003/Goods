package dbmodels

import (
	domainmodels "Goods/internal/domain/models"
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

func ConvertGoodUpdateIsDelAnswerToGoodFullInfo(answer GoodsUpdateIsDelAnswer) domainmodels.GoodFullInfo {
	return domainmodels.GoodFullInfo{
		GoodsId: answer.GoodsId,
		ChDt:    answer.ChDt,
		IsDel:   answer.IsDel,
	}
}

func ConvertGoodsUpdateIsDelAnswerToGoodsFullInfo(answers GoodsUpdateIsDelAnswers) domainmodels.GoodsFullInfo { //пока буду указатель отдавать
	var infos []domainmodels.GoodFullInfo
	for _, answer := range answers.GoodsUpdateIsDelAnswers {
		info := ConvertGoodUpdateIsDelAnswerToGoodFullInfo(answer)
		infos = append(infos, info)
	}

	return domainmodels.GoodsFullInfo{
		GoodsFullInfo: infos,
	}
}
