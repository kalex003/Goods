package domain

//поч импорт пакета целиком не работает?
import (
	goodsv1 "github.com/kalex003/Goods_Proto/gen/go/goods"
	"google.golang.org/protobuf/types/known/timestamppb"
)

import "time"

type GoodFullInfo struct {
	GoodsId      int64
	PlaceId      int64
	SkuId        *int64
	WbstickerId  *int64
	Barcode      *string
	StateId      *string
	ChEmployeeId int64
	OfficeId     *int64
	WhId         *int64
	TareId       *int64
	TareType     *string
	ChDt         time.Time
	IsDel        bool
}

type GoodsFullInfo struct {
	GoodsFullInfo []GoodFullInfo
}

// GetResponse
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

// InsertResponse
func ConvertGoodFullInfoToOneInsertResponse(good GoodFullInfo) *goodsv1.OneInsertResponse {
	return &goodsv1.OneInsertResponse{
		GoodsId: good.GoodsId,
		ChDt:    timestamppb.New(good.ChDt),
	}
}

func ConvertGoodsFullInfoToInsertResponse(goods GoodsFullInfo) *goodsv1.InsertResponse {
	var responses []*goodsv1.OneInsertResponse
	for _, good := range goods.GoodsFullInfo {
		response := ConvertGoodFullInfoToOneInsertResponse(good)
		responses = append(responses, response)
	}

	return &goodsv1.InsertResponse{
		Structs: responses,
	}
}

func ConvertOneInsertRequestToGoodFullInfo(req *goodsv1.OneInsertRequest) GoodFullInfo {
	return GoodFullInfo{
		PlaceId:      req.PlaceId,
		SkuId:        req.SkuId,
		WbstickerId:  req.WbstickerId,
		Barcode:      req.Barcode,
		StateId:      req.StateId,
		ChEmployeeId: req.ChEmployeeId,
		OfficeId:     req.OfficeId,
		WhId:         req.WhId,
		TareId:       req.TareId,
		TareType:     req.TareType,
	}
}

// Преобразование массива указателей на OneInsertRequest в массив GoodInfo
func ConvertInsertRequestToGoodFullInfo(reqs *goodsv1.InsertRequest) GoodsFullInfo {
	var goods []GoodFullInfo
	for _, req := range reqs.GetStructs() {
		good := ConvertOneInsertRequestToGoodFullInfo(req)
		goods = append(goods, good)
	}
	return GoodsFullInfo{
		GoodsFullInfo: goods,
	}
}

// UpdateResponse
func ConvertGoodFullInfoToOneUpdateResponse(good GoodFullInfo) *goodsv1.OneUpdateResponse {
	return &goodsv1.OneUpdateResponse{
		GoodsId: good.GoodsId,
		ChDt:    timestamppb.New(good.ChDt),
	}
}

func ConvertGoodsFullInfoToUpdateResponse(goods GoodsFullInfo) *goodsv1.UpdateResponse {
	var responses []*goodsv1.OneUpdateResponse
	for _, good := range goods.GoodsFullInfo {
		response := ConvertGoodFullInfoToOneUpdateResponse(good)
		responses = append(responses, response)
	}

	return &goodsv1.UpdateResponse{
		Structs: responses,
	}
}

func ConvertOneUpdateRequestToGoodFullInfo(req *goodsv1.OneUpdateRequest) GoodFullInfo {
	return GoodFullInfo{
		PlaceId:      req.PlaceId,
		SkuId:        req.SkuId,
		WbstickerId:  req.WbstickerId,
		Barcode:      req.Barcode,
		StateId:      req.StateId,
		ChEmployeeId: req.ChEmployeeId,
		OfficeId:     req.OfficeId,
		WhId:         req.WhId,
		TareId:       req.TareId,
		TareType:     req.TareType,
	}
}

func ConvertUpdateRequestToGoodsFullInfo(reqs *goodsv1.UpdateRequest) GoodsFullInfo {
	var goods []GoodFullInfo
	for _, req := range reqs.GetStructs() {
		good := ConvertOneUpdateRequestToGoodFullInfo(req)
		goods = append(goods, good)
	}
	return GoodsFullInfo{
		GoodsFullInfo: goods,
	}
}

// UpdateIsDelResponse
func ConvertGoodFullInfoToOneUpdateIsDelResponse(good GoodFullInfo) *goodsv1.OneUpdateIsDelResponse {
	return &goodsv1.OneUpdateIsDelResponse{
		GoodsId: good.GoodsId,
		ChDt:    timestamppb.New(good.ChDt),
		IsDel:   good.IsDel,
	}
}

func ConvertGoodsFullInfoToUpdateIsDelResponse(goods GoodsFullInfo) *goodsv1.UpdateIsDelResponse {
	var responses []*goodsv1.OneUpdateIsDelResponse
	for _, good := range goods.GoodsFullInfo {
		response := ConvertGoodFullInfoToOneUpdateIsDelResponse(good)
		responses = append(responses, response)
	}

	return &goodsv1.UpdateIsDelResponse{
		Structs: responses,
	}
}

func ConvertOneUpdateIsDelRequestToGoodFullInfo(req *goodsv1.OneUpdateIsDelRequest) GoodFullInfo {
	return GoodFullInfo{
		GoodsId: req.GoodsId,
		IsDel:   req.IsDel,
	}
}

func ConvertUpdateIsDelRequestToGoodsFullInfo(reqs *goodsv1.UpdateIsDelRequest) GoodsFullInfo {
	var goods []GoodFullInfo
	for _, req := range reqs.GetStructs() {
		good := ConvertOneUpdateIsDelRequestToGoodFullInfo(req)
		goods = append(goods, good)
	}
	return GoodsFullInfo{
		GoodsFullInfo: goods,
	}
}
