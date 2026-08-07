package models

type RegistroUsuarios struct {

	Usuarios map[int]string
}

func NuevoRegistroUsuarios() *RegistroUsuarios {

	return &RegistroUsuarios{

		Usuarios: make(map[int]string),
	}
}