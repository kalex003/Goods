package postgres

import (
	"Goods/internal/domain/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type GoodsDb struct {
	Db *sqlx.DB //пришлось сделать импортируемым ради крончика
}

func New(connstring string) (*GoodsDb, error) {

	db := sqlx.MustConnect("postgres", connstring)

	return &GoodsDb{Db: db}, nil
}

func (Gd *GoodsDb) SaveGoods(ctx context.Context, insmodel models.GoodsInfo) (models.GoodsInsertAnswers, error) {
	var result []models.GoodsInsertAnswer

	// Преобразуем структуру insmodel в JSON
	jsonData, marshalerr := json.Marshal(insmodel)
	if marshalerr != nil {
		return models.GoodsInsertAnswers{}, fmt.Errorf("SaveGoods: unable to marshal insmodel to JSON: %w", marshalerr)
	}

	err := Gd.Db.SelectContext(ctx, &result, `
	WITH cte AS (
		UPDATE goods.goods AS g
		SET place_id = c.place_id,
		    sku_id =  c.sku_id,
		    wbsticker_id = c.wbsticker_id,
		    barcode = c.barcode,
		    state_id = c.state_id,
		    ch_employee_id = c.ch_employee_id,
		    office_id = c.office_id,
		    wh_id = c.wh_id,
		    tare_id = c.tare_id,
		    tare_type = c.tare_type,
			ch_dt = CURRENT_TIMESTAMP
		FROM JSONB_TO_RECORDSET($1::JSONB -> 'data') AS c(goods_id BIGINT,
		    											  place_id BIGINT,
		                                 				  sku_id BIGINT,
		                                 				  wbsticker_id BIGINT,
		                                 				  barcode VARCHAR(30),
		                                 				  state_id CHAR(3),
		                                 				  ch_employee_id INTEGER,
		                                 				  office_id INTEGER,
		                                 				  wh_id INTEGER,
		                                 				  tare_id BIGINT,
		                                 				  tare_type CHAR(3))
		WHERE g.goods_id = c.goods_id
		RETURNING g.*)
	INSERT INTO goods.goodslog AS g 
	(
	    goods_id, 
	    place_id,
	    sku_id, 
	    wbsticker_id, 
	    barcode, 
	    state_id, 
	    ch_employee_id, 
	    office_id,
	    wh_id, 
	    tare_id, 
	    tare_type, 
	    ch_dt, 
	    is_del
	    )
	SELECT c.goods_id,
	       c.place_id,
	       c.sku_id,
	       c.wbsticker_id,
	       c.barcode,
	       c.state_id,
	       c.ch_employee_id,
	       c.office_id,
	       c.wh_id,
	       c.tare_id,
	       c.tare_type,
	       c.ch_dt,
	       c.is_del
	FROM cte c
	RETURNING g.goods_id,
	          g.ch_dt;`, jsonData)

	if err != nil {
		return models.GoodsInsertAnswers{}, fmt.Errorf("SaveGood: %w", err)
	}

	return models.GoodsInsertAnswers{
		GoodsInsertAnswers: result,
	}, nil
}

func (Gd *GoodsDb) UpdateGoods(ctx context.Context, updmodel models.GoodsUpdateInputs) (models.GoodsUpdateAnswers, error) {
	var result []models.GoodsUpdateAnswer

	// Преобразуем структуру insmodel в JSON
	jsonData, marshalerr := json.Marshal(updmodel)
	if marshalerr != nil {
		return models.GoodsUpdateAnswers{}, fmt.Errorf("UpdateGoods: unable to marshal updmodel to JSON: %w", marshalerr)
	}

	err := Gd.Db.SelectContext(ctx, &result, `
	WITH cte AS (
		UPDATE goods.goods AS g
		SET place_id = c.place_id,
		    sku_id =  c.sku_id,
		    wbsticker_id = c.wbsticker_id,
		    barcode = c.barcode,
		    state_id = c.state_id,
		    ch_employee_id = c.ch_employee_id,
		    office_id = c.office_id,
		    wh_id = c.wh_id,
		    tare_id = c.tare_id,
		    tare_type = c.tare_type,
			ch_dt = CURRENT_TIMESTAMP
		FROM JSONB_TO_RECORDSET($1::JSONB -> 'data') AS c(goods_id BIGINT,
		    											  place_id BIGINT,
		                                 				  sku_id BIGINT,
		                                 				  wbsticker_id BIGINT,
		                                 				  barcode VARCHAR(30),
		                                 				  state_id CHAR(3),
		                                 				  ch_employee_id INTEGER,
		                                 				  office_id INTEGER,
		                                 				  wh_id INTEGER,
		                                 				  tare_id BIGINT,
		                                 				  tare_type CHAR(3))
		WHERE g.goods_id = c.goods_id
		RETURNING g.*)
	INSERT INTO goods.goodslog AS g
	(
	    goods_id,
	    place_id,
	    sku_id,
	    wbsticker_id,
	    barcode,
	    state_id,
	    ch_employee_id,
	    office_id,
	    wh_id,
	    tare_id,
	    tare_type,
	    ch_dt,
	    is_del
	    )
	SELECT c.goods_id,
	       c.place_id,
	       c.sku_id,
	       c.wbsticker_id,
	       c.barcode,
	       c.state_id,
	       c.ch_employee_id,
	       c.office_id,
	       c.wh_id,
	       c.tare_id,
	       c.tare_type,
	       c.ch_dt,
	       c.is_del
	FROM cte c
	RETURNING g.goods_id,
	          g.ch_dt;`, jsonData)

	if err != nil {
		return models.GoodsUpdateAnswers{}, fmt.Errorf("UpdateGood: %w", err)
	}

	return models.GoodsUpdateAnswers{
		GoodsUpdateAnswers: result,
	}, nil

}

func (Gd *GoodsDb) SelectGoodsByIds(ctx context.Context, goodsIds []int64) (models.GoodsFullInfo, error) {

	var result []models.GoodFullInfo
	err := Gd.Db.SelectContext(ctx, &result, `SELECT * FROM goods.goods g WHERE g.goods_id = ANY($1)`, pq.Array(goodsIds))
	if err != nil {
		return models.GoodsFullInfo{}, fmt.Errorf("GetGood: %w", err)
	}
	return models.GoodsFullInfo{
		GoodsFullInfo: result,
	}, nil
}

func (Gd *GoodsDb) SelectGoodsByPlace(ctx context.Context, placeId int64) (models.GoodsFullInfo, error) {

	var result []models.GoodFullInfo
	err := Gd.Db.SelectContext(ctx, &result, `SELECT * FROM goods.goods g WHERE g.place_id = $1`, placeId)
	if err != nil {
		return models.GoodsFullInfo{}, fmt.Errorf("GetGoodbyplace: %w", err)
	}
	return models.GoodsFullInfo{
		GoodsFullInfo: result,
	}, nil
}

func (Gd *GoodsDb) SelectGoodsByTare(ctx context.Context, tareId int64) (models.GoodsFullInfo, error) {

	var result []models.GoodFullInfo
	err := Gd.Db.SelectContext(ctx, &result, `SELECT * FROM goods.goods g WHERE g.tare_id = $1`, tareId)
	if err != nil {
		return models.GoodsFullInfo{}, fmt.Errorf("GetGoodbytare: %w", err)
	}
	return models.GoodsFullInfo{
		GoodsFullInfo: result,
	}, nil
}

func (Gd *GoodsDb) SelectGoodsHistory(ctx context.Context, goodsId int64) (models.GoodsFullInfo, error) {

	var result []models.GoodFullInfo
	err := Gd.Db.SelectContext(ctx, &result, `SELECT * FROM goods.goodslog g WHERE g.goodsId = $1`, goodsId)
	if err != nil {
		return models.GoodsFullInfo{}, fmt.Errorf("GetGoodhistory: %w", err)
	}
	return models.GoodsFullInfo{
		GoodsFullInfo: result,
	}, nil
}

func (Gd *GoodsDb) UpdateIsDelOfGoods(ctx context.Context, updIsDelmodel models.GoodsUpdateIsDelInputs) (models.GoodsUpdateIsDelAnswers, error) {

	var result []models.GoodsUpdateIsDelAnswer

	// Преобразуем структуру insmodel в JSON
	jsonData, marshalerr := json.Marshal(updIsDelmodel)
	if marshalerr != nil {
		return models.GoodsUpdateIsDelAnswers{}, fmt.Errorf("SaveGoods: unable to marshal insmodel to JSON: %w", marshalerr)
	}

	err := Gd.Db.SelectContext(ctx, &result, `
	WITH cte AS (
		UPDATE goods.goods AS g
		SET is_del = c.is_del,
		ch_dt = CURRENT_TIMESTAMP
		FROM JSONB_TO_RECORDSET($1::JSONB -> 'data') AS c(goods_id BIGINT,
		                                 				   is_del BOOLEAN)
		WHERE g.goods_id = c.goods_id
		RETURNING g.*)
	INSERT INTO goods.goodslog AS g
	(
	    goods_id,
	    place_id,
	    sku_id,
	    wbsticker_id,
	    barcode,
	    state_id,
	    ch_employee_id,
	    office_id,
	    wh_id,
	    tare_id,
	    tare_type,
	    ch_dt,
	    is_del
	    )
	SELECT c.goods_id,
	       c.place_id,
	       c.sku_id,
	       c.wbsticker_id,
	       c.barcode,
	       c.state_id,
	       c.ch_employee_id,
	       c.office_id,
	       c.wh_id,
	       c.tare_id,
	       c.tare_type,
	       c.ch_dt,
	       c.is_del
	FROM cte c
	RETURNING g.goods_id,
	          g.ch_dt;`, jsonData) //Важен ли порядок в RETURNING?

	if err != nil {
		return models.GoodsUpdateIsDelAnswers{}, fmt.Errorf("UpdateIsDelOfGoods: %w", err)
	}

	return models.GoodsUpdateIsDelAnswers{
		GoodsUpdateIsDelAnswers: result,
	}, nil
}
