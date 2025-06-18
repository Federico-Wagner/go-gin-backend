package repository

import (
	"backend_gin/dto"
	"fmt"
)

var usuarios = []dto.Usuario{
	{ID: 1, Nombre: "Fede", Email: "fede@mail.com"},
}

// GetAll devuelve todos los usuarios
func GetAll() []dto.Usuario {
	return usuarios
}

// Add agrega un nuevo usuario y devuelve el agregado
func Add(u dto.Usuario) dto.Usuario {
	u.ID = len(usuarios) + 1
	usuarios = append(usuarios, u)
	return u
}

func DeleteUser(user dto.Usuario) {
	fmt.Print("Borrando en Repository", user)
}

func EditUser(user dto.Usuario) {
	fmt.Print("Editando en Repository", user)
}