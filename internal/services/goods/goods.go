package goods

import (
	domainmodels "Goods/internal/domain/models"
	dbmodels "Goods/internal/storage/models"
	"golang.org/x/net/context"
	"log/slog"
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
	SaveGoods(ctx context.Context, insmodel dbmodels.GoodsInsertInputs) (dbmodels.GoodsInsertAnswers, error)
}

type GoodsUpdater interface { //для типа сторэдж
	UpdateGoods(ctx context.Context, updmodel dbmodels.GoodsUpdateInputs) (dbmodels.GoodsUpdateAnswers, error)
}

type GoodsSelecter interface { //для типа сторэдж
	SelectGoodsByIds(ctx context.Context, goodsIds *[]int64) (dbmodels.GoodsGetAnswers, error)
	SelectGoodsByPlace(ctx context.Context, placeId int64) (dbmodels.GoodsGetAnswers, error)
	SelectGoodsByTare(ctx context.Context, placeId int64) (dbmodels.GoodsGetAnswers, error)
	SelectGoodsHistory(ctx context.Context, goodsId int64) (dbmodels.GoodsGetAnswers, error)
}

type GoodsDeleter interface {
	UpdateIsDelOfGoods(ctx context.Context, updIsDelmodel dbmodels.GoodsUpdateIsDelInputs) (dbmodels.GoodsUpdateIsDelAnswers, error)
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

func (g *Goods) InsertNewGoods(ctx context.Context, insmodel domainmodels.GoodsFullInfo) (domainmodels.GoodsFullInfo, error) { // Правда скорее всего это плохо, что я в 2 слоях использую одну и ту же модель
	const op = "Goods.Save"

	log := g.log.With(
		slog.String("op", op),
		slog.Any("insmodel", insmodel), // тут подумать надо
	)

	log.Info("attempting to insert goods")

	ans, err := g.GoodsSaver.SaveGoods(ctx, dbmodels.ConvertGoodsFullInfoToGoodsInsertInputs(insmodel)) //вызываю пг-шку

	if err != nil {
		return domainmodels.GoodsFullInfo{}, err
	}

	return dbmodels.ConvertGoodsinsertAnswersToGoodsFullInfo(ans), err

}

func (g *Goods) ChangeGoods(ctx context.Context, updmodel domainmodels.GoodsFullInfo) (domainmodels.GoodsFullInfo, error) {

	const op = "Goods.Update"

	log := g.log.With(
		slog.String("op", op),
		slog.Any("updmodel", updmodel), // тут подумать надо
	)

	log.Info("Attempting to update goods")

	ans, err := g.GoodsUpdater.UpdateGoods(ctx, dbmodels.ConvertGoodsFullInfoToGoodsUpdateInputs(updmodel)) //вызываю пг-шку

	if err != nil {
		return domainmodels.GoodsFullInfo{}, err
	}

	return dbmodels.ConvertGoodsUpdateAnswerToGoodsFullInfo(ans), nil

}

func (g *Goods) GetByIdsGoodsInfo(ctx context.Context, goodsIDs *[]int64) (domainmodels.GoodsFullInfo, error) {

	const op = "Goods.Get"

	log := g.log.With(
		slog.String("op", op),
		slog.Any("goodsIDs", goodsIDs), // тут подумать надо
	)

	log.Info("attempting to GetByIds")

	goodsInfo, err := g.GoodsSelecter.SelectGoodsByIds(ctx, goodsIDs) //вызываю пг-шку

	if err != nil {
		return domainmodels.GoodsFullInfo{}, err
	}

	return dbmodels.ConvertGoodsGetAnswerToGoodsFullInfo(goodsInfo), nil

}

func (g *Goods) GetByPlaceGoodsInfo(ctx context.Context, placeID int64) (domainmodels.GoodsFullInfo, error) {

	const op = "Goods.Getbyplace"

	log := g.log.With(
		slog.String("op", op),
		slog.Int("placeId", int(placeID)), // тут подумать надо
	)

	log.Info("attempting to Getbyplace")

	goodsInfo, err := g.GoodsSelecter.SelectGoodsByPlace(ctx, placeID) //вызываю пг-шку

	if err != nil {
		return domainmodels.GoodsFullInfo{}, err
	}

	return dbmodels.ConvertGoodsGetAnswerToGoodsFullInfo(goodsInfo), nil

}

func (g *Goods) GetByTareGoodsInfo(ctx context.Context, tareID int64) (domainmodels.GoodsFullInfo, error) {

	const op = "Goods.Getbytare"

	log := g.log.With(
		slog.String("op", op),
		slog.Int("tareId", int(tareID)), // тут подумать надо
	)

	log.Info("attempting to Getbytare")

	goodsInfo, err := g.GoodsSelecter.SelectGoodsByTare(ctx, tareID) //вызываю пг-шку

	if err != nil {
		return domainmodels.GoodsFullInfo{}, err
	}

	return dbmodels.ConvertGoodsGetAnswerToGoodsFullInfo(goodsInfo), nil

}

func (g *Goods) GetGoodsHistory(ctx context.Context, goodsID int64) (domainmodels.GoodsFullInfo, error) {

	const op = "Goods.GetHistory"

	log := g.log.With(
		slog.String("op", op),
		slog.Int("tareId", int(goodsID)), // тут подумать надо
	)

	log.Info("attempting to GetHistory")

	goodsInfo, err := g.GoodsSelecter.SelectGoodsHistory(ctx, goodsID) //вызываю пг-шку

	if err != nil {
		return domainmodels.GoodsFullInfo{}, err
	}

	return dbmodels.ConvertGoodsGetAnswerToGoodsFullInfo(goodsInfo), nil

}

func (g *Goods) ChangeIsDelOfGoods(ctx context.Context, updIsDelmodel domainmodels.GoodsFullInfo) (domainmodels.GoodsFullInfo, error) {

	const op = "Good.ChangeIsDel"

	log := g.log.With(
		slog.String("op", op),
		slog.Any("updmodel", updIsDelmodel), // тут подумать надо
	)

	log.Info("attempting to login user")

	ans, err := g.GoodsDeleter.UpdateIsDelOfGoods(ctx, dbmodels.ConvertGoodsFullInfoToGoodsUpdateIsDelInputs(updIsDelmodel)) //вызываю пг-шку

	if err != nil {
		return domainmodels.GoodsFullInfo{}, err
	}

	return dbmodels.ConvertGoodsUpdateIsDelAnswerToGoodsFullInfo(ans), nil

}
