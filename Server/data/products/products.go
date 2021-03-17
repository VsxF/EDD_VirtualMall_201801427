package products

type InventorysJSON struct {
	Inventorys []StoreJSON `json:"Inventarios"`
}

type StoreJSON struct {
	Name string `json:"Tienda"`
	Department string `json:"Departamento"`
	Qualif int `json:"Calificacion"`
	Products []ProductJSON `json:"Productos"`
}

type ProductJSON struct {
	Name string `json:"Nombre"`
	Code int `json:"Codigo"`
	Desc string `json:"Descripcion"`
	Price float64 `json:"Precio"`
	Quant int `json:"Cantidad"`
	Image string `json:"Imagen"`
}

func NewInventorys() *InventorysJSON {
	return &InventorysJSON{}
}

func NewStore() *StoreJSON {
	return &StoreJSON{}
}

func NewProduct() *ProductJSON {
	return &ProductJSON{}
}