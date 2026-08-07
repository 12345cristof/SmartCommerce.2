package main

import (
	"fmt"

	"SmartCommerce/internal/models"
)

func main() {

	fmt.Println("===== SMARTCOMMERCE =====")

	admin :=
		models.Administrador{}

	fmt.Println(
		admin.GenerarReporte(),
	)

	producto1 :=
		models.Producto{
			ID: 1,
			Nombre: "Laptop Gamer",
			Descripcion: "RTX 4060",
			Precio: 1500,
			Stock: 5,
		}

	inventario :=
		models.Inventario{}

	inventario.AgregarProducto(
		producto1,
	)

	fmt.Println(
		"Productos en inventario:",
		inventario.TotalProductos(),
	)

	historial :=
		models.Historial{}

	historial.RegistrarOperacion(
		"Crear Producto",
	)

	historial.RegistrarOperacion(
		"Actualizar Inventario",
	)

	fmt.Println(
		"Ultima Operacion:",
		historial.UltimaOperacion(),
	)

	gestor :=
		models.GestorPedidos{}

	gestor.AgregarPedido(
		"Pedido001",
	)

	gestor.AgregarPedido(
		"Pedido002",
	)

	fmt.Println(
		"Pedido Procesado:",
		gestor.ProcesarPedido(),
	)
}