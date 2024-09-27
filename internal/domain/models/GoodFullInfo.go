package models

//поч импорт пакета целиком не работает?
import (
	goodsv1 "github.com/kalex003/Goods_Proto/gen/go/goods"
	"google.golang.org/protobuf/types/known/timestamppb"
)

import "time"

type GoodFullInfo struct {
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

type GoodsFullInfo struct {
	GoodsFullInfo []GoodFullInfo
}

func NewFullInfo(GoodsId int64, PlaceId int64, SkuId *int64, WbstickerId *int64, Barcode *string, StateId *string, ChEmployeeId int64, OfficeId *int64, WhId *int64, TareId *int64, TareType *string, ChDt time.Time, IsDel bool) *GoodFullInfo {
	return &GoodFullInfo{
		GoodsId:      GoodsId,
		PlaceId:      PlaceId,
		SkuId:        SkuId,
		WbstickerId:  WbstickerId,
		Barcode:      Barcode,
		StateId:      StateId,
		ChEmployeeId: ChEmployeeId,
		OfficeId:     OfficeId,
		WhId:         WhId,
		TareId:       TareId,
		TareType:     TareType,
		ChDt:         ChDt,
		IsDel:        IsDel,
	}
}

// Преобразование одной структуры GoodFullInfo в OneGetResponse
func ConvertGoodFullInfoToOneGetResponse(good GoodFullInfo) *goodsv1.OneGetResponse {
	return &goodsv1.OneGetResponse{
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
		ChDt:         timestamppb.New(good.ChDt),
		IsDel:        good.IsDel,
	}
}

// Преобразование массива структур GoodFullInfo в массив GetResponse
func ConvertGoodsFullInfoToGetResponse(goods GoodsFullInfo) *goodsv1.GetResponse { //пока буду указатель отдавать
	var responses []*goodsv1.OneGetResponse
	for _, good := range goods.GoodsFullInfo {
		response := ConvertGoodFullInfoToOneGetResponse(good)
		responses = append(responses, response)
	}

	return &goodsv1.GetResponse{
		Structs: responses,
	}
}

type GoodInsertAnswer struct {
	GoodsId int64     `db:"goods_id"`
	ChDt    time.Time `db:"ch_dt"`
}

type GoodsInsertAnswer struct {
	GoodsInsertAnswer []GoodInsertAnswer
}

// Преобразование одной структуры GoodInsertAnswer в OneInsertResponse
func ConvertGoodInsertAnswerToOneInsertResponse(answer GoodInsertAnswer) *goodsv1.OneInsertResponse {
	return &goodsv1.OneInsertResponse{
		GoodsId: answer.GoodsId,
		ChDt:    timestamppb.New(answer.ChDt),
	}
}

func ConvertGoodsInsertAnswerToInsertResponse(answers GoodsInsertAnswer) *goodsv1.InsertResponse { //пока буду указатель отдавать
	var responses []*goodsv1.OneInsertResponse
	for _, answer := range answers.GoodsInsertAnswer {
		response := ConvertGoodInsertAnswerToOneInsertResponse(answer)
		responses = append(responses, response)
	}

	return &goodsv1.InsertResponse{
		Structs: responses,
	}
}
