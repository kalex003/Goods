package postgres

import (
	"Goods/internal/domain/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GoodsDb struct {
	Db *pgxpool.Pool // используем pgxpool для пула подключений
}

func New(connstring string) (*GoodsDb, error) {

	pool, err := pgxpool.New(context.Background(), connstring)
	if err != nil {
		return nil, err
	}

	return &GoodsDb{Db: pool}, nil
}

func (Gd *GoodsDb) SaveGoods(ctx context.Context, insmodel models.GoodsInfo) (models.GoodsInsertAnswers, error) {
	var result []models.GoodsInsertAnswer

	/*	// Преобразуем структуру insmodel в JSON
		jsonData, marshalerr := json.Marshal(insmodel)
		if marshalerr != nil {
			return models.GoodsInsertAnswers{}, fmt.Errorf("SaveGoods: unable to marshal insmodel to JSON: %w", marshalerr)
		}*/

	/*	rows, err := Gd.Db.Query(ctx, `
		WITH cte AS (
		INSERT INTO goods.goods AS g
		(
		    place_id,
		    sku_id,
		    wbsticker_id,
		    barcode,
		    state_id,
		    ch_employee_id,
		    office_id,
		    wh_id,
		    tare_id,
		    tare_type
		)
		SELECT c.place_id,
		       c.sku_id,
		       c.wbsticker_id,
		       c.barcode,
		       c.state_id,
		       c.ch_employee_id,
		       c.office_id,
		       c.wh_id,
		       c.tare_id,
		       c.tare_type
		FROM JSONB_TO_RECORDSET($1::JSONB -> 'data') AS c(place_id BIGINT,
		                                 				  sku_id BIGINT,
		                                 				  wbsticker_id BIGINT,
		                                 				  barcode VARCHAR(30),
		                                 				  state_id CHAR(3),
		                                 				  ch_employee_id INTEGER,
		                                 				  office_id INTEGER,
		                                 				  wh_id INTEGER,
		                                 				  tare_id BIGINT,
		                                 				  tare_type CHAR(3))
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
		          g.ch_dt;`, jsonData)*/

	cte_statement := squirrel.Insert("goods.goodslog AS g").
		Columns("place_id", "sku_id", "wbsticker_id", "barcode", "state_id", "ch_employee_id", "office_id", "wh_id", "tare_id", "tare_type").
		Prefix("WITH ins_cte AS (").
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING g.*)")

	for _, i := range insmodel.GoodsInfo {
		cte_statement = cte_statement.Values(i.PlaceId, i.SkuId, i.WbstickerId, i.Barcode, i.StateId, i.ChEmployeeId, i.OfficeId, i.WhId, i.TareId, i.TareType)
	}

	_, args, err := cte_statement.ToSql()
	if err != nil {
		return models.GoodsInsertAnswers{}, fmt.Errorf("SaveGoods: unable to build query: %w", err)
	}
	/*	sql, _, err := squirrel.StatementBuilder.
		Insert("").
		PrefixExpr(
			squirrel.Insert("").
				Select(
					squirrel.Select("c.place_id", "c.sku_id", "c.wbsticker_id", "c.barcode", "c.state_id", "c.ch_employee_id", "c.office_id", "c.wh_id", "c.tare_id", "c.tare_type").
						From("JSONB_TO_RECORDSET(?::JSONB -> 'data') AS c	(place_id BIGINT, sku_id BIGINT, wbsticker_id BIGINT, barcode VARCHAR(30), state_id CHAR(3), ch_employee_id INTEGER, office_id INTEGER, wh_id INTEGER, tare_id BIGINT, tare_type CHAR(3))"),
				).
				Into("goods.goods AS g").
				Columns("place_id", "sku_id", "wbsticker_id", "barcode", "state_id", "ch_employee_id", "office_id", "wh_id", "tare_id", "tare_type").
				Prefix("WITH ins_cte AS (").
				PlaceholderFormat(squirrel.Dollar).
				Suffix("RETURNING g.*)"),
		).
		Into("goods.goodslog AS g").
		Columns("goods_id", "place_id", "sku_id", "wbsticker_id", "barcode", "state_id", "ch_employee_id", "office_id", "wh_id", "tare_id", "tare_type", "ch_dt", "is_del").
		Select(
			squirrel.Select("c.goods_id", "c.place_id", "c.sku_id", "c.wbsticker_id", "c.barcode", "c.state_id", "c.ch_employee_id", "c.office_id", "c.wh_id", "c.tare_id", "c.tare_type", "c.ch_dt", "c.is_del").
				From("ins_cte c"),
		).
		Suffix("RETURNING g.goods_id, g.ch_dt").
		ToSql()*/

	sql, args, err := squirrel.StatementBuilder.
		Insert("goods.goodslog AS g").
		Columns("goods_id", "place_id", "sku_id", "wbsticker_id", "barcode", "state_id", "ch_employee_id", "office_id", "wh_id", "tare_id", "tare_type", "ch_dt", "is_del").
		Select(
			squirrel.Select("c.goods_id", "c.place_id", "c.sku_id", "c.wbsticker_id", "c.barcode", "c.state_id", "c.ch_employee_id", "c.office_id", "c.wh_id", "c.tare_id", "c.tare_type", "c.ch_dt", "c.is_del").
				From("ins_cte c"),
		).
		PrefixExpr(cte_statement).
		Suffix("RETURNING g.goods_id, g.ch_dt").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return models.GoodsInsertAnswers{}, fmt.Errorf("SaveGoods: unable to build query: %w", err)
	}

	// Выполняем запрос
	rows, err := Gd.Db.Query(ctx, sql, args...)
	if err != nil {
		return models.GoodsInsertAnswers{}, fmt.Errorf("SaveGoods: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var answer models.GoodsInsertAnswer
		if err := rows.Scan(&answer.GoodsId, &answer.ChDt); err != nil {
			return models.GoodsInsertAnswers{}, err
		}
		result = append(result, answer)
	}

	if rows.Err() != nil {
		return models.GoodsInsertAnswers{}, rows.Err()
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
	/*
		rows, err := Gd.Db.Query(ctx, `
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
		          g.ch_dt;`, jsonData)*/

	sql, _, err := squirrel.StatementBuilder.
		Insert("").
		PrefixExpr(
			squirrel.Update("").
				Set("ch_dt", squirrel.Expr("CURRENT_TIMESTAMP AT TIME ZONE 'Europe/Moscow'")).
				Set("place_id", squirrel.Expr("c.place_id")).
				Set("sku_id", squirrel.Expr("c.sku_id")).
				Set("wbsticker_id", squirrel.Expr("c.wbsticker_id")).
				Set("barcode", squirrel.Expr("c.barcode")).
				Set("state_id", squirrel.Expr("c.state_id")).
				Set("ch_employee_id", squirrel.Expr("c.ch_employee_id")).
				Set("office_id", squirrel.Expr("c.office_id")).
				Set("wh_id", squirrel.Expr("c.wh_id")).
				Set("tare_id", squirrel.Expr("c.tare_id")).
				Set("tare_type", squirrel.Expr("c.tare_type")).
				From("JSONB_TO_RECORDSET(?::JSONB -> 'data') AS c (goods_id BIGINT, place_id BIGINT, sku_id BIGINT, wbsticker_id BIGINT, barcode VARCHAR(30), state_id CHAR(3), ch_employee_id INTEGER, office_id INTEGER, wh_id INTEGER, tare_id BIGINT, tare_type CHAR(3))").
				Table("goods.goods AS g").
				Where("g.goods_id = c.goods_id").
				Prefix("WITH upd_cte AS (").
				PlaceholderFormat(squirrel.Dollar).
				Suffix("RETURNING g.*)"),
		).
		Into("goods.goodslog AS g").
		Columns("goods_id", "place_id", "sku_id", "wbsticker_id", "barcode", "state_id", "ch_employee_id", "office_id", "wh_id", "tare_id", "tare_type", "ch_dt", "is_del").
		Select(
			squirrel.Select("c.goods_id", "c.place_id", "c.sku_id", "c.wbsticker_id", "c.barcode", "c.state_id", "c.ch_employee_id", "c.office_id", "c.wh_id", "c.tare_id", "c.tare_type", "c.ch_dt", "c.is_del").
				From("upd_cte c"),
		).
		Suffix("RETURNING g.goods_id, g.ch_dt").
		ToSql()

	if err != nil {
		return models.GoodsUpdateAnswers{}, fmt.Errorf("UpdateGoods: unable to build query: %w", err)
	}

	// Выполняем запрос
	rows, err := Gd.Db.Query(ctx, sql, jsonData)
	if err != nil {
		return models.GoodsUpdateAnswers{}, fmt.Errorf("UpdateGoods: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var answer models.GoodsUpdateAnswer
		if err := rows.Scan(&answer.GoodsId, &answer.ChDt); err != nil {
			return models.GoodsUpdateAnswers{}, err
		}
		result = append(result, answer)
	}

	if rows.Err() != nil {
		return models.GoodsUpdateAnswers{}, rows.Err()
	}

	return models.GoodsUpdateAnswers{
		GoodsUpdateAnswers: result,
	}, nil

}

func (Gd *GoodsDb) SelectGoodsByIds(ctx context.Context, goodsIds *[]int64) (models.GoodsFullInfo, error) {

	var result []models.GoodFullInfo
	var err error
	var sql string
	var rows pgx.Rows
	if len(*goodsIds) == 0 {
		sql, _, err = squirrel.StatementBuilder.
			Select("g.goods_id", "g.place_id", "g.sku_id", "g.wbsticker_id", "g.barcode", "g.state_id", "g.ch_employee_id", "g.office_id", "g.wh_id", "g.tare_id", "g.tare_type", "g.ch_dt", "g.is_del").
			From("goods.goods AS g").
			ToSql()

		if err != nil {
			return models.GoodsFullInfo{}, fmt.Errorf("SelectGoodsByIds: unable to build query: %w", err)
		}

		rows, err = Gd.Db.Query(ctx, sql)

	} else {
		sql, _, err = squirrel.StatementBuilder.
			Select("g.goods_id", "g.place_id", "g.sku_id", "g.wbsticker_id", "g.barcode", "g.state_id", "g.ch_employee_id", "g.office_id", "g.wh_id", "g.tare_id", "g.tare_type", "g.ch_dt", "g.is_del").
			From("goods.goods AS g").
			Where("g.goods_id = ANY(?)").
			PlaceholderFormat(squirrel.Dollar).
			ToSql()

		if err != nil {
			return models.GoodsFullInfo{}, fmt.Errorf("SelectGoodsByIds: unable to build query: %w", err)
		}

		rows, err = Gd.Db.Query(ctx, sql, goodsIds)
	}

	if err != nil {
		return models.GoodsFullInfo{}, fmt.Errorf("SelectGoodsByIds: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var answer models.GoodFullInfo
		if err := rows.Scan(&answer.GoodsId, &answer.PlaceId, &answer.SkuId, &answer.WbstickerId, &answer.Barcode, &answer.StateId, &answer.ChEmployeeId, &answer.OfficeId, &answer.WhId, &answer.TareId, &answer.TareType, &answer.ChDt, &answer.IsDel); err != nil {
			return models.GoodsFullInfo{}, err
		}
		result = append(result, answer)
	}
	return models.GoodsFullInfo{
		GoodsFullInfo: result,
	}, nil

}

func (Gd *GoodsDb) SelectGoodsByPlace(ctx context.Context, placeId int64) (models.GoodsFullInfo, error) {

	var result []models.GoodFullInfo

	sql, args, err := squirrel.StatementBuilder.
		Select("g.goods_id", "g.place_id", "g.sku_id", "g.wbsticker_id", "g.barcode", "g.state_id", "g.ch_employee_id", "g.office_id", "g.wh_id", "g.tare_id", "g.tare_type", "g.ch_dt", "g.is_del").
		From("goods.goods AS g").
		Where("g.place_id = ?", placeId).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return models.GoodsFullInfo{}, fmt.Errorf("SelectGoodsByPlace: unable to build query: %w", err)
	}

	rows, err := Gd.Db.Query(ctx, sql, args...)
	if err != nil {
		return models.GoodsFullInfo{}, fmt.Errorf("SelectGoodsByPlace: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var answer models.GoodFullInfo
		if err := rows.Scan(&answer.GoodsId, &answer.PlaceId, &answer.SkuId, &answer.WbstickerId, &answer.Barcode, &answer.StateId, &answer.ChEmployeeId, &answer.OfficeId, &answer.WhId, &answer.TareId, &answer.TareType, &answer.ChDt, &answer.IsDel); err != nil {
			return models.GoodsFullInfo{}, err
		}
		result = append(result, answer)
	}
	return models.GoodsFullInfo{
		GoodsFullInfo: result,
	}, nil

}

func (Gd *GoodsDb) SelectGoodsByTare(ctx context.Context, tareId int64) (models.GoodsFullInfo, error) {

	var result []models.GoodFullInfo

	sql, args, err := squirrel.StatementBuilder.
		Select("g.goods_id", "g.place_id", "g.sku_id", "g.wbsticker_id", "g.barcode", "g.state_id", "g.ch_employee_id", "g.office_id", "g.wh_id", "g.tare_id", "g.tare_type", "g.ch_dt", "g.is_del").
		From("goods.goods AS g").
		Where("g.tare_id = ? ", tareId).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return models.GoodsFullInfo{}, fmt.Errorf("SelectGoodsByTare: unable to build query: %w", err)
	}

	rows, err := Gd.Db.Query(ctx, sql, args...)
	if err != nil {
		return models.GoodsFullInfo{}, fmt.Errorf("SelectGoodsByTare: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var answer models.GoodFullInfo
		if err := rows.Scan(&answer.GoodsId, &answer.PlaceId, &answer.SkuId, &answer.WbstickerId, &answer.Barcode, &answer.StateId, &answer.ChEmployeeId, &answer.OfficeId, &answer.WhId, &answer.TareId, &answer.TareType, &answer.ChDt, &answer.IsDel); err != nil {
			return models.GoodsFullInfo{}, err
		}
		result = append(result, answer)
	}
	return models.GoodsFullInfo{
		GoodsFullInfo: result,
	}, nil

}

func (Gd *GoodsDb) SelectGoodsHistory(ctx context.Context, goodsId int64) (models.GoodsFullInfo, error) {

	var result []models.GoodFullInfo

	sql, args, err := squirrel.StatementBuilder.
		Select("g.goods_id", "g.place_id", "g.sku_id", "g.wbsticker_id", "g.barcode", "g.state_id", "g.ch_employee_id", "g.office_id", "g.wh_id", "g.tare_id", "g.tare_type", "g.ch_dt", "g.is_del").
		From("goods.goodslog AS g").
		Where("g.goods_id = ? ", goodsId).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return models.GoodsFullInfo{}, fmt.Errorf("SelectGoodsHistory: unable to build query: %w", err)
	}

	rows, err := Gd.Db.Query(ctx, sql, args...)
	if err != nil {
		return models.GoodsFullInfo{}, fmt.Errorf("SelectGoodsHistory: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var answer models.GoodFullInfo
		if err := rows.Scan(&answer.GoodsId, &answer.PlaceId, &answer.SkuId, &answer.WbstickerId, &answer.Barcode, &answer.StateId, &answer.ChEmployeeId, &answer.OfficeId, &answer.WhId, &answer.TareId, &answer.TareType, &answer.ChDt, &answer.IsDel); err != nil {
			return models.GoodsFullInfo{}, err
		}
		result = append(result, answer)
	}
	return models.GoodsFullInfo{
		GoodsFullInfo: result,
	}, nil
}

func (Gd *GoodsDb) UpdateIsDelOfGoods(ctx context.Context, updIsDelModel models.GoodsUpdateIsDelInputs) (models.GoodsUpdateIsDelAnswers, error) {

	var result []models.GoodsUpdateIsDelAnswer

	// Преобразуем структуру insmodel в JSON
	jsonData, marshalErr := json.Marshal(updIsDelModel)

	if marshalErr != nil {
		return models.GoodsUpdateIsDelAnswers{}, fmt.Errorf("UpdateIsDelOfGoods: unable to marshal model to JSON: %w", marshalErr)
	}

	/*	rows, err := Gd.Db.Query(ctx, `
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
		          g.ch_dt;`, jsonData)*/

	sql, _, err := squirrel.StatementBuilder.
		Insert("").
		PrefixExpr(
			squirrel.Update("").
				Set("is_del", squirrel.Expr("c.is_del")).
				Set("ch_dt", squirrel.Expr("CURRENT_TIMESTAMP AT TIME ZONE 'Europe/Moscow'")).
				From("JSONB_TO_RECORDSET(?::JSONB -> 'data') AS c(goods_id BIGINT, is_del BOOLEAN)").
				Table("goods.goods AS g").
				Where("g.goods_id = c.goods_id").
				Prefix("WITH upd_cte AS (").
				PlaceholderFormat(squirrel.Dollar).
				Suffix("RETURNING g.*)"),
		).
		Into("goods.goodslog AS g").
		Columns("goods_id", "place_id", "sku_id", "wbsticker_id", "barcode", "state_id", "ch_employee_id", "office_id", "wh_id", "tare_id", "tare_type", "ch_dt", "is_del").
		Select(
			squirrel.Select("c.goods_id", "c.place_id", "c.sku_id", "c.wbsticker_id", "c.barcode", "c.state_id", "c.ch_employee_id", "c.office_id", "c.wh_id", "c.tare_id", "c.tare_type", "c.ch_dt", "c.is_del").
				From("upd_cte c"),
		).
		Suffix("RETURNING g.goods_id, g.ch_dt, g.is_del").
		ToSql()

	if err != nil {
		return models.GoodsUpdateIsDelAnswers{}, fmt.Errorf("UpdateIsDelOfGoods: unable to build query: %w", err)
	}

	// Выполняем запрос
	rows, err := Gd.Db.Query(ctx, sql, jsonData)
	if err != nil {
		return models.GoodsUpdateIsDelAnswers{}, fmt.Errorf("UpdateIsDelOfGoods: %w", err)
	}
	defer rows.Close()

	// Читаем результаты
	for rows.Next() {
		var answer models.GoodsUpdateIsDelAnswer
		if err := rows.Scan(&answer.GoodsId, &answer.ChDt, &answer.IsDel); err != nil {
			return models.GoodsUpdateIsDelAnswers{}, fmt.Errorf("UpdateIsDelOfGoods: unable to scan row: %w", err)
		}
		result = append(result, answer)
	}

	return models.GoodsUpdateIsDelAnswers{
		GoodsUpdateIsDelAnswers: result,
	}, nil
}
