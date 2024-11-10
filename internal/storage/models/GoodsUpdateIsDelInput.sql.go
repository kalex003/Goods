package dbmodels

import domainmodels "Goods/internal/domain/models"

type GoodsUpdateIsDelInput struct {
	GoodsId int64 `json:"goods_id"`
	IsDel   bool  `json:"is_del"`
}

type GoodsUpdateIsDelInputs struct {
	GoodsUpdateIsDelInputs []GoodsUpdateIsDelInput `json:"data"`
}

func ConvertGoodsFullInfoToGoodsUpdateIsDelInput(info domainmodels.GoodFullInfo) GoodsUpdateIsDelInput {
	return GoodsUpdateIsDelInput{
		GoodsId: info.GoodsId,
		IsDel:   info.IsDel,
	}
}

func ConvertGoodsFullInfoToGoodsUpdateIsDelInputs(infos domainmodels.GoodsFullInfo) GoodsUpdateIsDelInputs {
	var goods []GoodsUpdateIsDelInput
	for _, info := range infos.GoodsFullInfo {
		good := ConvertGoodsFullInfoToGoodsUpdateIsDelInput(info)
		goods = append(goods, good)
	}
	return GoodsUpdateIsDelInputs{
		GoodsUpdateIsDelInputs: goods,
	}
}
