package tests

import (
	"Goods/tests/suite"
	"github.com/brianvoe/gofakeit/v6"
	goodsv1 "github.com/kalex003/Goods_Proto/gen/go/goods"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// TODO: add token fail validation cases

func TestInsertLGoods(t *testing.T) {
	ctx, st := suite.New(t) //создаем структуру Suite для тестов

	PlaceId := int64(gofakeit.Number(1, 100))
	SkuId := int64(gofakeit.Number(1, 100))
	WbstickerId := int64(gofakeit.Number(1, 100))
	Barcode := gofakeit.Word()
	StateId := gofakeit.Word()
	ChEmployeeId := int64(gofakeit.Number(1, 100))
	OfficeId := int64(gofakeit.Number(1, 100))
	WhId := int64(gofakeit.Number(1, 100))
	TareId := int64(gofakeit.Number(1, 100))
	TareType := gofakeit.Word()

	respIns, err := st.GoodsClient.Insert(ctx, &goodsv1.InsertRequest{Structs: []*goodsv1.OneInsertRequest{
		{
			PlaceId:      PlaceId,
			SkuId:        &SkuId,
			WbstickerId:  &WbstickerId,
			Barcode:      &Barcode,
			StateId:      &StateId,
			ChEmployeeId: ChEmployeeId,
			OfficeId:     &OfficeId,
			WhId:         &WhId,
			TareId:       &TareId,
			TareType:     &TareType,
		},
	}})

	require.NoError(t, err)

	assert.NotEmpty(t, respIns.GetStructs())

	// Проверка, что массив ответов не пустой
	assert.NotEmpty(t, respIns.GetStructs())

	// Проверка каждого GoodsId в ответах
	for _, answer := range respIns.GetStructs() {
		assert.NotZero(t, answer.GoodsId, "GoodsId должен быть заполнен в ответе")
		assert.NotNil(t, answer.ChDt, "Chdt должен быть заполнен в ответе")
	}

}
