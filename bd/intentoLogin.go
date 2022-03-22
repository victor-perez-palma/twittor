package bd

import (
	"github/victor-perez-palma/twittor/models"
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}

	passwordBytes := []bytes(password)
	passwordBd := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword([]byte(passwordBd, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}