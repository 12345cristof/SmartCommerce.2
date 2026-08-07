package models

type Usuario struct {
	id     int
	nombre string
	correo string
}

func NuevoUsuario(id int, nombre string, correo string) *Usuario {

	return &Usuario{
		id:     id,
		nombre: nombre,
		correo: correo,
	}
}

func (u *Usuario) GetNombre() string {
	return u.nombre
}

func (u *Usuario) SetNombre(nombre string) {
	u.nombre = nombre
}