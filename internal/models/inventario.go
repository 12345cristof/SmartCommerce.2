package models

type Inventario struct {
	Productos []Producto
}

func (i *Inventario) AgregarProducto(
	p Producto,
) {

	i.Productos =
		append(i.Productos, p)
}

func (i *Inventario) TotalProductos() int {

	return len(i.Productos)
}