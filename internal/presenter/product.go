package presenter

type Products struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SupplierID  int64   `json:"supplier_id"`
}

type ProductsWithSuppliers struct {
	ProductID    int64   `json:"product_id"`
	ProductName  string  `json:"product_name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	SupplierID   int64   `json:"supplier_id"`
	SupplierName string  `json:"supplier_name"`
	ContactInfo  string  `json:"contact_info"`
}
