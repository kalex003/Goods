package dbmodels

//поч импорт пакета целиком не работает?
import (
	domainmodels "Goods/internal/domain/models"
)

import "time"

type GoodsGetAnswer struct {
	GoodsId      int64     `db:"goods_id"`
	PlaceId      int64     `db:"place_id"`
	SkuId        *int64    `db:"sku_id"`
	WbstickerId  *int64    `db:"wbsticker_id"`
	Barcode      *string   `db:"barcode"`
	StateId      *string   `db:"state_id"`
	ChEmployeeId int64     `db:"ch_employee_id"`
	OfficeId     *int64    `db:"office_id"`
	WhId         *int64    `db:"wh_id"`
	TareId       *int64    `db:"tare_id"`
	TareType     *string   `db:"tare_type"`
	ChDt         time.Time `db:"ch_dt"`
	IsDel        bool      `db:"is_del"`
}

type GoodsGetAnswers struct {
	GoodsGetAnswers []GoodsGetAnswer
}

func ConvertGoodsGetAnswerToGoodFullInfo(good GoodsGetAnswer) domainmodels.GoodFullInfo {
	return domainmodels.GoodFullInfo{
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
		ChDt:         good.ChDt,
		IsDel:        good.IsDel,
	}
}

func ConvertGoodsGetAnswerToGoodsFullInfo(goods GoodsGetAnswers) domainmodels.GoodsFullInfo { //пока буду указатель отдавать
	var infos []domainmodels.GoodFullInfo
	for _, good := range goods.GoodsGetAnswers {
		info := ConvertGoodsGetAnswerToGoodFullInfo(good)
		infos = append(infos, info)
	}

	return domainmodels.GoodsFullInfo{
		GoodsFullInfo: infos,
	}
}
