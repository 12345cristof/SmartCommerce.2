package models

type Historial struct {

	Operaciones []string
}

func (h *Historial) RegistrarOperacion(
	operacion string,
) {

	h.Operaciones =
		append(h.Operaciones, operacion)
}

func (h *Historial) UltimaOperacion() string {

	if len(h.Operaciones) == 0 {
		return ""
	}

	ultima :=
		h.Operaciones[len(h.Operaciones)-1]

	h.Operaciones =
		h.Operaciones[:len(h.Operaciones)-1]

	return ultima
}