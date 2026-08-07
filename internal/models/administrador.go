package models

type Administrador struct {
	Usuario
}

func (a *Administrador) GenerarReporte() string {

	return "Reporte generado correctamente"
}