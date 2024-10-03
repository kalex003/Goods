package goods

import "log/slog"

import (
	"Goods/internal/domain/models"
	"context"
)

type Goods struct {
	log           *slog.Logger
	GoodsSaver    GoodsSaver
	GoodsSelecter GoodsSelecter
	GoodsUpdater  GoodsUpdater
	GoodsDeleter  GoodsDeleter
}

// а мб не надо так много интерфейсов?
type GoodsSaver interface { //для типа сторэдж
	SaveGoods(ctx context.Context, insmodel models.GoodsInfo) (models.GoodsInsertAnswers, error)
}

type GoodsUpdater interface { //для типа сторэдж
	UpdateGoods(ctx context.Context, updmodel models.GoodsUpdateInputs) (models.GoodsUpdateAnswers, error)
}

type GoodsSelecter interface { //для типа сторэдж
	SelectGoodsByIds(ctx context.Context, goodsIds []int64) (models.GoodsFullInfo, error) // а почему нельзя вот так передать? goodsId models.GoodFullInfo.GoodsId
	SelectGoodsByPlace(ctx context.Context, placeId int64) (models.GoodsFullInfo, error)
	SelectGoodsByTare(ctx context.Context, placeId int64) (models.GoodsFullInfo, error)
	SelectGoodsHistory(context.Context, int64) (models.GoodsFullInfo, error)
}

type GoodsDeleter interface {
	UpdateIsDelOfGoods(ctx context.Context, updIsDelmodel models.GoodsUpdateIsDelInputs) (models.GoodsUpdateIsDelAnswers, error)
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

func (g *Goods) InsertNewGoods(ctx context.Context, insmodel models.GoodsInfo) (models.GoodsInsertAnswers, error) { // Правда скорее всего это плохо, что я в 2 слоях использую одну и ту же модель
	const op = "Goods.Save"

	log := g.log.With(
		slog.String("op", op),
		slog.Any("insmodel", insmodel), // тут подумать надо
	)

	log.Info("attempting to insert goods")

	ans, err := g.GoodsSaver.SaveGoods(ctx, insmodel) //вызываю пг-шку

	if err != nil {
		return models.GoodsInsertAnswers{}, err
	}

	return ans, err

}

func (g *Goods) ChangeGoods(ctx context.Context, updmodel models.GoodsUpdateInputs) (models.GoodsUpdateAnswers, error) {

	const op = "Goods.Update"

	log := g.log.With(
		slog.String("op", op),
		slog.Any("updmodel", updmodel), // тут подумать надо
	)

	log.Info("Attempting to update goods")

	ans, err := g.GoodsUpdater.UpdateGoods(ctx, updmodel) //вызываю пг-шку

	if err != nil {
		return models.GoodsUpdateAnswers{}, err
	}

	return ans, nil

}

func (g *Goods) GetByIdsGoodsInfo(ctx context.Context, goodsIDs []int64) (models.GoodsFullInfo, error) {

	const op = "Goods.Get"

	log := g.log.With(
		slog.String("op", op),
		slog.Any("goodsIDs", goodsIDs), // тут подумать надо
	)

	log.Info("attempting to GetByIds")

	goodsInfo, err := g.GoodsSelecter.SelectGoodsByIds(ctx, goodsIDs) //вызываю пг-шку

	if err != nil {
		return models.GoodsFullInfo{}, err
	}

	return goodsInfo, nil

}

func (g *Goods) GetByPlaceGoodsInfo(ctx context.Context, placeID int64) (models.GoodsFullInfo, error) {

	const op = "Goods.Getbyplace"

	log := g.log.With(
		slog.String("op", op),
		slog.Int("placeId", int(placeID)), // тут подумать надо
	)

	log.Info("attempting to Getbyplace")

	goodsInfo, err := g.GoodsSelecter.SelectGoodsByPlace(ctx, placeID) //вызываю пг-шку

	if err != nil {
		return models.GoodsFullInfo{}, err
	}

	return goodsInfo, nil

}

func (g *Goods) GetByTareGoodsInfo(ctx context.Context, tareID int64) (models.GoodsFullInfo, error) {

	const op = "Goods.Getbytare"

	log := g.log.With(
		slog.String("op", op),
		slog.Int("tareId", int(tareID)), // тут подумать надо
	)

	log.Info("attempting to Getbytare")

	goodsInfo, err := g.GoodsSelecter.SelectGoodsByTare(ctx, tareID) //вызываю пг-шку

	if err != nil {
		return models.GoodsFullInfo{}, err
	}

	return goodsInfo, nil

}

func (g *Goods) GetGoodsHistory(ctx context.Context, goodsID int64) (models.GoodsFullInfo, error) {

	const op = "Goods.GetHistory"

	log := g.log.With(
		slog.String("op", op),
		slog.Int("tareId", int(goodsID)), // тут подумать надо
	)

	log.Info("attempting to GetHistory")

	goodsInfo, err := g.GoodsSelecter.SelectGoodsHistory(ctx, goodsID) //вызываю пг-шку

	if err != nil {
		return models.GoodsFullInfo{}, err
	}

	return goodsInfo, nil

}

func (g *Goods) ChangeIsDelOfGoods(ctx context.Context, updIsDelmodel models.GoodsUpdateIsDelInputs) (models.GoodsUpdateIsDelAnswers, error) {

	const op = "Good.ChangeIsDel"

	log := g.log.With(
		slog.String("op", op),
		slog.Any("updmodel", updIsDelmodel), // тут подумать надо
	)

	log.Info("attempting to login user")

	ans, err := g.GoodsDeleter.UpdateIsDelOfGoods(ctx, updIsDelmodel) //вызываю пг-шку

	if err != nil {
		return models.GoodsUpdateIsDelAnswers{}, err
	}

	return ans, nil

}
