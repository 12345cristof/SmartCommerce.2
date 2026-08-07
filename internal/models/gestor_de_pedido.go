package models

type GestorPedidos struct {

	Pedidos []string
}

func (g *GestorPedidos) AgregarPedido(
	pedido string,
) {

	g.Pedidos =
		append(g.Pedidos, pedido)
}

func (g *GestorPedidos) ProcesarPedido() string {

	if len(g.Pedidos) == 0 {
		return ""
	}

	pedido := g.Pedidos[0]

	g.Pedidos =
		g.Pedidos[1:]

	return pedido
}