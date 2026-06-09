package inventory

type item struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type Inventory []item


