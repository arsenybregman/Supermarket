package sql

type ApiProducts struct {
	Products []Goods `json:"products"`
}

func (db *Storage) GetGoodsApi() (ApiProducts, error) {
	var res ApiProducts
	rows, err := db.db.Query("SELECT id, title, price, description, quantity FROM products")
	if err != nil {
		return ApiProducts{}, err
	}
	defer rows.Close()

	// Обрабатываем каждую запись
	for rows.Next() {
		good := Goods{}
		err := rows.Scan(&good.Id, &good.Title, &good.Price, &good.Description, &good.Quantity)
		if err == nil {
			res.Products = append(res.Products, good)
		}
	}
	return res, nil
}
