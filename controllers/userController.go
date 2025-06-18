package controllers

import (
	"backend_gin/dto"
	"backend_gin/repository"
)

func GetAllUsers() []dto.Usuario {
	var usuarios []dto.Usuario = repository.GetAll()
	return usuarios
}

func CreateNewUser(user dto.Usuario) dto.Usuario {
	var newUser = repository.Add(user)
	return newUser
}

func DeleteUser(user dto.Usuario) bool {
	repository.DeleteUser(user)
	return true
}

func EditUser(user dto.Usuario) bool {
	repository.EditUser(user)
	return true
}
