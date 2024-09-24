package grpcgoods

import (
	"Goods/internal/domain/models"
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	goodsv1 "github.com/kalex003/Goods_Proto/gen/go/goods"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Goods interface { // тут основная логика будет (создали интерфейс, в services.goods.goods.go будет структура Goods, которая будет реализоывать контракт)
	InsertNewGood(context.Context, models.GoodInfo) (int64, *timestamp.Timestamp, error)
	GetGoodInfo(context.Context, int64) (models.GoodFullInfo, error)
	ChangeGood(context.Context, models.GoodFullInfo) error
	RemoveGood(context.Context, int64) error
}

type serverAPI struct { //для реализация интерфейсов, сгенерированнхы прото файлом
	goodsv1.UnimplementedGoodsServer //метод, сгенерированный grpc
	goods                            Goods
}

func Register(gRPCServer *grpc.Server, goods Goods) { //создаем структуру выше
	goodsv1.RegisterGoodsServer(gRPCServer, &serverAPI{goods: goods})
}

func (s *serverAPI) Insert(ctx context.Context, req *goodsv1.InsertRequest) (*goodsv1.InsertResponse, error) {

	err := ValidateInsert(req)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	goodID, chDt, err := s.goods.InsertNewGood(ctx, *models.NewInfo(req.GetPlaceId(), req.GetSkuId(), req.GetWbstickerId(), req.GetBarcode(), req.GetStateId(), req.GetChEmployeeId(), req.GetOfficeId(), req.GetWhId(), req.GetTareId(), req.GetTareType()))

	if err != nil {
		return nil, err
	}

	return &goodsv1.InsertResponse{
		GoodsId: goodID,
		ChDt:    chDt, // поч grpc заставляет меня использовать указатель
	}, nil

}

func ValidateInsert(req *goodsv1.InsertRequest) error {

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
}

func (s *serverAPI) Get(ctx context.Context, req *goodsv1.GetRequest) (*goodsv1.GetResponse, error) {
	if err := ValidateGet(req); err != nil {
		return nil, err
	}

	resp, err := s.goods.GetGoodInfo(ctx, req.GetGoodsId())

	if err != nil {
		return nil, err
	}

	return &goodsv1.GetResponse{
		GoodsId:      resp.GoodsId,
		PlaceId:      resp.PlaceId,
		SkuId:        resp.SkuId,
		WbstickerId:  resp.WbstickerId,
		Barcode:      resp.Barcode,
		StateId:      resp.StateId,
		ChEmployeeId: resp.ChEmployeeId,
		OfficeId:     resp.OfficeId,
		WhId:         resp.WhId,
		TareId:       resp.TareId,
		TareType:     resp.TareType,
		ChDt:         timestamppb.New(resp.ChDt),
	}, nil
}

func ValidateGet(req *goodsv1.GetRequest) error {
	if req.GetGoodsId() == 0 {
		return status.Error(codes.InvalidArgument, "missing goods_id")
	}

	return nil

}

func (s *serverAPI) Update(ctx context.Context, req *goodsv1.UpdateRequest) (*goodsv1.UpdateResponse, error) {
	if err := ValidateUpdate(req); err != nil {
		return nil, err
	}

	err := s.goods.ChangeGood(ctx, models.GoodFullInfo{
		GoodsId:      req.GetGoodsId(),
		PlaceId:      req.GetPlaceId(),
		SkuId:        req.GetSkuId(),
		WbstickerId:  req.GetWbstickerId(),
		Barcode:      req.GetBarcode(),
		StateId:      req.GetStateId(),
		ChEmployeeId: req.GetChEmployeeId(),
		OfficeId:     req.GetOfficeId(),
		WhId:         req.GetWhId(),
		TareId:       req.GetTareId(),
		TareType:     req.GetTareType(),
	})

	if err != nil {
		return nil, err
	}

	return &goodsv1.UpdateResponse{}, nil
}

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
}

func (s *serverAPI) Delete(ctx context.Context, req *goodsv1.DeleteRequest) (*goodsv1.DeleteResponse, error) {

	if err := ValidateDelete(req); err != nil {
		return nil, err
	}

	err := s.goods.RemoveGood(ctx, req.GetGoodsId())

	if err != nil {
		return nil, err
	}

	return &goodsv1.DeleteResponse{}, nil

}

func ValidateDelete(req *goodsv1.DeleteRequest) error {

	if req.GetGoodsId() == 0 {
		return status.Error(codes.InvalidArgument, "missing goods_id")
	}

	return nil

}
