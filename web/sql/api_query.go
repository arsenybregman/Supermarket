package sql

type ApiProducts struct {
	Products []Goods `json:"products"`
}

func (db *Storage) GetGoodsApi() (ApiProducts, error) {
	rows, err := db.GetGoods()
	if err != nil {
		return ApiProducts{}, err
	}
	// Обрабатываем каждую запись
	
	return ApiProducts{rows}, nil
}
