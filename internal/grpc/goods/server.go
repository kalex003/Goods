package grpcgoods

import (
	domainmodels "Goods/internal/domain/models"
	"context"
	goodsv1 "github.com/kalex003/Goods_Proto/gen/go/goods"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Goods interface { // тут основная логика будет (создали интерфейс, в services.goods.goods.go будет структура Goods, которая будет реализоывать контракт)
	InsertNewGoods(context.Context, domainmodels.GoodsFullInfo) (domainmodels.GoodsFullInfo, error)
	GetByIdsGoodsInfo(context.Context, *[]int64) (domainmodels.GoodsFullInfo, error)
	ChangeGoods(context.Context, domainmodels.GoodsFullInfo) (domainmodels.GoodsFullInfo, error)
	ChangeIsDelOfGoods(context.Context, domainmodels.GoodsFullInfo) (domainmodels.GoodsFullInfo, error)
	GetByPlaceGoodsInfo(context.Context, int64) (domainmodels.GoodsFullInfo, error)
	GetByTareGoodsInfo(context.Context, int64) (domainmodels.GoodsFullInfo, error)
	GetGoodsHistory(context.Context, int64) (domainmodels.GoodsFullInfo, error)
}

type serverAPI struct { //для реализация интерфейсов, сгенерированнхы прото файлом
	goodsv1.UnimplementedGoodsServer //А зачем вот это используется? (сейчас нахожу это странным)
	goods                            Goods
}

func Register(gRPCServer *grpc.Server, goods Goods) { //создаем структуру выше
	goodsv1.RegisterGoodsServer(gRPCServer, &serverAPI{goods: goods})
}

func (s *serverAPI) Insert(ctx context.Context, req *goodsv1.InsertRequest) (*goodsv1.InsertResponse, error) {

	//	err := ValidateInsert(req)

	/*	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}*/

	ans, err := s.goods.InsertNewGoods(ctx, domainmodels.ConvertInsertRequestToGoodFullInfo(req))

	if err != nil {
		return nil, err
	}

	return domainmodels.ConvertGoodsFullInfoToInsertResponse(ans), nil

}

/*func ValidateInsert(req *goodsv1.InsertRequest) error {

	if req.GetPlaceId() <= 0 {
		return status.Error(codes.InvalidArgument, "missing place_id")
	}

	if req.GetSkuId() <= 0 {
		return status.Error(codes.InvalidArgument, "missing sku_id")
	}

	if req.GetWbstickerId() <= 0 {
		return status.Error(codes.InvalidArgument, "missing wbsticker_id")
	}

	if req.GetBarcode() == "" {
		return status.Error(codes.InvalidArgument, "missing barcode")
	}

	if req.GetStateId() == "" {
		return status.Error(codes.InvalidArgument, "missing state_id")
	}

	if req.GetChEmployeeId() <= 0 {
		return status.Error(codes.InvalidArgument, "missing ch_employee_id")
	}

	if req.GetOfficeId() <= 0 {
		return status.Error(codes.InvalidArgument, "missing office_id")
	}

	if req.GetWhId() <= 0 {
		return status.Error(codes.InvalidArgument, "missing wh_id")
	}

	if req.GetTareId() <= 0 {
		return status.Error(codes.InvalidArgument, "missing tare_id")
	}

	if req.GetTareType() == "" {
		return status.Error(codes.InvalidArgument, "missing tare_type")
	}

	return nil
}*/

func (s *serverAPI) GetById(ctx context.Context, req *goodsv1.GetByIdRequest) (*goodsv1.GetResponse, error) {

	for _, v := range req.GetStructs() {
		if err := ValidateGetById(v.GoodsId); err != nil {
			return nil, err
		}
	}

	goodsIds := make([]int64, len(req.GetStructs()))

	// Заполняем массив значениями GoodsId
	for i, item := range req.GetStructs() {
		goodsIds[i] = *item.GoodsId
	}

	resp, err := s.goods.GetByIdsGoodsInfo(ctx, &goodsIds)

	if err != nil {
		return nil, err
	}

	return domainmodels.ConvertGoodsFullInfoToGetResponse(resp), nil
}

func ValidateGetById(goodsId *int64) error {
	if *goodsId <= 0 {
		return status.Error(codes.InvalidArgument, "missing goods_id")
	}

	return nil
}

func (s *serverAPI) GetByPlace(ctx context.Context, req *goodsv1.GetByPlaceRequest) (*goodsv1.GetResponse, error) {

	placeId := req.GetPlaceId()
	if err := ValidateGetByPlace(placeId); err != nil {
		return nil, err
	}

	resp, err := s.goods.GetByPlaceGoodsInfo(ctx, placeId)

	if err != nil {
		return nil, err
	}

	return domainmodels.ConvertGoodsFullInfoToGetResponse(resp), nil
}

func ValidateGetByPlace(placeId int64) error {
	if placeId <= 0 {
		return status.Error(codes.InvalidArgument, "wrong place_id")
	}

	return nil

}

func (s *serverAPI) GetByTare(ctx context.Context, req *goodsv1.GetByTareRequest) (*goodsv1.GetResponse, error) {

	tareId := req.GetTareId()
	if err := ValidateGetByTare(tareId); err != nil {
		return nil, err
	}

	resp, err := s.goods.GetByTareGoodsInfo(ctx, tareId)

	if err != nil {
		return nil, err
	}

	return domainmodels.ConvertGoodsFullInfoToGetResponse(resp), nil
}

func ValidateGetByTare(tareId int64) error {
	if tareId <= 0 {
		return status.Error(codes.InvalidArgument, "wrong tare_id")
	}

	return nil

}

func (s *serverAPI) GetHistory(ctx context.Context, req *goodsv1.OneGetByIdRequest) (*goodsv1.GetResponse, error) {

	goodsId := req.GetGoodsId()
	if err := ValidateGetHistory(goodsId); err != nil {
		return nil, err
	}

	resp, err := s.goods.GetGoodsHistory(ctx, goodsId)

	if err != nil {
		return nil, err
	}

	return domainmodels.ConvertGoodsFullInfoToGetResponse(resp), nil
}

func ValidateGetHistory(goodsId int64) error {
	if goodsId <= 0 {
		return status.Error(codes.InvalidArgument, "missing goods_id")
	}

	return nil
}

func (s *serverAPI) Update(ctx context.Context, req *goodsv1.UpdateRequest) (*goodsv1.UpdateResponse, error) {
	/*	if err := ValidateUpdate(req); err != nil {
		return nil, err
	}*/

	ans, err := s.goods.ChangeGoods(ctx, domainmodels.ConvertUpdateRequestToGoodsFullInfo(req))

	if err != nil {
		return nil, err
	}

	return domainmodels.ConvertGoodsFullInfoToUpdateResponse(ans), nil
}

/*
func ValidateUpdate(req *goodsv1.UpdateRequest) error {

	if req.GetGoodsId() == 0 {
		return status.Error(codes.InvalidArgument, "missing goods_id")
	}

	if req.GetPlaceId() == 0 {
		return status.Error(codes.InvalidArgument, "missing place_id")
	}

	if req.GetSkuId() == 0 {
		return status.Error(codes.InvalidArgument, "missing sku_id")
	}

	if req.GetWbstickerId() == 0 {
		return status.Error(codes.InvalidArgument, "missing wbsticker_id")
	}

	if req.GetBarcode() == "" {
		return status.Error(codes.InvalidArgument, "missing barcode")
	}

	if req.GetStateId() == "" {
		return status.Error(codes.InvalidArgument, "missing state_id")
	}

	if req.GetChEmployeeId() == 0 {
		return status.Error(codes.InvalidArgument, "missing ch_employee_id")
	}

	if req.GetOfficeId() == 0 {
		return status.Error(codes.InvalidArgument, "missing office_id")
	}

	if req.GetWhId() == 0 {
		return status.Error(codes.InvalidArgument, "missing wh_id")
	}

	if req.GetTareId() == 0 {
		return status.Error(codes.InvalidArgument, "missing tare_id")
	}

	if req.GetTareType() == "" {
		return status.Error(codes.InvalidArgument, "missing tare_type")
	}

	return nil
}*/

func (s *serverAPI) UpdateIsDel(ctx context.Context, req *goodsv1.UpdateIsDelRequest) (*goodsv1.UpdateIsDelResponse, error) {

	for _, v := range req.GetStructs() {
		if err := ValidateUpdateIsDel(v.GoodsId); err != nil {
			return nil, err
		}
	}

	ans, err := s.goods.ChangeIsDelOfGoods(ctx, domainmodels.ConvertUpdateIsDelRequestToGoodsFullInfo(req))

	if err != nil {
		return nil, err
	}

	return domainmodels.ConvertGoodsFullInfoToUpdateIsDelResponse(ans), nil

}

func ValidateUpdateIsDel(goodsId int64) error {

	if goodsId <= 0 {
		return status.Error(codes.InvalidArgument, "missing goods_id")
	}

	return nil

}
