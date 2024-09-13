package postgres

import (
	"Goods/internal/domain/models"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type GoodsDb struct {
	db *sqlx.DB
}

func New(connstring string) (*GoodsDb, error) {

	db := sqlx.MustConnect("postgres", connstring)

	return &GoodsDb{db: db}, nil
}

func (Gd *GoodsDb) SaveGood(ctx context.Context, insmodel models.GoodInfo) (int64, timestamp.Timestamp, error) {
	var result struct {
		GoodsId  int64
		CreateDt timestamp.Timestamp
	}

	err := Gd.db.Get(&result, `insert into goods.goods (place_id, sku_id, wbsticker_id, barcode, state_id, ch_employee_id, office_id, wh_id, tare_id, tare_type) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, returning goods_id, ch_dt`, insmodel.PlaceId, insmodel.SkuId, insmodel.WbstickerId, insmodel.Barcode, insmodel.StateId, insmodel.ChEmployeeId, insmodel.OfficeId, insmodel.WhId, insmodel.TareId, insmodel.TareType)
	if err != nil {
		var err_time timestamp.Timestamp //надо будет здесь какое-то время сгенерировать
		return 0, err_time, fmt.Errorf("SaveGood: %w", err)
	}
	return result.GoodsId, result.CreateDt, nil
}

func (Gd *GoodsDb) UpdateGood(ctx context.Context, updmodel models.GoodFullInfo) error {

	err := Gd.db.MustExec(`UPDATE goods.goods AS g
								 SET place_id = COALESCE($1, place_id), 
                       				 sku_id = COALESCE($2, sku_id), 
                       				 wbsticker_id = COALESCE($3, wbsticker_id),
                       				 barcode = COALESCE($4, barcode),
                       				 state_id = COALESCE($5, state_id), 
                       				 ch_employee_id = COALESCE($6, ch_employee_id),
                       				 office_id = COALESCE($7, office_id),
                       				 wh_id = COALESCE($8, wh_id), 
                       				 tare_id = COALESCE($9, tare_id),
                       				 tare_type = COALESCE($10, tare_type),
									 ch_dt = CURRENT_TIMESTAMP
                       				 WHERE g.goods_id = $11 `, updmodel.PlaceId, updmodel.SkuId, updmodel.WbstickerId, updmodel.Barcode, updmodel.StateId, updmodel.ChEmployeeId, updmodel.OfficeId, updmodel.WhId, updmodel.TareId, updmodel.TareType, updmodel.GoodsId)
	if err != nil {
		return fmt.Errorf("SaveGood: %w", err)
	}
	return nil
}

func (Gd *GoodsDb) SelectGood(ctx context.Context, goodsId int64) (models.GoodFullInfo, error) {

	var result models.GoodFullInfo
	err := Gd.db.Get(&result, `SELECT * FROM goods.goods g WHERE g.goods_id = $1`, goodsId)
	if err != nil {
		return models.GoodFullInfo{}, fmt.Errorf("SaveGood: %w", err)
	}
	return result, nil
}

func (Gd *GoodsDb) DeleteGood(ctx context.Context, goods_id int64) error {

	err := Gd.db.MustExec(`DELETE FROM goods.goods g WHERE g.goods_id = $1`, goods_id)
	if err != nil {
		return fmt.Errorf("SaveGood: %w", err)
	}
	return nil
}
