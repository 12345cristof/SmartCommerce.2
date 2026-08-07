package services

import (
	"errors"

	"SmartCommerce/internal/models"
)

type ProductoService struct {
	Productos []models.Producto
}

func (ps *ProductoService) AgregarProducto(
	p models.Producto,
) {

	ps.Productos =
		append(ps.Productos, p)
}

func (ps *ProductoService) BuscarProducto(
	id int,
) (*models.Producto, error) {

	for _, p := range ps.Productos {

		// Busca el producto por ID

		if p.ID == id {
			return &p, nil
		}
	}

	return nil,
		errors.New(
			"producto no encontrado",
		)
}