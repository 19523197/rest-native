package repository

import (
	"rest-native/model"
)

type ProductRepository struct {
	BaseRepository BaseRepository
}

func (h *ProductRepository) GetAllProduct() (products []model.Product, err error) {
	var product model.Product
	rows, err := h.BaseRepository.Sql.Query("SELECT nama_produk, harga_produk FROM toko")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&product.Name, &product.Price)
		products = append(products, product)
	}

	rows.Close()
	return
}

func (h *ProductRepository) GetOneProduct(id int64) (product model.Product, err error) {
	row := h.BaseRepository.Sql.QueryRow("SELECT nama_produk, harga_produk FROM toko WHERE id = ?", id)
	if err != nil {
		return
	}

	row.Scan(&product.Name, &product.Price)

	return
}
