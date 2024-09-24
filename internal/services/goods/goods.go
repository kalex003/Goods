package goods

import (
	"Goods/internal/domain/models"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

type Goods struct {
	log           *slog.Logger
	GoodsSaver    GoodsSaver
	GoodsSelecter GoodsSelecter
	GoodsUpdater  GoodsUpdater
	GoodsDeleter  GoodsDeleter
}

type GoodsSaver interface { //для типа сторэдж
	SaveGood(ctx context.Context, insmodel models.GoodInfo) (int64, time.Time, error)
}

type GoodsUpdater interface { //для типа сторэдж
	UpdateGood(ctx context.Context, updmodel models.GoodFullInfo) error
}

type GoodsSelecter interface { //для типа сторэдж
	SelectGood(ctx context.Context, goodsId int64) (models.GoodFullInfo, error) // а почему нельзя вот так передать? goodsId models.GoodFullInfo.GoodsId
}

type GoodsDeleter interface {
	DeleteGood(ctx context.Context, goods_id int64) error
}

func New(log *slog.Logger, goodsSaver GoodsSaver, goodsUpdater GoodsUpdater, goodsSelecter GoodsSelecter, goodsDeleter GoodsDeleter) *Goods {
	return &Goods{
		log:           log,
		GoodsSaver:    goodsSaver,
		GoodsUpdater:  goodsUpdater,
		GoodsSelecter: goodsSelecter,
		GoodsDeleter:  goodsDeleter,
	}
}

func (g *Goods) InsertNewGood(ctx context.Context, insmodel models.GoodInfo) (int64, *timestamppb.Timestamp, error) {
	const op = "Good.Save"

	log := g.log.With(
		slog.String("op", op),
		slog.Any("insmodel", insmodel), // тут подумать надо
	)

	log.Info("attempting to login user")

	goodId, chDt, err := g.GoodsSaver.SaveGood(ctx, insmodel) //вызываю пг-шку

	if err != nil {
		return -1, &timestamppb.Timestamp{}, err
	}

	return goodId, timestamppb.New(chDt), err

}

func (g *Goods) ChangeGood(ctx context.Context, updmodel models.GoodFullInfo) error {

	const op = "Good.Update"

	log := g.log.With(
		slog.String("op", op),
		slog.Any("updmodel", updmodel), // тут подумать надо
	)

	log.Info("attempting to login user")

	err := g.GoodsUpdater.UpdateGood(ctx, updmodel) //вызываю пг-шку

	if err != nil {
		return err
	}

	return nil

}

func (g *Goods) GetGoodInfo(ctx context.Context, goodsID int64) (models.GoodFullInfo, error) {

	const op = "Good.Get"

	log := g.log.With(
		slog.String("op", op),
		slog.Int("updmodel", int(goodsID)), // тут подумать надо
	)

	log.Info("attempting to login user")

	goodInfo, err := g.GoodsSelecter.SelectGood(ctx, goodsID) //вызываю пг-шку

	if err != nil {
		return models.GoodFullInfo{}, err
	}

	return goodInfo, nil

}

func (g *Goods) RemoveGood(ctx context.Context, goodsID int64) error {

	const op = "Good.Get"

	log := g.log.With(
		slog.String("op", op),
		slog.Int("updmodel", int(goodsID)), // тут подумать надо
	)

	log.Info("attempting to login user")

	err := g.GoodsDeleter.DeleteGood(ctx, goodsID) //вызываю пг-шку

	if err != nil {
		return err
	}

	return nil

}
