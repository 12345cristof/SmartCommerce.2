package models

type Pedido struct {
	ID        int
	Cliente   string
	Productos []Producto
	Total     float64
}