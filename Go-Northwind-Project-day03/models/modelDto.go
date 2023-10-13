package models

type CreateProductDto struct {
	ProductID       int16   `db:"product_id" json:"product_id"`
	ProductName     string  `db:"product_name" json:"product_name"`
	SupplierID      int16   `db:"supplier_id" json:"supplier_id"`
	CategoryID      int16   `db:"category_id" json:"category_id"`
	QuantityPerUnit string  `db:"quantity_per_unit" json:"quantity_per_unit"`
	UnitPrice       float64 `db:"unit_price" json:"unit_price"`
	UnitsInStock    int16   `db:"units_in_stock" json:"units_in_stock"`
	UnitsOnOrder    int16   `db:"units_on_order" json:"units_on_order"`
	ReorderLevel    int16   `db:"reorder_level" json:"reorder_level"`
	Discontinued    int32   `db:"discontinued" json:"discontinued"`
}

type ListCategoryProductDto []struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	Picture      string `json:"picture"`
	Products     []struct {
		ProductID       int     `json:"product_id"`
		ProductName     string  `json:"product_name"`
		SupplierID      int     `json:"supplier_id"`
		CategoryID      int     `json:"category_id"`
		QuantityPerUnit string  `json:"quantity_per_unit"`
		UnitPrice       float64 `json:"unit_price"`
		UnitsInStock    int     `json:"units_in_stock"`
		UnitsOnOrder    int     `json:"units_on_order"`
		ReorderLevel    int     `json:"reorder_level"`
		Discontinued    int     `json:"discontinued"`
	} `json:"products"`
}

type CreateCategoryDto struct {
	CategoryID   int16  `db:"category_id" json:"category_id"`
	CategoryName string `db:"category_name" json:"category_name"`
	Description  string `db:"description" json:"description,omitempty"`
	Picture      []byte `db:"picture" json:"picture,omitempty"`
}

type CreateCategoryProductDto struct {
	CreateCategoryDto
	Products []CreateProductDto
}
